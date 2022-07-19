package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/app/interface/restful/presenter"
	"github.com/thvinhtruong/legoha/app/registry"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
	"github.com/thvinhtruong/legoha/pkg/sqlconn"
)

func getAdminAccess(db *sql.DB) registry.AdminHook {
	return registry.BuildAdminHook(db)
}

func AdminHandler(app fiber.Router) {
	app.Post("/assign?userId=?&todoId=?", assignUserToTodo)
	app.Delete("/revoke?userId=?&todoId=?", revokeUserFromTodo)
	app.Patch("/user/:userId/todo/:todoid/done", completedUserTodo)
	app.Patch("/user/:userId/todo/:todoid/undo", undoUserTodo)
	app.Delete("/todo/delete?todoId=?", deleteTodo)
	app.Delete("/user/delete?userId=?", deleteUser)
}

func assignUserToTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	var record dto.TodoUser
	record.UserID = userId
	record.TodoID = todoId

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.AssignTask(ctx, record)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func revokeUserFromTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.RevokeTask(ctx, userId, todoId)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func completedUserTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}
	todoId, err := strconv.Atoi(c.Params("todoid"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.CompletedTask(ctx, userId, todoId)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func undoUserTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}
	todoId, err := strconv.Atoi(c.Params("todoid"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.UndoTask(ctx, userId, todoId)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func deleteTodo(c *fiber.Ctx) error {
	ctx := context.Background()
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.DeleteTodo(ctx, todoId)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func deleteUser(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	admin_access := getAdminAccess(db)
	err = admin_access.AdminService.DeleteUser(ctx, userId)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}
