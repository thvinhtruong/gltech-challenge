package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/usecase/tasklist"
)

func NewAdminHandler(app fiber.Router, service tasklist.TaskListUseCase) {
	app.Post("/assign?userId=?&todoId=?", assignTask(service))
	app.Post("/user/:userId/todo/:todoid/done", completedTask(service))
	app.Post(" /user/:userId/todo/:todoid/undo", undoTask(service))
	app.Delete("revoke?userId=?&todoId=?", revokeTask(service))
	app.Get("/todo/:todoid/user", listUserTask(service))
}

func assignTask(service tasklist.TaskListUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		err = service.AssignTask(userId, todoId)
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

func completedTask(service tasklist.TaskListUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		err = service.CompletedTask(userId, todoId)
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

func undoTask(service tasklist.TaskListUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		err = service.UndoTask(userId, todoId)
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

func revokeTask(service tasklist.TaskListUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		err = service.RevokeTask(userId, todoId)
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

func listUserTask(service tasklist.TaskListUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := strconv.Atoi(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		tasks, err := service.ListUsersTasks(userId)
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
			"data":   tasks,
		})
	}
}
