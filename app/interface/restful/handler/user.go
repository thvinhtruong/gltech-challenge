package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	userservice "github.com/thvinhtruong/legoha/app/usecase/user/service"
)

type UserHandler struct {
	service userservice.UserUseCase
}

func NewUserHandler(service userservice.UserUseCase) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
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

	user := &entity.User{Name: registerDTO.Name, Username: registerDTO.Username, Password: registerDTO.Password}
	err = u.service.RegisterUser(user.Name, user.Username, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Users Created",
	})

}

func (u *UserHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := u.service.FindAllUsers()
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

func (u *UserHandler) FindUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("userId"))
	user, err := u.service.FindUserById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Found",
		"data":    user,
	})
}

func (u *UserHandler) UpdateUserInfor(c *fiber.Ctx) error {
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
	err = u.service.PatchUserInfor(id, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Updated",
	})
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("userId"))
	err := u.service.DeleteUser(id)
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

func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
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

	user, err := u.service.LoginUser(loginDTO.Username, loginDTO.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":       "error",
			"error_detail": err,
			"error":        err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Logged in",
		"data":    user,
	})
}
