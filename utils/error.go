package utils

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Errors struct {
	Err     error
	Data    any
	Message *string
}

func NewErrors(err error, message *string, data any) *Errors {
	return &Errors{Err: err, Data: data, Message: message}
}

func StrPtr(s string) *string {
	return &s
}

func (e Errors) GetErrorResponse(c *fiber.Ctx) error {
	if e.Err == nil {
		return nil
	}

	code, message := parseError(e.Err, e.Message)

	return c.Status(code).JSON(ResponseWrapper{Message: message, Data: e.Data})
}

func getErrMsg(customMsg *string, err error) string {
	m := err.Error()
	if customMsg != nil {
		m = *customMsg
	}
	return m
}

func parseError(err error, customMsg *string) (int, string) {
	// fiber error
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		switch {
		case errors.Is(err, fiber.ErrUnprocessableEntity):
			return fiber.StatusBadRequest, "failed get input"

		default:
			return fiberErr.Code, getErrMsg(customMsg, err)
		}
	}

	// validaiton err
	if errors.Is(err, ErrValidation) {
		return fiber.StatusBadRequest, err.Error()
	}

	// gorm
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.StatusNotFound, getErrMsg(customMsg, err)
	}

	return fiber.StatusInternalServerError, err.Error()
}

func (e *Errors) IsErrGormUnique(constraints map[string]string) *Errors {
	var pgErr *pgconn.PgError

	if errors.As(e.Err, &pgErr) && pgErr.Code == "23505" {
		if c, ok := constraints[pgErr.ConstraintName]; ok {
			validationErr := []ValidationError{
				{Field: c, Index: nil, Message: fmt.Sprintf("%v already used", c)},
			}

			e.Err = ErrValidation
			e.Data = validationErr
			e.Message = nil
		}
	}

	return e
}

func (e *Errors) IsNotFound(msg string) *Errors {
	if errors.Is(e.Err, gorm.ErrRecordNotFound) {
		e.Message = &msg
	}

	return e
}
