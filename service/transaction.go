package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"wekasir/config"
	"wekasir/entity"
	"wekasir/model"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func GetAllTransactions(c *fiber.Ctx, params utils.DataTableQueryParams) (*[]entity.TransactionWithJoin, *paginate.Page, *Errors) {
	scopes := []func(db *gorm.DB) *gorm.DB{}
	scopes = append(scopes, model.TransactionWithJoins())
	searchScope := model.SearchScope(model.TransactionSearchScopeQuery(), params.Search)
	if err := utils.DataTableGetScopes(params, &scopes, searchScope); err != nil {
		return nil, nil, NewErrors(err, StrPtr("failed get datatable scopes"), nil)
	}

	transactions := []entity.TransactionWithJoin{}
	queryRes := model.GetAllTransactions(nil, &transactions, scopes)
	if queryRes.Error != nil {
		return nil, nil, NewErrors(queryRes.Error, StrPtr(fmt.Sprintf("failed get %s", strings.ToLower(config.Product.TableName))), nil)
	}

	if params.Paginate {
		page, err := PaginateResults(c, queryRes, &transactions)
		if err != nil {
			return nil, nil, NewErrors(err, StrPtr("failed paginate results"), nil)
		}

		return &transactions, &page, nil
	}

	return &transactions, nil, nil
}

func CreateTransaction(c *fiber.Ctx, req entity.TransactionRequest) (*entity.Transaction, *Errors) {
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	transaction := new(entity.Transaction)
	req.FillWithRequest(transaction)
	transaction.CreatedBy = GetUserinfo(c).Username

	if err := WithTransaction(func(tx *gorm.DB) error {
		// fill detail
		details := []entity.TransactionDetail{}
		productIds := []uint{}

		var totalAmount float64
		var totalQty uint

		for _, reqDetail := range req.Details {
			detail := entity.TransactionDetail{}
			reqDetail.FillWithRequest(&detail)
			detail.CreatedBy = GetUserinfo(c).Username
			details = append(details, detail)
			productIds = append(productIds, detail.ProductID)
			totalQty += detail.Qty
			totalAmount += detail.Price
		}

		// get product by detail
		productScope := func(db *gorm.DB) *gorm.DB {
			return db.Where("products.id IN ?", productIds)
		}

		_products := []entity.Product{}
		if res := model.GetAllProducts(tx, &_products, []func(db *gorm.DB) *gorm.DB{productScope}); res.Error != nil {
			return res.Error
		}

		products := map[uint]entity.Product{}
		for _, p := range _products {
			products[p.ID] = p
		}

		for _, d := range details {
			if _, ok := products[d.ProductID]; ok {
				currProductQty := products[d.ProductID].Qty
				res := currProductQty - d.Qty
				if currProductQty < d.Qty {
					message := fmt.Sprintf("qty %s is insufficient", products[d.ProductID].Name)
					return errors.New(message)
				}

				p := products[d.ProductID]
				p.Qty = res
				products[d.ProductID] = p

			} else {
				return gorm.ErrRecordNotFound
			}
		}

		// create transaction
		code := time.Now().Format("20060102150405")
		transaction.Code = fmt.Sprintf("TR-%s", code)
		transaction.TotalAmount = totalAmount
		transaction.TotalQty = float64(totalQty)
		if transaction.PaidAmount < totalAmount {
			return errors.New("insufficient funds to complete the transaction")
		}
		transaction.ChangeAmount = transaction.PaidAmount - totalAmount

		if err := model.CreateTransaction(tx, transaction); err != nil {
			return err
		}

		for i := range details {
			details[i].TransactionID = transaction.ID
		}

		//  create transaction detail
		if err := model.CreateTransactionDetails(tx, &details); err != nil {
			return err
		}

		// update product
		_p := []entity.Product{}
		for _, v := range products {
			_p = append(_p, v)
		}
		if err := model.UpdateProducts(tx, &_p); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, NewErrors(err, nil, nil)
	}

	return transaction, nil
}

func GetTransaction(c *fiber.Ctx, id string) (*entity.TransactionWithJoin, *[]entity.TransactionDetailWithJoin, *Errors) {
	transaction := new(entity.TransactionWithJoin)

	if err := model.GetTransaction(nil, transaction, model.TransactionWithJoins(), model.WhereScope(config.Transaction.TableName, "id", "=", id)); err != nil {
		return nil, nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.Transaction.PageName))), nil).IsNotFound(fmt.Sprintf("%s not found", ToLower(config.Transaction.PageName)))
	}

	transactionDetails := []entity.TransactionDetailWithJoin{}
	if err := model.GetTransactionDetails(nil, &transactionDetails, model.TransactionDetailWithJoins(), model.WhereScope(config.TransactionDetail.TableName, "transaction_id", "=", transaction.ID)); err != nil {
		return nil, nil, NewErrors(err, sp(fmt.Sprintf("failed get %s", ToLower(config.TransactionDetail.PageName))), nil)
	}

	return transaction, &transactionDetails, nil
}

func DeleteTransaction(c *fiber.Ctx, id string) *Errors {
	if err := WithTransaction(func(tx *gorm.DB) error {
		if err := model.DeleteTransaction(tx, id); err != nil {
			return err
		}

		if err := model.DeleteTransactionDetails(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return NewErrors(err, sp("failed delete transaction"), nil)
	}
	return nil
}
