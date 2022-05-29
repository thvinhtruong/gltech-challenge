package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	todoservice "github.com/thvinhtruong/legoha/app/usecase/todo/service"
)

type TodoHandler struct {
	service todoservice.TodoUseCase
}

func NewTodoHandler(service todoservice.TodoUseCase) *TodoHandler {
	return &TodoHandler{service: service}
}

func (t *TodoHandler) PostNewTodo(c *fiber.Ctx) error {
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

	err = t.service.PostNewTodo(todoDTO.Title, todoDTO.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "new todo created",
		"error":   nil,
	})
}

func (t *TodoHandler) ShowAllTodos(c *fiber.Ctx) error {
	todos, err := t.service.ShowAllTodos()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "all todos are shown",
		"error":   nil,
		"data":    todos,
	})
}

func (t *TodoHandler) ShowTodoByID(c *fiber.Ctx) error {
	todoId, _ := strconv.Atoi(c.Params("todoId"))
	todo, err := t.service.ShowTodoByID(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "todo is shown",
		"error":   nil,
		"data":    todo,
	})
}

func (t *TodoHandler) UpdateTodoInfor(c *fiber.Ctx) error {
	type UpdateTodoDTO struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
	}

	var todoDTO UpdateTodoDTO
	err := c.BodyParser(&todoDTO)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	todoId, _ := strconv.Atoi(c.Params("todoId"))
	todo := &entity.Todo{
		Title:       todoDTO.Title,
		Description: todoDTO.Description,
	}

	err = t.service.UpdateTodoInfor(todoId, todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "todo updated",
		"error":   nil,
	})
}

func (t *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	todoId, _ := strconv.Atoi(c.Params("todoId"))
	err := t.service.DeleteTodo(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "todo deleted",
		"error":   nil,
	})
}
