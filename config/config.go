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
