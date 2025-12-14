package model

import (
	"fmt"
	"strings"

	"wekasir/database"

	"gorm.io/gorm"
)

func GetConn(tx *gorm.DB) *gorm.DB {
	conn := tx
	if tx == nil {
		conn = database.DB
	}
	return conn
}

func WhereScope(tableName, column, operator string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		head := fmt.Sprintf("%s.%s %s ?", tableName, column, operator)
		return db.Where(head, value)
	}
}

func SearchScope(query, search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		values := GetSearchScopeValues(query, search)
		return db.Where(query, values...)
	}
}

func GetSearchScopeValues(query string, value any) []any {
	count := strings.Count(query, "?")
	values := []any{}
	for range count {
		values = append(values, "%"+value.(string)+"%")
	}

	return values
}
