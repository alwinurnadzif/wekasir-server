// Package middleware
package middleware

import (
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
)

func UserAuthentication(ctx *fiber.Ctx) error {
	// if strings.HasPrefix(ctx.Path(), "/v1/public/") || ctx.Path() == "/v1/login" {
	// 	return ctx.Next()
	// }
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token not found",
		})
	}

	const prefix = "Bearer "
	var token string
	if len(tokenHeader) > len(prefix) && tokenHeader[:len(prefix)] == prefix {
		token = tokenHeader[len(prefix):]
	} else {
		token = tokenHeader
	}

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token not found",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorization",
		})
	}

	ctx.Locals("userinfo", claims)

	return ctx.Next()
}
