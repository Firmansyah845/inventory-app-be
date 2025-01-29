package warehouse

type IncomingData struct {
	WhsIdf       int        `json:"whsIdf"`
	TrxInNo      string     `json:"trxInNo"`
	TrxInDate    string     `json:"trxInDate"`
	TrxInSuppIdf int        `json:"trxInSuppIdf"`
	TrxInNotes   string     `json:"trxInNotes"`
	Products     []Products `json:"products"`
}

type Products struct {
	ProductId int `json:"productId"`
	QtyDus    int `json:"qtyDus"`
	QtyPcs    int `json:"qtyPcs"`
}

type OutgoingData struct {
	WhsIdf        int        `json:"whsIdf"`
	TrxOutNo      string     `json:"trxOutNo"`
	TrxOutDate    string     `json:"trxOutDate"`
	TrxOutSuppIdf int        `json:"trxOutSuppIdf"`
	TrxOutNotes   string     `json:"trxOutNotes"`
	Products      []Products `json:"products"`
}

type Stock struct {
	WhsName     string `json:"whsName"`
	ProductName string `json:"productName"`
	QtyDus      int    `json:"qtyDus"`
	QtyPcs      int    `json:"qtyPcs"`
}
