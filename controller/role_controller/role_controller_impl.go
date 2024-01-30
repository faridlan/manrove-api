package rolecontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nostracode/mangrove-api/model/web"
	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
	roleservice "github.com/nostracode/mangrove-api/service/role_service"
)

type RoleControllerImpl struct {
	RoleService roleservice.RoleService
}

func NewRoleController(roleService roleservice.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Create(ctx *fiber.Ctx) error {

	request := new(roleweb.RoleCreateReq)
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	role, err := controller.RoleService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   role,
	}

	return ctx.JSON(webResponse)

}

func (controller *RoleControllerImpl) Update(ctx *fiber.Ctx) error {

	request := new(roleweb.RoleUpdateReq)
	id := ctx.Params("roleId")
	request.ID = id

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	role, err := controller.RoleService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   role,
	}

	return ctx.JSON(webReponse)

}

func (controller *RoleControllerImpl) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("roleId")

	err := controller.RoleService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webReponse)

}

func (controller *RoleControllerImpl) FindById(ctx *fiber.Ctx) error {

	id := ctx.Params("roleId")

	user, err := controller.RoleService.FindById(ctx.Context(), id)
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

func (controller *RoleControllerImpl) FindAll(ctx *fiber.Ctx) error {

	users, err := controller.RoleService.FindAll(ctx.Context())
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
