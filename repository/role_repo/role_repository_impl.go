package rolerepo

import (
	"context"
	"errors"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error) {

	// err := db.Omit("UID").Create(&role).Error
	err := db.Omit("UID").Clauses(clause.Returning{}).Select("name").Create(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil

}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error) {

	err := db.Save(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil

}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, role *domain.Role) error {

	err := db.Delete(&role).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *RoleRepositoryImpl) FindByUID(ctx context.Context, db *gorm.DB, roleId string) (*domain.Role, error) {

	role := domain.Role{}
	err := db.First(&role, "uid = ?", roleId).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &role, nil

}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Role, error) {

	role := []*domain.Role{}
	err := db.Find(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil

}