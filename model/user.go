// Package model
package model

import (
	"wekasir/database"
	"wekasir/entity"

	"gorm.io/gorm"
)

func UserSearchScopeQuery() string {
	return "users.username ILIKE ? OR users.name ILIKE ? OR users.email ILIKE ?"
}

func GetAllUser(tx *gorm.DB, scopes []func(db *gorm.DB) *gorm.DB) (*[]entity.User, *gorm.DB) {
	conn := GetConn(tx)
	users := []entity.User{}
	res := conn.Scopes(scopes...).Find(&users)
	return &users, res
}

func GetUser(tx *gorm.DB, scopes []func(db *gorm.DB) *gorm.DB) (*entity.User, error) {
	conn := GetConn(tx)
	user := entity.User{}
	err := conn.Scopes(scopes...).First(&user).Error
	return &user, err
}

func CreateUser(req entity.UserRequest, user *entity.User) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&user).Error
	})
}

func UpdateUser(tx *gorm.DB, id string, user *entity.User) error {
	conn := GetConn(tx)
	return conn.Save(user).Where("users.id = ?", id).Error
}

func DeleteUser(tx *gorm.DB, id string) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Delete(&entity.User{}).Error
}
