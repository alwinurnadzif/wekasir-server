package handler

import (
	"fmt"

	"wekasir/entity"
	"wekasir/service"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllTransactions(c *fiber.Ctx) error {
	params := utils.DataTableQueryParams{}
	if err := c.QueryParser(&params); err != nil {
		return utils.NewErrors(fiber.ErrBadRequest, utils.StrPtr("failed get datatable params"), nil).GetErrorResponse(c)
	}

	products, page, err := service.GetAllTransactions(c, params)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	if params.Paginate {
		return c.JSON(res{Message: "success", Data: page})
	}

	return c.JSON(res{Message: "success", Data: products})
}

func CreateTransaction(c *fiber.Ctx) error {
	input := entity.TransactionRequest{}
	if err := c.BodyParser(&input); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	product, err := service.CreateTransaction(c, input)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: product})
}

func GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	transaction, transactionDetails, err := service.GetTransaction(c, id)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	fmt.Printf("transaction: %v\n", transaction)
	fmt.Printf("transactionDetails: %v\n", transactionDetails)

	result := struct {
		entity.TransactionWithJoin
		Details []entity.TransactionDetailWithJoin `json:"details"`
	}{
		TransactionWithJoin: *transaction,
		Details:             *transactionDetails,
	}

	return c.JSON(res{Message: "success", Data: result})
}

func UpdateTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(entity.ProductRequest)
	if err := c.BodyParser(body); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	product, err := service.UpdateProduct(c, *body, id)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: product})
}

func DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id", "0")

	if err := service.DeleteTransaction(c, id); err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success"})
}
