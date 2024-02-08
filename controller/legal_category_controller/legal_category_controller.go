package legalcategorycontroller

import "github.com/gofiber/fiber/v2"

type LegalCategoryController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
