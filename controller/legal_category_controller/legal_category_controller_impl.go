package legalcategorycontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/model/web"
	legalcategoryweb "github.com/nostracode/mangrove-api/model/web/legal-category-web"
	LegalCategoryService "github.com/nostracode/mangrove-api/service/legal_category_service"
)

type LegalCategoryControllerImpl struct {
	LegalCategoryService LegalCategoryService.LegalCategoryService
}

func NewLegalCategoryController(legalCategoryService LegalCategoryService.LegalCategoryService) LegalCategoryController {
	return &LegalCategoryControllerImpl{
		LegalCategoryService: legalCategoryService,
	}
}

func (controller *LegalCategoryControllerImpl) Create(ctx *fiber.Ctx) error {

	request := new(legalcategoryweb.LegalCategoryCreateReq)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	legalCategory, err := controller.LegalCategoryService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   legalCategory,
	}

	return ctx.JSON(webResponse)

}

func (controller *LegalCategoryControllerImpl) Update(ctx *fiber.Ctx) error {

	request := new(legalcategoryweb.LegalCategoryUpdateReq)

	id := ctx.Params("legalCategoryId")
	request.ID = id

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	legalCategory, err := controller.LegalCategoryService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   legalCategory,
	}

	return ctx.JSON(webReponse)
}

func (controller *LegalCategoryControllerImpl) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("legalCategoryId")

	err := controller.LegalCategoryService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webReponse)

}

func (controller *LegalCategoryControllerImpl) FindById(ctx *fiber.Ctx) error {

	id := ctx.Params("legalCategoryId")

	legalCategory, err := controller.LegalCategoryService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   legalCategory,
	}

	return ctx.JSON(webReponse)

}

func (controller *LegalCategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {

	legalCategorys, err := controller.LegalCategoryService.FindAll(ctx.Context())
	if err != nil {
		return err
	}
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   legalCategorys,
	}

	return ctx.JSON(webReponse)

}
