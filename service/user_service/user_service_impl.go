package userservice

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/helper/model"
	"github.com/nostracode/mangrove-api/model/domain"
	userweb "github.com/nostracode/mangrove-api/model/web/user_web"
	userrepo "github.com/nostracode/mangrove-api/repository/user_repo"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo userrepo.UserRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(userRepo userrepo.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request *userweb.UserCreateReq) (*userweb.UserResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	user := &domain.User{
		Email:       request.Email,
		Name:        request.Name,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		RoleId:      request.RoleId,
		ImageUrl:    request.ImageUrl,
	}

	_, err = service.UserRepo.FindUsername(ctx, tx, user.Name)
	if err != nil {
		return nil, &exception.ConflictError{
			Message: err.Error(),
		}
	}

	_, err = service.UserRepo.FindEmail(ctx, tx, user.Email)
	if err != nil {
		return nil, &exception.ConflictError{
			Message: err.Error(),
		}
	}

	response, err := service.UserRepo.Save(ctx, tx, user)

	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToUserResponse(response), nil

}

func (service *UserServiceImpl) Update(ctx context.Context, request *userweb.UserUpdateReq) (*userweb.UserResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	userResponse, err := service.UserRepo.FindByID(ctx, tx, request.ID)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	userResponse.Email = request.Email
	userResponse.Name = request.Name
	userResponse.PhoneNumber = request.PhoneNumber
	userResponse.RoleId = request.RoleId
	userResponse.ImageUrl = request.ImageUrl

	_, err = service.UserRepo.FindUsername(ctx, tx, request.Name)
	if err != nil {
		return nil, &exception.ConflictError{
			Message: err.Error(),
		}
	}

	_, err = service.UserRepo.FindEmail(ctx, tx, request.Email)
	if err != nil {
		return nil, &exception.ConflictError{
			Message: err.Error(),
		}
	}

	response, err := service.UserRepo.Update(ctx, tx, userResponse)

	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToUserResponse(response), nil

}

func (service *UserServiceImpl) Delete(ctx context.Context, userID string) error {

	tx := service.DB.Begin()
	defer tx.Rollback()

	user, err := service.UserRepo.FindByID(ctx, tx, userID)
	if err != nil {
		return &exception.NotFoundError{
			Message: err.Error(),
		}
	}
	err = service.UserRepo.Delete(ctx, tx, user)

	if err == nil {
		tx.Commit()
	} else {
		panic(err)
	}

	return nil

}

func (service *UserServiceImpl) FindById(ctx context.Context, userID string) (*userweb.UserResponse, error) {

	user, err := service.UserRepo.FindByID(ctx, service.DB, userID)

	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	return model.ToUserResponse(user), nil

}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]*userweb.UserResponse, error) {

	users, err := service.UserRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return model.ToUserResponses(users), nil

}
