package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/model/web"
)

func ExceptionError(ctx *fiber.Ctx, err error) error {

	if errMessage, ok := err.(*NotFoundError); ok {
		return notFoundError(ctx, errMessage.Error())
	} else if badRequest, ok := err.(*BadRequestError); ok {
		return badRequestError(ctx, badRequest.Error())
	} else if conflict, ok := err.(*ConflictError); ok {
		return conflictError(ctx, conflict.Error())
	} else {
		return internalServerError(ctx, err)
	}

}

func internalServerError(ctx *fiber.Ctx, err error) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	}

	return ctx.JSON(webResponse)

}

func notFoundError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusNotFound)
	webResponse := web.WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT FOUND",
		Data:   err,
	}

	return ctx.JSON(webResponse)

}

func badRequestError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusBadRequest)
	webResponse := web.WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "BAD REQUEST",
		Data:   err,
	}

	return ctx.JSON(webResponse)

}

func conflictError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusConflict)
	webResponse := web.WebResponse{
		Code:   fiber.StatusConflict,
		Status: "CONFLICT",
		Data:   err,
	}

	return ctx.JSON(webResponse)

}
