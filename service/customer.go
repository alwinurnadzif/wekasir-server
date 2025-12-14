package service

import (
	"fmt"
	"strings"

	"wekasir/config"
	"wekasir/entity"
	"wekasir/model"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func GetAllCustomers(c *fiber.Ctx, params utils.DataTableQueryParams) (*[]entity.Customer, *paginate.Page, *Errors) {
	scopes := []func(db *gorm.DB) *gorm.DB{}
	searchScope := model.SearchScope(model.CustomerSearchScopeQuery(), params.Search)
	if err := utils.DataTableGetScopes(params, &scopes, searchScope); err != nil {
		return nil, nil, NewErrors(err, StrPtr("failed get datatable scopes"), nil)
	}

	customers := []entity.Customer{}
	queryRes := model.GetAllCustomers(nil, &customers, scopes)
	if queryRes.Error != nil {
		return nil, nil, NewErrors(queryRes.Error, StrPtr(fmt.Sprintf("failed get %s", strings.ToLower(config.Customer.TableName))), nil)
	}

	if params.Paginate {
		page, err := PaginateResults(c, queryRes, &customers)
		if err != nil {
			return nil, nil, NewErrors(err, StrPtr("failed paginate results"), nil)
		}

		return &customers, &page, nil
	}

	return &customers, nil, nil
	// return users, nil, nil
}

func CreateCustomer(c *fiber.Ctx, req entity.CustomerRequest) (*entity.Customer, *Errors) {
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	customer := new(entity.Customer)
	req.FillWithRequest(customer)
	customer.CreatedBy = GetUserinfo(c).Username

	if err := model.CreateCustomer(nil, customer); err != nil {
		return nil, NewErrors(err, nil, nil)
	}

	return customer, nil
}

func GetCustomer(c *fiber.Ctx, id string) (*entity.Customer, *Errors) {
	customer := new(entity.Customer)

	if err := model.GetCustomer(nil, customer, model.WhereScope(config.Customer.TableName, "id", "=", id)); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.Customer.PageName))), nil).IsNotFound(fmt.Sprintf("%s not found", ToLower(config.Customer.PageName)))
	}

	return customer, nil
}

func UpdateCustomer(c *fiber.Ctx, req entity.CustomerRequest, id string) (*entity.Customer, *Errors) {
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	// get Customer
	customer := new(entity.Customer)
	if err := model.GetCustomer(nil, customer, model.WhereScope(config.Customer.TableName, "id", "=", id)); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.Customer.PageName))), nil).IsNotFound(fmt.Sprintf("%s not found", ToLower(config.Customer.PageName)))
	}

	req.FillWithRequest(customer)
	customer.UpdatedBy = GetUserinfo(c).Username

	if err := model.UpdateCustomer(nil, customer, customer.ID); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed update %s", ToLower(config.Customer.PageName))), nil)
	}

	return customer, nil
}

func DeleteCustomer(c *fiber.Ctx, id string) *Errors {
	if err := model.DeleteCustomer(nil, id); err != nil {
		return NewErrors(err, sp(fmt.Sprintf("failed delete %s", ToLower(config.Customer.PageName))), nil)
	}

	return nil
}
