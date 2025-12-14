package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
)

type Userinfo struct {
	Username string
	UserID   uint
}

func GetUserInfo(c *fiber.Ctx) Userinfo {
	res := Userinfo{}

	claims := c.Locals("userinfo")
	if claims == nil {
		return res
	}

	userID, ok := claims.(jwt.MapClaims)["id"]
	if !ok {
		log.Warn("claims id not found")
	}

	res.UserID = uint(userID.(float64))

	username, ok := claims.(jwt.MapClaims)["username"]
	if !ok {
		log.Warn("claims username not found")
	}

	res.Username = username.(string)

	return res
}
