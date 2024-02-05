// main.go
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/config/db/conn"
	rolecontroller "github.com/nostracode/mangrove-api/controller/role_controller"
	usercontroller "github.com/nostracode/mangrove-api/controller/user_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	userrepo "github.com/nostracode/mangrove-api/repository/user_repo"
	roleservice "github.com/nostracode/mangrove-api/service/role_service"
	userservice "github.com/nostracode/mangrove-api/service/user_service"
)

func main() {

	db := conn.NewDatabase()
	validator := validator.New()

	//Role
	roleRepo := rolerepo.NewRoleRepository()
	roleService := roleservice.NewRoleService(roleRepo, db, validator)
	roleController := rolecontroller.NewRoleController(roleService)

	//User
	userRepo := userrepo.NewUserRepository()
	userService := userservice.NewUserService(userRepo, db, validator)
	userController := usercontroller.NewUserController(userService)

	// Create a new Fiber instance
	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ExceptionError,
		},
	)

	//Endpoint Of Role
	app.Post("/api/roles", roleController.Create)
	app.Get("/api/roles", roleController.FindAll)
	app.Get("/api/roles/:roleId", roleController.FindById)
	app.Put("/api/roles/:roleId", roleController.Update)
	app.Delete("/api/roles/:roleId", roleController.Delete)

	//Endpoint Of User
	app.Post("/api/users", userController.Register)
	app.Get("/api/users", userController.FindAll)
	app.Get("/api/users/:userId", userController.FindById)
	app.Put("/api/users/:userId", userController.Update)
	app.Delete("/api/users/:userId", userController.Delete)

	// Start the Fiber app on port 3030
	err := app.Listen(":3030")
	helper.PanicIfError(err)
}
