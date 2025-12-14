// Package service
package service

import (
	"fmt"

	"wekasir/entity"
	"wekasir/model"
	"wekasir/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx, params utils.DataTableQueryParams) (*[]entity.User, *paginate.Page, *Errors) {
	scopes := []func(db *gorm.DB) *gorm.DB{}
	searchScope := model.SearchScope(model.UserSearchScopeQuery(), params.Search)
	if err := utils.DataTableGetScopes(params, &scopes, searchScope); err != nil {
		return nil, nil, NewErrors(err, StrPtr("failed get datatable scopes"), nil)
	}

	users, queryRes := model.GetAllUser(nil, scopes)
	if queryRes.Error != nil {
		return nil, nil, NewErrors(queryRes.Error, StrPtr("failed get users"), nil)
	}

	if params.Paginate {
		page, err := PaginateResults(c, queryRes, users)
		if err != nil {
			return nil, nil, NewErrors(err, StrPtr("failed paginate results"), nil)
		}

		return users, &page, nil
	}

	return users, nil, nil
}

func CreateUser(c *fiber.Ctx, req *entity.UserRequest) (*entity.User, *utils.Errors) {
	// validate
	if validationErr, err := utils.ValidatePayload(req); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	user := entity.User{}
	req.FillWithRequest(&user)
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, utils.NewErrors(err, utils.StrPtr("failed hash password"), nil)
	}
	user.Password = hashedPassword
	user.CreatedBy = GetUserinfo(c).Username

	if err := model.CreateUser(*req, &user); err != nil {
		return nil, utils.NewErrors(err, utils.StrPtr("failed create user"), nil).IsErrGormUnique(entity.UserConstraints)
	}

	return &user, nil
}

func GetUser(c *fiber.Ctx, id string) (*entity.User, *Errors) {
	scope := model.WhereScope("users", "id", "=", id)
	users, err := model.GetUser(nil, []func(db *gorm.DB) *gorm.DB{scope})
	if err != nil {
		return nil, NewErrors(err, nil, nil)
	}

	return users, nil
}

func UpdateUser(c *fiber.Ctx, id string, input entity.UserRequest) (*entity.User, *Errors) {
	// validate input
	if validationErr, err := utils.ValidatePayload(input); err != nil {
		return nil, utils.NewErrors(err, nil, validationErr)
	}

	// get users
	scope := model.WhereScope("users", "id", "=", id)
	user, err := model.GetUser(nil, []func(db *gorm.DB) *gorm.DB{scope})
	if err != nil {
		return nil, NewErrors(err, nil, nil)
	}

	input.FillWithRequest(user)
	user.UpdatedBy = GetUserinfo(c).Username
	fmt.Printf("user: %v\n", user)
	if input.UpdatePassword != "" {
		hashed, err := utils.HashPassword(input.UpdatePassword)
		if err != nil {
			return nil, NewErrors(fiber.ErrInternalServerError, StrPtr("failed hash password"), nil)
		}
		user.Password = hashed
	}

	if err := model.UpdateUser(nil, id, user); err != nil {
		return nil, utils.NewErrors(err, StrPtr("failed update user"), nil).IsErrGormUnique(entity.UserConstraints)
	}

	return user, nil
}

func DeleteUser(c *fiber.Ctx, id string) *Errors {
	if err := model.DeleteUser(nil, id); err != nil {
		return utils.NewErrors(err, StrPtr("failed delete user"), nil)
	}

	return nil
}
