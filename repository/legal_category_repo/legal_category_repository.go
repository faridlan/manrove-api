package legalcategoryrepo

import (
	"context"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
)

type LegalCategoryRepository interface {
	Save(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) (*domain.LegalCategory, error)
	Update(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) (*domain.LegalCategory, error)
	Delete(ctx context.Context, db *gorm.DB, legalCategory *domain.LegalCategory) error
	FindById(ctx context.Context, db *gorm.DB, legalCategoryId string) (*domain.LegalCategory, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.LegalCategory, error)
}
