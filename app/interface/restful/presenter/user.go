package presenter

import "github.com/gofiber/fiber/v2"

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// user not auth
func BindUserRequest(c *fiber.Ctx, WithAuth bool) (UserRequest, error) {
	var req UserRequest
	if err := c.Bind(fiber.Map{
		"name":     &req.Name,
		"username": &req.Username,
		"password": &req.Password,
	}); err != nil {
		c.Status(400).SendString(err.Error())
		return req, err
	}
	if WithAuth {
		if err := ValidateAuth(req.Username, req.Password); err != nil {
			c.Status(400).SendString(err.Error())
			return req, err
		}
	}

	return req, nil
}

// login, require auth for user
func BindAuthUserRequest(c *fiber.Ctx) (UserRequest, error) {
	var req UserRequest
	if err := c.Bind(fiber.Map{
		"username": &req.Username,
		"password": &req.Password,
	}); err != nil {
		c.Status(400).SendString(err.Error())
		return req, err
	}
	if err := ValidateAuth(req.Username, req.Password); err != nil {
		c.Status(400).SendString(err.Error())
		return req, err
	}

	return req, nil
}
