package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/thvinhtruong/legoha/api/handler"
	"github.com/thvinhtruong/legoha/persistent"
	"github.com/thvinhtruong/legoha/repository"
	"github.com/thvinhtruong/legoha/usecase/tasklist"
	"github.com/thvinhtruong/legoha/usecase/todo"
	"github.com/thvinhtruong/legoha/usecase/user"
)

func Run(port int) {
	// making repo & create database
	repo := repository.Repository{
		DB: persistent.GetDB(),
	}

	app := fiber.New(fiber.Config{
		AppName:      "Go Todo List",
		ServerHeader: "Fiber - Gorm",
	})

	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 20,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "too much requests",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	user_repo := repository.NewUserRepository(repo.DB)
	user_service := user.NewUserService(user_repo)
	user_handler := app.Group("/user")
	handler.NewUserHandler(user_handler, user_service)

	todo_repo := repository.NewTodoRepository(repo.DB)
	todo_service := todo.NewTodoService(todo_repo)
	todo_handler := app.Group("/todo")
	handler.NewTodoHandler(todo_handler, todo_service)

	tasklist_repo := repository.NewTaskListRepository(repo.DB)
	tasklist_service := tasklist.NewTaskService(tasklist_repo)
	tasklist_handler := app.Group("/")
	handler.NewAdminHandler(tasklist_handler, tasklist_service)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 3000.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
