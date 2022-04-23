package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entity "github.com/thvinhtruong/legoha/entities"
	"github.com/thvinhtruong/legoha/usecase/user"
)

func NewUserHandler(app fiber.Router, service user.UserUseCase) {
	app.Post("/", createUser(service))
	app.Post("/login", loginUser(service))
	app.Get("/all", listUsers(service))
	app.Get("/:userId", getUser(service))
	app.Patch("/:userId", patchUser(service))
	app.Delete("/:userId", deleteUser(service))
}

func createUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		type RegisterDTO struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var registerDTO RegisterDTO
		err := c.BodyParser(&registerDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		err = service.CreateUser(registerDTO.Name, registerDTO.Username, registerDTO.Password)
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

func listUsers(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.ListUsers()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    users,
		})
	}
}

func getUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("userId"))
		user, err := service.GetUserById(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    user,
		})
	}
}

func patchUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type UpdateDTO struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var updateDTO UpdateDTO
		err := c.BodyParser(&updateDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}
		id, _ := strconv.Atoi(c.Params("userId"))
		user := &entity.User{Name: updateDTO.Name, Username: updateDTO.Username, Password: updateDTO.Password}
		err = service.PatchUser(id, user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Updated",
		})
	}
}

func deleteUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("userId"))
		err := service.DeleteUser(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Deleted",
		})
	}
}

func loginUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginDTO struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var loginDTO LoginDTO
		err := c.BodyParser(&loginDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user, err := service.LoginUser(loginDTO.Username, loginDTO.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    user,
		})
	}
}
