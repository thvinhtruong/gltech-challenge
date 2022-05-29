package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	tasklistservice "github.com/thvinhtruong/legoha/app/usecase/tasklist/service"
)

type AdminHandler struct {
	service tasklistservice.TaskListUseCase
}

func NewAdminHandler(service tasklistservice.TaskListUseCase) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

func (a *AdminHandler) AssignTask(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	todoId, err := strconv.Atoi(c.Query("todoId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	err = a.service.AssignTask(userId, todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "task assigned",
		"error":   nil,
	})
}

func (a *AdminHandler) CompletedTask(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	todoId, err := strconv.Atoi(c.Params("todoid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	err = a.service.CompletedTask(userId, todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "task completed",
		"error":   nil,
	})
}

func (a *AdminHandler) UndoTask(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	todoId, err := strconv.Atoi(c.Params("todoid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	err = a.service.UndoTask(userId, todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "task undone",
		"error":   nil,
	})
}

func (a *AdminHandler) RevokeTask(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	todoId, err := strconv.Atoi(c.Query("todoId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	err = a.service.RevokeTask(userId, todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "task revoked",
		"error":   nil,
	})
}

func (a *AdminHandler) ListUserTask(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	tasks, err := a.service.ListUsersTasks(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "tasks is shown for user " + c.Params("userId"),
		"error":   nil,
		"data":    tasks,
	})
}
