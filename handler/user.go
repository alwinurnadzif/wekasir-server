// Package handler
package handler

import (
	"wekasir/entity"
	"wekasir/service"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	params := utils.DataTableQueryParams{}
	if err := c.QueryParser(&params); err != nil {
		return utils.NewErrors(fiber.ErrBadRequest, utils.StrPtr("failed get datatable params"), nil).GetErrorResponse(c)
	}

	users, page, err := service.GetAllUsers(c, params)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	if params.Paginate {
		return c.JSON(res{Message: "success", Data: page})
	}

	return c.JSON(res{Message: "success", Data: users})
}

func CreateUser(c *fiber.Ctx) error {
	req := entity.UserRequest{}

	if err := c.BodyParser(&req); err != nil {
		utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	user, err := service.CreateUser(c, &req)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: user})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id", "0")
	user, err := service.GetUser(c, id)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: user})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id", "0")
	req := entity.UserRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	users, err := service.UpdateUser(c, id, req)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: users})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id", "0")

	if err := service.DeleteUser(c, id); err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success"})
}
