package main

import (
	"fmt"
	"log"

	"github.com/thvinhtruong/legoha/cmd/core"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/thvinhtruong/legoha/app/api"
	repository "github.com/thvinhtruong/legoha/app/interface/persistence/rdbms/gormdb"
	"github.com/thvinhtruong/legoha/app/interface/restful/handler"
	tasklistservice "github.com/thvinhtruong/legoha/app/usecase/tasklist/service"
	todoservice "github.com/thvinhtruong/legoha/app/usecase/todo/service"
	userservice "github.com/thvinhtruong/legoha/app/usecase/user/service"
)

func userStart(app fiber.Router, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	service := userservice.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	api.NewUserRoutes(app, handler)
}

func todoStart(app fiber.Router, db *gorm.DB) {
	repo := repository.NewTodoRepository(db)
	service := todoservice.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)
	api.NewTodoRoutes(app, handler)
}

func adminStart(app fiber.Router, db *gorm.DB) {
	repo := repository.NewTaskListRepository(db)
	service := tasklistservice.NewTaskService(repo)
	handler := handler.NewAdminHandler(service)
	api.NewAdminRoute(app, handler)
}

func Run(port int) {
	// connect to db
	db := core.GetDB()
	defer db.Close()

	// create app
	app := api.Restful()

	// start api
	userStart(app.Group("/user"), db)
	todoStart(app.Group("/todo"), db)
	adminStart(app.Group("/"), db)

	// Listen to port 3000.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

func main() {
	Run(3000)
}
