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

func GetAllProducts(c *fiber.Ctx, params utils.DataTableQueryParams) (*[]entity.Product, *paginate.Page, *Errors) {
	scopes := []func(db *gorm.DB) *gorm.DB{}
	searchScope := model.SearchScope(model.ProductSearchScopeQuery(), params.Search)
	if err := utils.DataTableGetScopes(params, &scopes, searchScope); err != nil {
		return nil, nil, NewErrors(err, StrPtr("failed get datatable scopes"), nil)
	}

	products := []entity.Product{}
	queryRes := model.GetAllProducts(nil, &products, scopes)
	if queryRes.Error != nil {
		return nil, nil, NewErrors(queryRes.Error, StrPtr(fmt.Sprintf("failed get %s", strings.ToLower(config.Product.TableName))), nil)
	}

	if params.Paginate {
		page, err := PaginateResults(c, queryRes, &products)
		if err != nil {
			return nil, nil, NewErrors(err, StrPtr("failed paginate results"), nil)
		}

		return &products, &page, nil
	}

	return &products, nil, nil
	// return users, nil, nil
}

func CreateProduct(c *fiber.Ctx, req entity.ProductRequest) (*entity.Product, *Errors) {
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	product := new(entity.Product)
	req.FillWithRequest(product)
	product.CreatedBy = GetUserinfo(c).Username

	if err := model.CreateProduct(nil, product); err != nil {
		return nil, NewErrors(err, nil, nil)
	}

	return product, nil
}

func GetProduct(c *fiber.Ctx, id string) (*entity.Product, *Errors) {
	product := new(entity.Product)

	if err := model.GetProduct(nil, product, model.WhereScope(config.Product.TableName, "id", "=", id)); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.Product.PageName))), nil).IsNotFound(fmt.Sprintf("%s not found", ToLower(config.Product.PageName)))
	}

	return product, nil
}

func UpdateProduct(c *fiber.Ctx, req entity.ProductRequest, id string) (*entity.Product, *Errors) {
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	// get Product
	product := new(entity.Product)
	if err := model.GetProduct(nil, product, model.WhereScope(config.Product.TableName, "id", "=", id)); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.Product.PageName))), nil).IsNotFound(fmt.Sprintf("%s not found", ToLower(config.Product.PageName)))
	}

	req.FillWithRequest(product)
	product.UpdatedBy = GetUserinfo(c).Username

	if err := model.UpdateProduct(nil, product, product.ID); err != nil {
		return nil, NewErrors(err, sp(fmt.Sprintf("failed update %s", ToLower(config.Product.PageName))), nil)
	}

	return product, nil
}

func DeleteProduct(c *fiber.Ctx, id string) *Errors {
	if err := model.DeleteProduct(nil, id); err != nil {
		return NewErrors(err, sp(fmt.Sprintf("failed delete %s", ToLower(config.Product.PageName))), nil)
	}

	return nil
}
