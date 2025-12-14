package model

import (
	"fmt"

	"wekasir/config"
	"wekasir/entity"

	"gorm.io/gorm"
)

func TransactionSearchScopeQuery() string {
	return fmt.Sprintf("%s.code ILIKE ? OR %s.name ILIKE ? OR users.name", config.Transaction.TableName, config.Customer.TableName)
}

func TransactionWithJoins() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		selectQuery := fmt.Sprintf("%s.*, users.name as user_name, customers.name as customer_name", config.Transaction.TableName)
		join := fmt.Sprintf("left join users on users.id = %s.user_id", config.Transaction.TableName)
		join2 := fmt.Sprintf("left join customers on customers.id = %s.customer_id", config.Transaction.TableName)

		return db.Select(selectQuery).Joins(join).Joins(join2)
	}
}

func GetAllTransactions[T any](tx *gorm.DB, dest *[]T, scopes []func(db *gorm.DB) *gorm.DB) *gorm.DB {
	conn := GetConn(tx)
	res := conn.Table(config.Transaction.TableName).Scopes(scopes...).Find(&dest)
	return res
}

func CreateTransaction(tx *gorm.DB, dest *entity.Transaction) error {
	conn := GetConn(tx)
	return conn.Create(&dest).Error
}

func GetTransaction[T any](tx *gorm.DB, dest *T, scopes ...func(db *gorm.DB) *gorm.DB) error {
	conn := GetConn(tx)
	return conn.Table(config.Transaction.TableName).Scopes(scopes...).First(&dest).Error
}

func DeleteTransaction(tx *gorm.DB, id string) error {
	conn := GetConn(tx)
	return conn.Where("id = ?", id).Delete(&entity.Transaction{}).Error
}
