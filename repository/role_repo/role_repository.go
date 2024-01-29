package rolerepo

import (
	"context"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error)
	Update(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error)
	Delete(ctx context.Context, db *gorm.DB, role *domain.Role) error
	FindByUID(ctx context.Context, db *gorm.DB, roleId string) (*domain.Role, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Role, error)
	FindByName(ctx context.Context, db *gorm.DB, name string) (*domain.Role, error)
}
