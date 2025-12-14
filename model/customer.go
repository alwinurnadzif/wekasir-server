package model

import (
	"fmt"

	"wekasir/config"
	"wekasir/entity"

	"gorm.io/gorm"
)

func CustomerSearchScopeQuery() string {
	return fmt.Sprintf("%s.name ILIKE ? OR %s.description ILIKE ?", config.Customer.TableName, config.Customer.TableName)
}

func GetAllCustomers[T any](tx *gorm.DB, dest *[]T, scopes []func(db *gorm.DB) *gorm.DB) *gorm.DB {
	conn := GetConn(tx)
	res := conn.Table(config.Customer.TableName).Scopes(scopes...).Find(&dest)
	return res
}

func CreateCustomer(tx *gorm.DB, dest *entity.Customer) error {
	conn := GetConn(tx)
	return conn.Create(&dest).Error
}

func GetCustomer[T any](tx *gorm.DB, dest *T, scopes ...func(db *gorm.DB) *gorm.DB) error {
	conn := GetConn(tx)
	return conn.Table(config.Customer.TableName).Scopes(scopes...).First(&dest).Error
}

func UpdateCustomer(tx *gorm.DB, dest *entity.Customer, id uint) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Save(dest).Error
}

func DeleteCustomer(tx *gorm.DB, id string) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Delete(&entity.Customer{}).Error
}
