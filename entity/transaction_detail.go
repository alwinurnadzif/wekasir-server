package entity

import "time"

type TransactionDetail struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TransactionID uint      `json:"transactionId"`
	ProductID     uint      `json:"productId"`
	Qty           uint      `json:"qty"`
	Price         float64   `json:"price"`
	CreatedBy     string    `json:"createdBy"`
	UpdatedBy     string    `json:"updatedBy"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type TransactionDetailRequest struct {
	ProductID uint    `json:"productId"`
	Qty       uint    `json:"qty"`
	Price     float64 `json:"price"`
}

type TransactionDetailWithJoin struct {
	TransactionDetail
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
}

func (req TransactionDetailRequest) FillWithRequest(dest *TransactionDetail) {
	dest.ProductID = req.ProductID
	dest.Qty = req.Qty
	dest.Price = req.Price
}
