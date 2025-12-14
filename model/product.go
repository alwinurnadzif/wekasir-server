package model

import (
	"fmt"

	"wekasir/config"
	"wekasir/entity"

	"gorm.io/gorm"
)

func ProductSearchScopeQuery() string {
	return fmt.Sprintf("%s.name ILIKE ? OR %s.description ILIKE ?", config.Product.TableName, config.Product.TableName)
}

func GetAllProducts[T any](tx *gorm.DB, dest *[]T, scopes []func(db *gorm.DB) *gorm.DB) *gorm.DB {
	conn := GetConn(tx)
	res := conn.Table(config.Product.TableName).Scopes(scopes...).Find(&dest)
	return res
}

func CreateProduct(tx *gorm.DB, dest *entity.Product) error {
	conn := GetConn(tx)
	return conn.Create(&dest).Error
}

func GetProduct[T any](tx *gorm.DB, dest *T, scopes ...func(db *gorm.DB) *gorm.DB) error {
	conn := GetConn(tx)
	return conn.Table(config.Product.TableName).Scopes(scopes...).First(&dest).Error
}

func UpdateProduct(tx *gorm.DB, dest *entity.Product, id uint) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Save(dest).Error
}

func DeleteProduct(tx *gorm.DB, id string) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Delete(&entity.Product{}).Error
}

func UpdateProducts(tx *gorm.DB, dest *[]entity.Product) error {
	conn := GetConn(tx)
	return conn.Save(&dest).Error
}
