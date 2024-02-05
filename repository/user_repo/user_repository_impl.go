package userrepo

import (
	"context"
	"errors"

	"github.com/nostracode/mangrove-api/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error) {

	err := db.Omit("ID", "FirstVisit").Clauses(clause.Returning{}).Select("email", "name", "password", "phone_number", "role_id", "image_url").Create(&user).Error
	// err := db.Omit("ID", "FirstVisit").Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (repository *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error) {

	err := db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, user *domain.User) error {

	err := db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, userID string) (*domain.User, error) {

	user := domain.User{}
	err := db.First(&user, "ID = ?", userID).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil

}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.User, error) {

	users := []*domain.User{}
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (repository *UserRepositoryImpl) FindUsername(ctx context.Context, db *gorm.DB, username string) (*domain.User, error) {

	user := domain.User{}
	err := db.First(&user, "name = ?", username).Error
	if err == nil {
		return nil, errors.New("username has been taken")
	}

	return &user, nil

}

func (repository *UserRepositoryImpl) FindEmail(ctx context.Context, db *gorm.DB, email string) (*domain.User, error) {

	user := domain.User{}
	err := db.First(&user, "email = ?", email).Error
	if err == nil {
		return nil, errors.New("email has been taken")
	}

	return &user, nil

}

func (repository *UserRepositoryImpl) FindUsernameId(ctx context.Context, db *gorm.DB, username string, userId string) (*domain.User, error) {

	user := domain.User{}
	err := db.Not("id = ?", userId).First(&user, "name = ?", username).Error
	// err := db.Where(domain.User{
	// 	ID:   userId,
	// 	Name: username,
	// }).First(&user).Error

	if err == nil {
		return nil, errors.New("username has been taken")
	}

	return &user, nil

}

func (repository *UserRepositoryImpl) FindEmailId(ctx context.Context, db *gorm.DB, email string, userId string) (*domain.User, error) {

	user := domain.User{}
	err := db.Not("id = ?", userId).First(&user, "email = ?", email).Error
	// err := db.Where(domain.User{
	// 	ID:    userId,
	// 	Email: email,
	// }).First(&user).Error

	if err == nil {
		return nil, errors.New("email has been taken")
	}

	return &user, nil

}
