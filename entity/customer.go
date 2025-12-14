package entity

import "time"

type Customer struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func (req CustomerRequest) FillWithRequest(dest *Customer) {
	dest.Name = req.Name
	dest.PhoneNumber = req.PhoneNumber
	dest.Email = req.Email
	dest.Address = req.Address
}
