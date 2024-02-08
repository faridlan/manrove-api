package legalcategoryservice

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/helper/model"
	"github.com/nostracode/mangrove-api/model/domain"
	legalcategoryweb "github.com/nostracode/mangrove-api/model/web/legal-category-web"
	legalcategoryrepo "github.com/nostracode/mangrove-api/repository/legal_category_repo"
	"gorm.io/gorm"
)

type LegalCategoryServiceImpl struct {
	LegalCategoryRepo legalcategoryrepo.LegalCategoryRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewLegalCategoryService(legalCategoryRepo legalcategoryrepo.LegalCategoryRepository, db *gorm.DB, validate *validator.Validate) LegalCategoryService {
	return &LegalCategoryServiceImpl{
		LegalCategoryRepo: legalCategoryRepo,
		DB:                db,
		Validate:          validate,
	}
}

func (service *LegalCategoryServiceImpl) Create(ctx context.Context, request *legalcategoryweb.LegalCategoryCreateReq) (*legalcategoryweb.LegalCategoryResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	legalCategory := &domain.LegalCategory{
		Name: request.Name,
	}

	// _, err = service.LegalCategoryRepo.FindByName(ctx, tx, legalCategory.Name)
	// if err != nil {
	// 	return nil, &exception.ConflictError{
	// 		Message: err.Error(),
	// 	}
	// }

	response, err := service.LegalCategoryRepo.Save(ctx, tx, legalCategory)

	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	return model.ToLegalCategoryResponse(response), nil

}

func (service *LegalCategoryServiceImpl) Update(ctx context.Context, request *legalcategoryweb.LegalCategoryUpdateReq) (*legalcategoryweb.LegalCategoryResponse, error) {

	err := service.Validate.Struct(request)
	errorString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errorString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	legalCategory, err := service.LegalCategoryRepo.FindById(ctx, tx, request.ID)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	legalCategory.Name = request.Name

	legalCategory, err = service.LegalCategoryRepo.Update(ctx, tx, legalCategory)

	if err == nil {
		tx.Commit()
	} else {
		panic(err)
	}

	return model.ToLegalCategoryResponse(legalCategory), nil

}

func (service *LegalCategoryServiceImpl) Delete(ctx context.Context, legalCategoryId string) error {

	tx := service.DB.Begin()
	defer tx.Rollback()

	legalCategory, err := service.LegalCategoryRepo.FindById(ctx, tx, legalCategoryId)
	if err != nil {
		return &exception.NotFoundError{
			Message: err.Error(),
		}
	}
	err = service.LegalCategoryRepo.Delete(ctx, tx, legalCategory)

	if err == nil {
		tx.Commit()
	} else {
		panic(err)
	}

	return nil

}

func (service *LegalCategoryServiceImpl) FindById(ctx context.Context, legalCategoryId string) (*legalcategoryweb.LegalCategoryResponse, error) {

	legalCategory, err := service.LegalCategoryRepo.FindById(ctx, service.DB, legalCategoryId)

	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	return model.ToLegalCategoryResponse(legalCategory), nil

}

func (service *LegalCategoryServiceImpl) FindAll(ctx context.Context) ([]*legalcategoryweb.LegalCategoryResponse, error) {

	legalCategories, err := service.LegalCategoryRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return model.ToLegalCategoryResponses(legalCategories), nil

}
