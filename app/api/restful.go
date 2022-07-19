package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/thvinhtruong/legoha/app/interface/restful/handler"
	"github.com/thvinhtruong/legoha/app/interface/restful/middleware"
)

func Restful() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Go Todo List",
		ServerHeader: "Fiber - Gorm",
	})

	// middleware
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.RateLimitMiddleware())

	// routes
	app.Route("/admin/", handler.AdminHandler)
	app.Route("/user/", handler.UserHandler)
	app.Route("/todo/", handler.TodoHandler)

	return app
}
