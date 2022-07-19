package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/thvinhtruong/legoha/app/interface/restful/presenter"
	"github.com/thvinhtruong/legoha/app/registry"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
	"github.com/thvinhtruong/legoha/pkg/hasher"
	"github.com/thvinhtruong/legoha/pkg/sqlconn"
)

func getUserAcess(db *sql.DB) registry.UserHook {
	return registry.BuildUserHook(db)
}

func UserHandler(app fiber.Router) {
	app.Post("/register", registerUser)
	app.Get("/:userId", getUserByID)
	app.Patch("/:userId", updateUser)
	app.Patch("/password/:userId", changePassword)
}

func registerUser(c *fiber.Ctx) error {
	ctx := context.Background()
	data, err := presenter.BindUserRequest(c, false)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	var user_record dto.User
	err = copier.Copy(user_record, data)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	password, err := hasher.HashPassword(user_record.Password)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	user_record.Password = password

	db := sqlconn.DB
	user_access := getUserAcess(db)
	err = user_access.UserService.RegisterUser(ctx, user_record)
	if err != nil {
		log.Println("Error: ", err)
		return presenter.Response(c, http.StatusOK, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func getUserByID(c *fiber.Ctx) error {
	ctx := context.Background()
	userID, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	db := sqlconn.DB
	user_access := getUserAcess(db)
	user, err := user_access.UserService.FindUserByID(ctx, userID)
	if err != nil {
		return presenter.Response(c, http.StatusOK, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, user)
}

func updateUser(c *fiber.Ctx) error {
	ctx := context.Background()
	userID, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	data, err := presenter.BindUserRequest(c, true)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	// only update name and username
	user_record := dto.User{
		Name:     data.Name,
		Username: data.Username,
	}

	db := sqlconn.DB
	user_access := getUserAcess(db)
	err = user_access.UserService.PatchUserInfor(ctx, userID, user_record)
	if err != nil {
		return presenter.Response(c, http.StatusOK, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}

func changePassword(c *fiber.Ctx) error {
	ctx := context.Background()
	userID, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	data, err := presenter.BindUserRequest(c, true)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	var user_record dto.User
	user_record.Password = data.Password
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	new_password, err := hasher.HashPassword(user_record.Password)
	if err != nil {
		return presenter.Response(c, http.StatusInternalServerError, err, nil)
	}

	user_record.Password = new_password

	db := sqlconn.DB
	user_access := getUserAcess(db)
	err = user_access.UserService.PatchUserInfor(ctx, userID, user_record)
	if err != nil {
		return presenter.Response(c, http.StatusOK, err, nil)
	}

	return presenter.Response(c, http.StatusOK, nil, nil)
}
