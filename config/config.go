// Package config
package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

var Product = struct {
	TableName string
	PageName  string
}{
	TableName: "products",
	PageName:  "Products",
}

var Customer = struct {
	TableName string
	PageName  string
}{
	TableName: "customers",
	PageName:  "Customers",
}

var Transaction = struct {
	TableName string
	PageName  string
}{
	TableName: "transactions",
	PageName:  "Transactions",
}

var TransactionDetail = struct {
	TableName string
	PageName  string
}{
	TableName: "transaction_details",
	PageName:  "Transaction details",
}
