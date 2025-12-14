// Package routes provides utility functions such as error handling, logging, and helpers.
package routes

import (
	"wekasir/config"
	"wekasir/handler"
	"wekasir/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	v1 := r.Group("v1")

	v1.Static("/public", config.ProjectRootPath+"/public")
	v1.Post("login", handler.Login)

	requiresAuth := v1.Group("", middleware.UserAuthentication)

	// users
	users := requiresAuth.Group("users")
	users.Get("/", handler.GetAllUser)
	users.Post("/", handler.CreateUser)
	users.Get(":id", handler.GetUser)
	users.Put(":id/update", handler.UpdateUser)
	users.Delete(":id/delete", handler.DeleteUser)

	// products
	products := requiresAuth.Group("products")
	products.Get("/", handler.GetAllProducts)
	products.Post("/", handler.CreateProduct)
	products.Get(":id", handler.GetProduct)
	products.Put(":id/update", handler.UpdateProduct)
	products.Delete(":id/delete", handler.DeleteProduct)
}
