package service

import (
	"errors"
	"time"

	"wekasir/entity"
	"wekasir/model"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, req *entity.LoginRequest) (string, *utils.Errors) {
	// validate input
	validationErr, err := utils.ValidatePayload(req)
	if err != nil {
		return "", utils.NewErrors(err, nil, validationErr)
	}

	// get user by username
	user, err := model.GetUser(nil, []func(db *gorm.DB) *gorm.DB{model.WhereScope("users", "username", "=", req.Username)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", utils.NewErrors(err, StrPtr("username not found"), nil)
		}
		return "", utils.NewErrors(err, nil, validationErr)
	}

	// check password
	if ok := utils.CheckHashPassword(req.Password, user.Password); !ok {
		return "", utils.NewErrors(fiber.ErrBadRequest, utils.StrPtr("invalid password"), nil)
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return "", utils.NewErrors(fiber.ErrInternalServerError, utils.StrPtr("failed generate token"), nil)
	}

	return token, nil
}
