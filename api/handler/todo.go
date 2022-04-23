package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entity "github.com/thvinhtruong/legoha/entities"
	"github.com/thvinhtruong/legoha/usecase/todo"
)

func NewTodoHandler(app fiber.Router, service todo.TodoUseCase) {
	app.Post("/todo", createTodo(service))
	app.Get("/todo/all", listTodo(service))
	app.Get("/todo/:todoId", getTodo(service))
	app.Patch("/todo/:todoId", patchTodo(service))
	app.Delete("/todo/:todoId", deleteTodo(service))
}

func createTodo(service todo.TodoUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		type NewTodoDTO struct {
			Title       string `json:"title"`
			Description string `json:"desc"`
		}

		var todoDTO NewTodoDTO
		err := c.BodyParser(&todoDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		err = service.CreateTodo(todoDTO.Title, todoDTO.Description)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
		})
	}
}

func listTodo(service todo.TodoUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todos, err := service.ListTodos()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
			"data":   todos,
		})
	}
}

func getTodo(service todo.TodoUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoId, _ := strconv.Atoi(c.Params("todoId"))
		todo, err := service.GetTodoByID(todoId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
			"data":   todo,
		})
	}
}

func patchTodo(service todo.TodoUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type UpdateDTO struct {
			Title       string `json:"title"`
			Description string `json:"desc"`
			Completed   bool   `json:"completed"`
		}

		var updateDTO UpdateDTO
		err := c.BodyParser(&updateDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		todoId, _ := strconv.Atoi(c.Params("todoId"))
		todo := &entity.Todo{Description: updateDTO.Description, Title: updateDTO.Title, Completed: updateDTO.Completed}
		err = service.PatchTodo(todoId, todo)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
			"data":   todo,
		})
	}
}

func deleteTodo(service todo.TodoUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoId, _ := strconv.Atoi(c.Params("todoId"))
		err := service.DeleteTodo(todoId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
		})
	}
}
