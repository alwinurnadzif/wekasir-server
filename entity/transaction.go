package entity

import "time"

type Transaction struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Code         string    `json:"code"`
	UserID       uint      `json:"userId"`
	CustomerID   uint      `json:"customerId"`
	Date         time.Time `json:"date"`
	TotalAmount  float64   `json:"totalAmount"`
	TotalQty     float64   `json:"totalQty"`
	PaidAmount   float64   `json:"paidAmount"`
	ChangeAmount float64   `json:"changeAmount"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type TransactionRequest struct {
	UserID      uint                       `json:"userId" validate:"required"`
	CustomerID  uint                       `json:"customerId" validate:"required"`
	Date        time.Time                  `json:"date" validate:"required"`
	TotalAmount float64                    `json:"totalAmount"`
	TotalQty    float64                    `json:"totalQty"`
	PaidAmount  float64                    `json:"paidAmount"`
	Details     []TransactionDetailRequest `json:"details"`
}

type TransactionWithJoin struct {
	Transaction
	UserName     string `json:"userName"`
	CustomerName string `json:"customerName"`
}

func (req TransactionRequest) FillWithRequest(dest *Transaction) {
	dest.UserID = req.UserID
	dest.CustomerID = req.CustomerID
	dest.Date = req.Date
	dest.PaidAmount = req.PaidAmount
}
