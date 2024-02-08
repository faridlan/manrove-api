package legalcategoryrepo

import (
	"context"
	"errors"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LegalCategoryRepositoryImpl struct {
}

func NewLegalCategoryRepository() LegalCategoryRepository {
	return &LegalCategoryRepositoryImpl{}
}

func (repository *LegalCategoryRepositoryImpl) Save(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) (*domain.LegalCategory, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("name").Create(&legalCategory).Error
	if err != nil {
		return nil, err
	}

	return legalCategory, nil

}

func (repository *LegalCategoryRepositoryImpl) Update(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) (*domain.LegalCategory, error) {

	err := db.Save(&legalCategory).Error
	if err != nil {
		return nil, err
	}

	return legalCategory, nil

}

func (repository *LegalCategoryRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) error {

	err := db.Delete(&legalCategory).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *LegalCategoryRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, legalCategoryId string) (*domain.LegalCategory, error) {

	legalCategory := domain.LegalCategory{}
	err := db.First(&legalCategory, "ID = ?", legalCategoryId).Error
	if err != nil {
		return nil, errors.New("legal category not found")
	}

	return &legalCategory, nil

}

func (repository *LegalCategoryRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.LegalCategory, error) {

	legalCategory := []*domain.LegalCategory{}
	err := db.Find(&legalCategory).Error
	if err != nil {
		return nil, err
	}

	return legalCategory, nil

}
