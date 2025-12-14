package handler

import (
	"wekasir/entity"
	"wekasir/service"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomers(c *fiber.Ctx) error {
	params := utils.DataTableQueryParams{}
	if err := c.QueryParser(&params); err != nil {
		return utils.NewErrors(fiber.ErrBadRequest, utils.StrPtr("failed get datatable params"), nil).GetErrorResponse(c)
	}

	customers, page, err := service.GetAllCustomers(c, params)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	if params.Paginate {
		return c.JSON(res{Message: "success", Data: page})
	}

	return c.JSON(res{Message: "success", Data: customers})
}

func CreateCustomer(c *fiber.Ctx) error {
	input := entity.CustomerRequest{}
	if err := c.BodyParser(&input); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	customer, err := service.CreateCustomer(c, input)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: customer})
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	customer, err := service.GetCustomer(c, id)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: customer})
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(entity.CustomerRequest)
	if err := c.BodyParser(body); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	customer, err := service.UpdateCustomer(c, *body, id)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: customer})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id", "0")

	if err := service.DeleteCustomer(c, id); err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success"})
}
