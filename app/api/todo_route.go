package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/app/interface/restful/handler"
)

func NewTodoRoutes(app fiber.Router, t *handler.TodoHandler) {
	app.Post("/", postNewTodo(t))
	app.Get("/all", showAllTodos(t))
	app.Get("/:todoId", showTodoByID(t))
	app.Put("/:todoId", updateTodo(t))
	app.Delete("/:todoId", deleteTodo(t))
}

func postNewTodo(t *handler.TodoHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return t.PostNewTodo(c)
	}
}

func showAllTodos(t *handler.TodoHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return t.ShowAllTodos(c)
	}
}

func showTodoByID(t *handler.TodoHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return t.ShowTodoByID(c)
	}
}

func updateTodo(t *handler.TodoHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return t.UpdateTodoInfor(c)
	}
}

func deleteTodo(t *handler.TodoHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return t.DeleteTodo(c)
	}
}
