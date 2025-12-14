// Package entity
package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex; not null"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRequest struct {
	ID             uint   `json:"id" `
	Username       string `json:"username" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Password       string `json:"password" validate:"omitempty,min=4"`
	Email          string `json:"email" validate:"required"`
	UpdatePassword string `json:"updatePassword" validate:"omitempty,min=4"`
}

func (u UserRequest) FillWithRequest(user *User) {
	user.Username = u.Username
	user.Name = u.Name
	user.Email = u.Email
}

var UserConstraints = map[string]string{
	"idx_users_email":    "email",
	"idx_users_username": "username",
}
