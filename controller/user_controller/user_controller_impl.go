package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/model/web"
	userweb "github.com/nostracode/mangrove-api/model/web/user_web"
	userservice "github.com/nostracode/mangrove-api/service/user_service"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {

	request := new(userweb.UserCreateReq)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	user, err := controller.UserService.Register(ctx.Context(), request)
	if err != nil {
		return err
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	return ctx.JSON(webResponse)

}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {

	request := new(userweb.UserUpdateReq)
	id := ctx.Params("userId")
	request.ID = id

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	user, err := controller.UserService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	return ctx.JSON(webReponse)

}

func (controller *UserControllerImpl) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("userId")

	err := controller.UserService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webReponse)

}

func (controller *UserControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")

	user, err := controller.UserService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	return ctx.JSON(webReponse)

}

func (controller *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {

	users, err := controller.UserService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	return ctx.JSON(webReponse)

}
