package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/app/interface/restful/handler"
)

func NewAdminRoute(app fiber.Router, a *handler.AdminHandler) {
	app.Post("/assign?userId=?&todoId=?", assignTask(a))
	app.Post("/user/:userId/todo/:todoid/done", completedTask(a))
	app.Post(" /user/:userId/todo/:todoid/undo", undoTask(a))
	app.Delete("revoke?userId=?&todoId=?", revokeTask(a))
	app.Get("/todo/:todoid/user", listUserTask(a))
}

func assignTask(a *handler.AdminHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return a.AssignTask(c)
	}
}

func completedTask(a *handler.AdminHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return a.CompletedTask(c)
	}
}

func undoTask(a *handler.AdminHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return a.UndoTask(c)
	}
}

func revokeTask(a *handler.AdminHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return a.RevokeTask(c)
	}
}

func listUserTask(a *handler.AdminHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return a.ListUserTask(c)
	}
}
