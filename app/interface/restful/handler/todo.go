package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/thvinhtruong/legoha/app/apperror"
	"github.com/thvinhtruong/legoha/app/interface/restful/presenter"
	"github.com/thvinhtruong/legoha/app/registry"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
	"github.com/thvinhtruong/legoha/pkg/sqlconn"
)

func getTodoAccess(db *sql.DB) registry.TodoHook {
	return registry.BuildTodoHook(db)
}

func TodoHandler(app fiber.Router) {
	app.Get("/:todoId", getTodoByID)
	app.Get("/search/todo?title=?&description=?", searchTodo)
	app.Post("/todo", createTodo)
	app.Patch("/:todoId", updateTodo)
}

func getTodoByID(c *fiber.Ctx) error {
	ctx := context.Background()
	todo_access := getTodoAccess(sqlconn.DB)
	todo_id, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	todo, err := todo_access.TodoService.ShowTodoByID(ctx, todo_id)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, todo)
}

func searchTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	todo_access := getTodoAccess(sqlconn.DB)
	title := c.Query("title")
	description := c.Query("description")

	if len(title) == 0 || len(description) == 0 {
		return presenter.Response(c, http.StatusOK, apperror.ErrorInputInvalid, nil)
	}

	var todo_record dto.Todo
	if title != "" || description != "" {
		todo_record.Title = title
		todo_record.Description = description
	}

	todo, err := todo_access.TodoService.FindTodos(ctx, todo_record)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, todo)
}

func createTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	todo_access := getTodoAccess(sqlconn.DB)
	data, err := presenter.BindTodoRequest(c)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	var todo_record dto.Todo
	err = copier.Copy(todo_record, data)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	err = todo_access.TodoService.PostNewTodo(ctx, todo_record)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func updateTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	todo_access := getTodoAccess(sqlconn.DB)
	todo_id, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	data, err := presenter.BindTodoRequest(c)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	var todo_record dto.Todo
	err = copier.Copy(todo_record, data)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	err = todo_access.TodoService.UpdateTodoInfor(ctx, todo_id, todo_record)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}
