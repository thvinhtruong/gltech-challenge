package registry

import (
	"database/sql"

	"github.com/thvinhtruong/legoha/app/usecase"
	"github.com/thvinhtruong/legoha/app/usecase/interactor"
)

// hook creates a bridge between data layer to usecase and domain layer
type UserHook struct {
	UserService usecase.UserService
}

// this func initialize data layer and usecase layer at the same time
func BuildUserHook(db *sql.DB) UserHook {
	repo := GetRepository(db)
	userusecase := interactor.NewUserUseCase(&repo)
	return UserHook{UserService: userusecase}
}

type TodoHook struct {
	TodoService usecase.TodoService
}

// this func initializes todo repository and usecase
func BuildTodoHook(db *sql.DB) TodoHook {
	repo := GetRepository(db)
	todousecase := interactor.NewTodoUseCase(&repo)
	return TodoHook{TodoService: todousecase}
}

type AdminHook struct {
	AdminService usecase.AdminService
}

// this func initializes todouser repository and usecase
func BuildAdminHook(db *sql.DB) AdminHook {
	repo := GetRepository(db)
	adminusecase := interactor.NewAdminUseCase(&repo)
	return AdminHook{AdminService: adminusecase}
}
