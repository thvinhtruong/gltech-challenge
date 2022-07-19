package presenter

import "github.com/gofiber/fiber/v2"

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

func BindTodoRequest(c *fiber.Ctx) (TodoRequest, error) {
	var req TodoRequest
	if err := c.Bind(fiber.Map{
		"title":       &req.Title,
		"description": &req.Description,
		"created_at":  &req.CreatedAt,
	}); err != nil {
		c.Status(400).SendString(err.Error())
		return req, err
	}

	if err := ValidateTodo(req.Title); err != nil {
		c.Status(400).SendString(err.Error())
		return req, err
	}

	return req, nil
}
