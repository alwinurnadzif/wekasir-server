package model

import (
	"fmt"

	"wekasir/config"
	"wekasir/entity"

	"gorm.io/gorm"
)

func TransactionDetailWithJoins() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		selectQuery := fmt.Sprintf("%s.*, products.name as product_name", config.TransactionDetail.TableName)
		join := fmt.Sprintf("left join products on products.id = %s.product_id", config.TransactionDetail.TableName)

		return db.Select(selectQuery).Joins(join)
	}
}

func CreateTransactionDetails(tx *gorm.DB, b *[]entity.TransactionDetail) error {
	conn := GetConn(tx)
	return conn.Create(&b).Error
}

func GetTransactionDetails[T any](tx *gorm.DB, dest *[]T, scopes ...func(db *gorm.DB) *gorm.DB) error {
	conn := GetConn(tx)
	return conn.Table(config.TransactionDetail.TableName).Scopes(scopes...).Find(&dest).Error
}

func DeleteTransactionDetails(tx *gorm.DB, id string) error {
	conn := GetConn(tx)
	return conn.Where("transaction_id = ?", id).Delete(&entity.TransactionDetail{}).Error
}
