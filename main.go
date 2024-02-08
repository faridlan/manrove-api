// main.go
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/config/db/conn"
	financecontroller "github.com/nostracode/mangrove-api/controller/finance_controller"
	legalcategorycontroller "github.com/nostracode/mangrove-api/controller/legal_category_controller"
	rolecontroller "github.com/nostracode/mangrove-api/controller/role_controller"
	usercontroller "github.com/nostracode/mangrove-api/controller/user_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	financerepo "github.com/nostracode/mangrove-api/repository/finance_repo"
	legalcategoryrepo "github.com/nostracode/mangrove-api/repository/legal_category_repo"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	userrepo "github.com/nostracode/mangrove-api/repository/user_repo"
	financeservice "github.com/nostracode/mangrove-api/service/finance_service"
	legalcategoryservice "github.com/nostracode/mangrove-api/service/legal_category_service"
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

	//Finance
	financeRepo := financerepo.NewFinanceRepository()
	financeService := financeservice.NewFinanceRepository(financeRepo, db, validator)
	financeController := financecontroller.NewFinanceController(financeService)

	//Legal Category
	legalCategoryRepo := legalcategoryrepo.NewLegalCategoryRepository()
	legalCategoryService := legalcategoryservice.NewLegalCategoryService(legalCategoryRepo, db, validator)
	legalCategoryController := legalcategorycontroller.NewLegalCategoryController(legalCategoryService)

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

	//Endpoint Of Finance
	app.Post("/api/finances", financeController.Create)
	app.Get("/api/finances", financeController.FindAll)
	app.Get("/api/finances/:financeId", financeController.FindById)
	app.Put("/api/finances/:financeId", financeController.Update)
	app.Delete("/api/finances/:financeId", financeController.Delete)

	//Endpoint Of Legal Category
	app.Post("/api/legal/categories", legalCategoryController.Create)
	app.Get("/api/legal/categories", legalCategoryController.FindAll)
	app.Get("/api/legal/categories/:legalCategoryId", legalCategoryController.FindById)
	app.Put("/api/legal/categories/:legalCategoryId", legalCategoryController.Update)
	app.Delete("/api/legal/categories/:legalCategoryId", legalCategoryController.Delete)

	// Start the Fiber app on port 3030
	err := app.Listen(":3030")
	helper.PanicIfError(err)
}
