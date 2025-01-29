package warehouse

import (
	"context"
)

func NewService(warehouseRepository WarehouseRepository) JayaService {
	return &service{warehouseRepository: warehouseRepository}
}

type service struct {
	warehouseRepository WarehouseRepository
}

func (s *service) IncomingGoods(ctx context.Context, req IncomingData) error {
	return s.warehouseRepository.IncomingGoods(ctx, req)

}

func (s *service) OutgoingGoods(ctx context.Context, req OutgoingData) error {
	return s.warehouseRepository.OutgoingGoods(ctx, req)
}

func (s *service) StockReport(ctx context.Context) (*[]Stock, error) {
	return s.warehouseRepository.StockReport(ctx)

}

type JayaService interface {
	IncomingGoods(ctx context.Context, req IncomingData) error
	OutgoingGoods(ctx context.Context, req OutgoingData) error
	StockReport(ctx context.Context) (*[]Stock, error)
}
