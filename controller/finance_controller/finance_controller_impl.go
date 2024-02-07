package financecontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/model/web"
	financeweb "github.com/nostracode/mangrove-api/model/web/finance_web"
	financeservice "github.com/nostracode/mangrove-api/service/finance_service"
)

type FinanceControllerImpl struct {
	FinanceService financeservice.FinanceService
}

func NewFinanceController(financeService financeservice.FinanceService) FinanceController {
	return &FinanceControllerImpl{
		FinanceService: financeService,
	}
}

func (controller *FinanceControllerImpl) Create(ctx *fiber.Ctx) error {

	request := new(financeweb.FinanceCreateReq)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	finance, err := controller.FinanceService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   finance,
	}

	return ctx.JSON(webResponse)
}

func (controller *FinanceControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(financeweb.FinanceUpdateReq)
	id := ctx.Params("financeId")
	request.ID = id

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	finance, err := controller.FinanceService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   finance,
	}

	return ctx.JSON(webResponse)
}

func (controller *FinanceControllerImpl) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("financeId")

	err := controller.FinanceService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webResponse)

}

func (controller *FinanceControllerImpl) FindById(ctx *fiber.Ctx) error {

	id := ctx.Params("financeId")

	financeResponse, err := controller.FinanceService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   financeResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *FinanceControllerImpl) FindAll(ctx *fiber.Ctx) error {

	financeResponses, err := controller.FinanceService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   financeResponses,
	}

	return ctx.JSON(webResponse)

}
