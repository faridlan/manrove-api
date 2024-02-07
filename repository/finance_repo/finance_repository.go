package financerepo

import (
	"context"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
)

type FinanceRepository interface {
	Save(ctx context.Context, db *gorm.DB, finance *domain.Finance) (*domain.Finance, error)
	Update(ctx context.Context, db *gorm.DB, finance *domain.Finance) (*domain.Finance, error)
	Delete(ctx context.Context, db *gorm.DB, finance *domain.Finance) error
	FindById(ctx context.Context, db *gorm.DB, financeId string) (*domain.Finance, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Finance, error)
}
