package entity

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Qty         uint      `json:"qty"`
	Price       float64   `json:"price"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Qty         uint    `json:"qty"`
	Price       float64 `json:"price"`
}

func (req ProductRequest) FillWithRequest(dest *Product) {
	dest.Name = req.Name
	dest.Description = req.Description
	dest.Qty = req.Qty
	dest.Price = req.Price
}
