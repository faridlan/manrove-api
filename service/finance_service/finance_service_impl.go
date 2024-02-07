package financeservice

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/helper/model"
	"github.com/nostracode/mangrove-api/model/domain"
	financeweb "github.com/nostracode/mangrove-api/model/web/finance_web"
	financerepo "github.com/nostracode/mangrove-api/repository/finance_repo"
	"gorm.io/gorm"
)

type FinanceServiceImpl struct {
	FinanceRepo financerepo.FinanceRepository
	DB          *gorm.DB
	Validate    *validator.Validate
}

func NewFinanceRepository(financeRepo financerepo.FinanceRepository, db *gorm.DB, validate *validator.Validate) FinanceService {
	return &FinanceServiceImpl{
		FinanceRepo: financeRepo,
		DB:          db,
		Validate:    validate,
	}
}

func (service *FinanceServiceImpl) Create(ctx context.Context, request *financeweb.FinanceCreateReq) (*financeweb.FinanceResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	finance := &domain.Finance{
		Date:        request.Date,
		IsDebit:     request.IsDebit,
		UserId:      request.UserId,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
	}

	financeResponse, err := service.FinanceRepo.Save(ctx, tx, finance)
	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToFinanceResponse(financeResponse), nil

}

func (service *FinanceServiceImpl) Update(ctx context.Context, request *financeweb.FinanceUpdateReq) (*financeweb.FinanceResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	finance, err := service.FinanceRepo.FindById(ctx, tx, request.ID)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	finance.Date = request.Date
	finance.IsDebit = request.IsDebit
	finance.UserId = request.UserId
	finance.Description = request.Description
	finance.ImageUrl = request.ImageUrl

	financeResponse, err := service.FinanceRepo.Update(ctx, tx, finance)
	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToFinanceResponse(financeResponse), nil

}

func (service *FinanceServiceImpl) Delete(ctx context.Context, financeId string) error {

	tx := service.DB.Begin()
	defer tx.Rollback()

	finance, err := service.FinanceRepo.FindById(ctx, tx, financeId)
	if err != nil {
		return &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	err = service.FinanceRepo.Delete(ctx, tx, finance)
	if err == nil {
		tx.Commit()
	} else {
		return err
	}

	return nil
}

func (service *FinanceServiceImpl) FindById(ctx context.Context, financeId string) (*financeweb.FinanceResponse, error) {

	tx := service.DB.Begin()
	defer tx.Rollback()

	financeResponse, err := service.FinanceRepo.FindById(ctx, tx, financeId)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	return model.ToFinanceResponse(financeResponse), nil
}

func (service *FinanceServiceImpl) FindAll(ctx context.Context) ([]*financeweb.FinanceResponse, error) {

	tx := service.DB.Begin()
	defer tx.Rollback()

	financeResponses, err := service.FinanceRepo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return model.ToFinanceResponses(financeResponses), nil

}
