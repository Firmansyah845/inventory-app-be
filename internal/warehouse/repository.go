package warehouse

import (
	"context"
	"database/sql"
)

func NewDBStore(db *sql.DB) WarehouseRepository {
	return &warehouseRepo{db: db}
}

type warehouseRepo struct {
	db *sql.DB
}

func (r *warehouseRepo) IncomingGoods(ctx context.Context, incomingData IncomingData) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Insert header (Transaction)
	var trxInPK int64
	query := `INSERT INTO TransaksiPenerimaanBarangHeader (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes) 
	          VALUES (?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, incomingData.TrxInNo, incomingData.WhsIdf, incomingData.TrxInDate, incomingData.TrxInSuppIdf, incomingData.TrxInNotes)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Get the last inserted ID
	trxInPK, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert details (Products)
	for _, product := range incomingData.Products {
		query = `INSERT INTO TransaksiPenerimaanBarangDetail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs)
		         VALUES (?, ?, ?, ?)`
		_, err := tx.ExecContext(ctx, query, trxInPK, product.ProductId, product.QtyDus, product.QtyPcs)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *warehouseRepo) OutgoingGoods(ctx context.Context, outgoingData OutgoingData) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Insert header (Transaction)
	var trxOutPK int64
	query := `INSERT INTO TransaksiPengeluaranBarangHeader (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes) 
	          VALUES (?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, outgoingData.TrxOutNo, outgoingData.WhsIdf, outgoingData.TrxOutDate, outgoingData.TrxOutSuppIdf, outgoingData.TrxOutNotes)
	if err != nil {
		tx.Rollback()
		return err
	}

	trxOutPK, err = res.LastInsertId()

	// Insert details (Products)
	for _, product := range outgoingData.Products {
		query = `INSERT INTO TransaksiPengeluaranBarangDetail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs)
		         VALUES (?, ?, ?, ?)`
		_, err := tx.ExecContext(ctx, query, trxOutPK, product.ProductId, product.QtyDus, product.QtyPcs)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *warehouseRepo) StockReport(ctx context.Context) (*[]Stock, error) {
	var result []Stock
	query := `SELECT whs.WhsName, p.ProductName,
                     COALESCE(SUM(d.TrxInDQtyDus), 0) - COALESCE(SUM(od.TrxOutDQtyDus), 0) AS QtyDus,
                     COALESCE(SUM(d.TrxInDQtyPcs), 0) - COALESCE(SUM(od.TrxOutDQtyPcs), 0) AS QtyPcs
	          FROM MasterWarehouse whs
	          LEFT JOIN TransaksiPenerimaanBarangHeader h ON whs.WhsPK = h.WhsIdf
	          LEFT JOIN TransaksiPenerimaanBarangDetail d ON h.TrxInPK = d.TrxInIDF
	          LEFT JOIN MasterProduct p ON p.ProductPK = d.TrxInDProductIdf
	          LEFT JOIN TransaksiPengeluaranBarangDetail od ON od.TrxOutDProductIdf = d.TrxInDProductIdf
	          GROUP BY whs.WhsName, p.ProductName`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Stock
		if err := rows.Scan(&item.WhsName, &item.ProductName, &item.QtyDus, &item.QtyPcs); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return &result, nil
}

type WarehouseRepository interface {
	IncomingGoods(ctx context.Context, incomingData IncomingData) error
	OutgoingGoods(ctx context.Context, outgoingData OutgoingData) error
	StockReport(ctx context.Context) (*[]Stock, error)
}
