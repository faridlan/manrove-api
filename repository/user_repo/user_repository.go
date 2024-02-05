package userrepo

import (
	"context"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, db *gorm.DB, user *domain.User) error
	FindByID(ctx context.Context, db *gorm.DB, userID string) (*domain.User, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.User, error)
	FindUsername(ctx context.Context, db *gorm.DB, username string) (*domain.User, error)
	FindEmail(ctx context.Context, db *gorm.DB, email string) (*domain.User, error)
	FindUsernameId(ctx context.Context, db *gorm.DB, username string, userId string) (*domain.User, error)
	FindEmailId(ctx context.Context, db *gorm.DB, email string, userId string) (*domain.User, error)
}
