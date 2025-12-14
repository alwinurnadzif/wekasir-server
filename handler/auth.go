package handler

import (
	"wekasir/entity"
	"wekasir/service"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	req := entity.LoginRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrors(err, nil, nil).GetErrorResponse(c)
	}

	token, err := service.Login(c, &req)
	if err != nil {
		return err.GetErrorResponse(c)
	}

	return c.JSON(res{Message: "success", Data: token})
}
