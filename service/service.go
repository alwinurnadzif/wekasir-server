package service

import (
	"errors"
	"strings"

	"wekasir/database"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type Errors = utils.Errors

var (
	NewErrors   = utils.NewErrors
	StrPtr      = utils.StrPtr
	GetUserinfo = utils.GetUserInfo
	sp          = utils.StrPtr
)

func PaginateResults(c *fiber.Ctx, res *gorm.DB, dest any) (paginate.Page, error) {
	pg := paginate.New()
	page := pg.With(res).Request(c.Request()).Response(dest)
	var err error
	if page.Error {
		err = fiber.ErrInternalServerError
	}
	return page, err
}

func IsGormErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func ToLower(str string) string {
	return strings.ToLower(str)
}
