package financerepo

import (
	"context"
	"errors"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FinanceRepositoryImpl struct {
}

func NewFinanceRepository() FinanceRepository {
	return &FinanceRepositoryImpl{}
}

func (repository *FinanceRepositoryImpl) Save(ctx context.Context, db *gorm.DB, finance *domain.Finance) (*domain.Finance, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("date", "is_debit", "user_id", "description", "image_url").Create(&finance).Error
	if err != nil {
		return nil, err
	}

	return finance, nil

}

func (repository *FinanceRepositoryImpl) Update(ctx context.Context, db *gorm.DB, finance *domain.Finance) (*domain.Finance, error) {

	err := db.Save(&finance).Error
	if err != nil {
		return nil, err
	}

	return finance, nil

}

func (repository *FinanceRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, finance *domain.Finance) error {

	err := db.Delete(&finance).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *FinanceRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, financeId string) (*domain.Finance, error) {

	finance := domain.Finance{}
	err := db.First(&finance, "ID = ?", financeId).Error
	if err != nil {
		return nil, errors.New("finance not found")
	}

	return &finance, nil

}

func (repository *FinanceRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Finance, error) {

	finances := []*domain.Finance{}
	err := db.Find(&finances).Error
	if err != nil {
		return nil, err
	}

	return finances, nil

}
