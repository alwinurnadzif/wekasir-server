package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type QueryParamOrder struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type QueryParamFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchScopeFunc func(search string) func(db *gorm.DB) *gorm.DB

type DataTableQueryParams struct {
	Search   string `query:"search,omitempty"`
	Orders   string `query:"orders,omitempty"`
	Filters  string `query:"filters,omitempty"`
	Paginate bool   `query:"paginate,omitempty"`
}

func getOrderScope(order QueryParamOrder) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		orderQuery := order.Field + " " + order.Value
		return db.Order(orderQuery)
	}
}

func GetOrderScopes(orderParams string, orderScopes *[]func(db *gorm.DB) *gorm.DB) error {
	// parse json
	var orders []QueryParamOrder

	if err := json.Unmarshal([]byte(orderParams), &orders); err != nil {
		return err
	}

	for _, order := range orders {
		if order.Value != "" {
			scope := getOrderScope(order)
			*orderScopes = append(*orderScopes, scope)
		}
	}

	return nil
}

func getFilterScope(filter QueryParamFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.Operator == "ILIKE" {
			filter.Value = "%" + filter.Value + "%"
		}

		if filter.Operator == "in" {
			values := strings.Split(filter.Value, ",")
			filterQuery := filter.Field + " " + filter.Operator + " ?"
			return db.Where(filterQuery, values)
		}

		if filter.Operator == "not in" {
			values := strings.Split(filter.Value, ",")
			filterQuery := filter.Field + " " + filter.Operator + " ?"
			return db.Where(filterQuery, values)
		}

		if filter.Operator == "is null" {
			filterQuery := filter.Field + " IS NULL"
			return db.Where(filterQuery)
		}

		if filter.Operator == "between" {
			dates := strings.Split(filter.Value, ".")
			if len(dates) != 2 {
				return db
			}
			filterQuery := filter.Field + " " + "BETWEEN ? AND ?"
			return db.Where(filterQuery, dates[0], dates[1])
		}

		filterQuery := filter.Field + " " + filter.Operator + " ?"
		return db.Where(filterQuery, filter.Value)
	}
}

func GetFilterScopes(filterParams string, filterScopes *[]func(db *gorm.DB) *gorm.DB) error {
	var filters []QueryParamFilter

	if err := json.Unmarshal([]byte(filterParams), &filters); err != nil {
		return err
	}

	for _, filter := range filters {
		if filter.Operator != "" && filter.Value != "" {
			scope := getFilterScope(filter)
			*filterScopes = append(*filterScopes, scope)
		}
	}

	return nil
}

func GetFilterQueryParams(params DataTableQueryParams) (*[]QueryParamFilter, error) {
	var filters []QueryParamFilter

	if err := json.Unmarshal([]byte(params.Filters), &filters); err != nil {
		return nil, err
	}

	return &filters, nil
}

func DataTableGetScopes(params DataTableQueryParams, scopes *[]func(db *gorm.DB) *gorm.DB, searchScope func(db *gorm.DB) *gorm.DB) error {
	if params.Search != "" {
		*scopes = append(*scopes, searchScope)
	}

	if params.Filters != "" {
		var filterScopes []func(db *gorm.DB) *gorm.DB
		if err := GetFilterScopes(params.Filters, &filterScopes); err != nil {
			return err
		}
		*scopes = append(*scopes, filterScopes...)
	}

	if params.Orders != "" {
		var orderScopes []func(db *gorm.DB) *gorm.DB
		if err := GetOrderScopes(params.Orders, &orderScopes); err != nil {
			return err
		}
		*scopes = append(*scopes, orderScopes...)
	}

	return nil
}

func DataTableFilterToQuery(param DataTableQueryParams) (string, error) {
	var query string
	var filters []QueryParamFilter

	if err := json.Unmarshal([]byte(param.Filters), &filters); err != nil {
		return "", err
	}

	for _, filter := range filters {
		if filter.Operator != "" && filter.Value != "" {
			query += fmt.Sprintf(" AND %v %v %v", filter.Field, filter.Operator, filter.Value)
		}
	}

	return query, nil
}

func ArrayToFilter(filters []string) string {
	if len(filters) <= 0 {
		return ""
	}

	results := ""
	for i, f := range filters {
		if i == 0 {
			results += fmt.Sprintf("WHERE %v", f)
		} else {
			results += fmt.Sprintf(" AND %v", f)
		}
	}
	return results
}
