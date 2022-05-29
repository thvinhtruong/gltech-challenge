package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/app/interface/restful/handler"
)

func NewUserRoutes(app fiber.Router, u *handler.UserHandler) {
	app.Post("/", registerUser(u))
	app.Post("/login", loginUser(u))
	app.Get("/all", findAllUsers(u))
	app.Get("/:userId", getUserByID(u))
	app.Put("/:userId", updateUser(u))
	app.Delete("/:userId", deleteUser(u))
}

func registerUser(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.RegisterUser(c)
	}
}

func loginUser(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.LoginUser(c)
	}
}

func findAllUsers(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.FindAllUsers(c)
	}
}

func getUserByID(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.FindUserById(c)
	}
}

func updateUser(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.UpdateUserInfor(c)
	}
}

func deleteUser(u *handler.UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return u.DeleteUser(c)
	}
}
