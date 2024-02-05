package userrest

import (
	"context"
	"fmt"
	"testing"

	"github.com/nostracode/mangrove-api/config/db/conn"
	"github.com/nostracode/mangrove-api/model/domain"
	userrepo "github.com/nostracode/mangrove-api/repository/user_repo"
	"github.com/stretchr/testify/assert"
)

var db = conn.NewDatabase()
var repo = userrepo.NewUserRepository()

func userTruncate() error {
	err := db.Exec("TRUNCATE users CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

func userCreate() (*domain.User, error) {

	user := &domain.User{
		Email:       "user1@mail.com",
		Name:        "user1",
		Password:    "secret1234565",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	userResponse, err := repo.Save(context.Background(), db, user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func TestCreateUserRepo(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	user := &domain.User{
		Email:       "contoh@mail.com",
		Name:        "ini_contoh_3",
		Password:    "rahasia",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	userResponse, err := repo.Save(context.Background(), db, user)
	fmt.Println(userResponse)
	assert.Nil(t, err)
	assert.Equal(t, "ini_contoh_3", userResponse.Name)

}

func TestFindAllUserRepo(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		_, err = userCreate()
		assert.Nil(t, err)
	}

	user, err := repo.FindAll(context.Background(), db)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(user))

}

func TestFindIdUserRepo(t *testing.T) {

	err := userTruncate()
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	userResponse, err := repo.FindByID(context.Background(), db, user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userResponse.Email)
	assert.Equal(t, user.Name, userResponse.Name)
	assert.Equal(t, user.PhoneNumber, userResponse.PhoneNumber)
	assert.Equal(t, user.FirstVisit, userResponse.FirstVisit)

}

func TestDeleteUserRepo(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	err = repo.Delete(context.Background(), db, user)
	assert.Nil(t, err)

	_, err = repo.FindByID(context.Background(), db, user.ID)
	assert.NotNil(t, err)

}

func TestUpdateUserRepo(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	user, err = repo.FindByID(context.Background(), db, user.ID)
	assert.Nil(t, err)

	user.Email = "email_updated"
	user.Name = "name_updated"
	user.PhoneNumber = "phone_number_updated"
	user.ImageUrl = "image_url_updated"

	userResponse, err := repo.Update(context.Background(), db, user)
	assert.Nil(t, err)

	assert.Equal(t, user.Name, userResponse.Name)
}
