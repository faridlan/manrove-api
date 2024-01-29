// main.go
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/config/db/conn"
	rolecontroller "github.com/nostracode/mangrove-api/controller/role_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	roleservice "github.com/nostracode/mangrove-api/service/role_service"
)

func main() {

	db := conn.NewDatabase()
	validator := validator.New()

	//Role
	roleRepo := rolerepo.NewRoleRepository()
	roleService := roleservice.NewRoleService(roleRepo, db, validator)
	roleController := rolecontroller.NewRoleController(roleService)

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

	// Start the Fiber app on port 3030
	err := app.Listen(":3030")
	helper.PanicIfError(err)
}
