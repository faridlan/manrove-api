package roleservice

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/helper/model"
	"github.com/nostracode/mangrove-api/model/domain"
	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	"gorm.io/gorm"
)

type RoleServiceImpl struct {
	RoleRepo rolerepo.RoleRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewRoleService(roleRepo rolerepo.RoleRepository, db *gorm.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepo: roleRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request *roleweb.RoleCreateReq) (*roleweb.RoleResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	role := &domain.Role{
		Name: request.Name,
	}

	_, err = service.RoleRepo.FindByName(ctx, tx, role.Name)
	if err != nil {
		return nil, &exception.ConflictError{
			Message: err.Error(),
		}
	}

	response, err := service.RoleRepo.Save(ctx, tx, role)

	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToRoleResponse(response), nil

}

func (service *RoleServiceImpl) Update(ctx context.Context, request *roleweb.RoleUpdateReq) (*roleweb.RoleResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	role, err := service.RoleRepo.FindByID(ctx, tx, request.ID)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	role.Name = request.Name

	role, err = service.RoleRepo.Update(ctx, tx, role)

	if err == nil {
		tx.Commit()
	} else {
		panic(err)
	}

	return model.ToRoleResponse(role), nil

}

func (service *RoleServiceImpl) Delete(ctx context.Context, roleId string) error {

	tx := service.DB.Begin()
	defer tx.Rollback()

	role, err := service.RoleRepo.FindByID(ctx, tx, roleId)
	if err != nil {
		return &exception.NotFoundError{
			Message: err.Error(),
		}
	}
	err = service.RoleRepo.Delete(ctx, tx, role)

	if err == nil {
		tx.Commit()
	} else {
		panic(err)
	}

	return nil

}

func (service *RoleServiceImpl) FindById(ctx context.Context, roleId string) (*roleweb.RoleResponse, error) {

	role, err := service.RoleRepo.FindByID(ctx, service.DB, roleId)

	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	return model.ToRoleResponse(role), nil

}

func (service *RoleServiceImpl) FindAll(ctx context.Context) ([]*roleweb.RoleResponse, error) {

	roles, err := service.RoleRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return model.ToRoleResponses(roles), nil

}
