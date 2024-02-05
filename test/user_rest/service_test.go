package userrest

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	userweb "github.com/nostracode/mangrove-api/model/web/user_web"
	userservice "github.com/nostracode/mangrove-api/service/user_service"
	"github.com/stretchr/testify/assert"
)

var validate = validator.New()
var service = userservice.NewUserService(repo, db, validate)

func TestFindAllUserService(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		_, err = userCreate()
		assert.Nil(t, err)
	}

	userResponse, err := service.FindAll(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 2, len(userResponse))
}

func TestFindByIDUserServiceSuccess(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	userResponse, err := service.FindById(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userResponse.Email)
	assert.Equal(t, user.Name, userResponse.Name)
	assert.Equal(t, user.PhoneNumber, userResponse.PhoneNumber)
	assert.Equal(t, user.FirstVisit, userResponse.FirstVisit)
}

func TestFindByIDUserServiceNotFound(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	userResponse, err := service.FindById(context.Background(), "salah")
	assert.Nil(t, userResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestRegisterServiceSuccess(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	user := &userweb.UserCreateReq{
		Email:       "user1@mail.com",
		Name:        "user1",
		Password:    "secret1",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	userResponse, err := service.Register(context.Background(), user)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userResponse.Email)
	assert.Equal(t, user.Name, userResponse.Name)
	assert.Equal(t, user.PhoneNumber, userResponse.PhoneNumber)
}

func TestRegisterServiceUsernameConflict(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	_, err = userCreate()
	assert.Nil(t, err)

	userCreate := &userweb.UserCreateReq{
		Email:       "user2@mail.com",
		Name:        "user1",
		Password:    "secret12345656",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	userResponse, err := service.Register(context.Background(), userCreate)
	assert.NotNil(t, err)
	assert.Nil(t, userResponse)
	assert.Equal(t, "username has been taken", err.Error())
}

func TestRegisterServiceEmailConflict(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	_, err = userCreate()
	assert.Nil(t, err)

	userCreate := &userweb.UserCreateReq{
		Email:       "user1@mail.com",
		Name:        "user2",
		Password:    "secret1",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	userResponse, err := service.Register(context.Background(), userCreate)
	assert.NotNil(t, err)
	assert.Nil(t, userResponse)
	assert.Equal(t, "email has been taken", err.Error())
}

func TestRegisterServiceEmailFailed(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	userCreate := &userweb.UserCreateReq{
		Email:       "",
		Name:        "user2",
		Password:    "secret1234565",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	_, err = service.Register(context.Background(), userCreate)
	assert.NotNil(t, err)
	assert.Equal(t, "Email is a required field", err.Error())
}

func TestUpdateServiceSuccess(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)
	user, err := userCreate()
	assert.Nil(t, err)

	userUpdate := &userweb.UserUpdateReq{
		ID:          user.ID,
		Email:       "email_updated@mail.com",
		Name:        "user1",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default_updated.png",
	}

	userResponse, err := service.Update(context.Background(), userUpdate)
	assert.Nil(t, err)
	assert.Equal(t, userUpdate.Email, userResponse.Email)
	assert.Equal(t, userUpdate.Name, userResponse.Name)
	assert.Equal(t, userUpdate.PhoneNumber, userResponse.PhoneNumber)
}

func TestUpdateServiceNotFound(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	_, err = userCreate()
	assert.Nil(t, err)

	userUpdate := &userweb.UserUpdateReq{
		ID:          "salah",
		Email:       "user2_updated@mail.com",
		Name:        "user1",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default_updated.png",
	}

	userResponse, err := service.Update(context.Background(), userUpdate)
	assert.NotNil(t, err)
	assert.Nil(t, userResponse)
	assert.Equal(t, "user not found", err.Error())
}

func TestUpdateServiceUsernameConflict(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	userCreateReq := &userweb.UserCreateReq{
		Email:       "user1@mail.com",
		Name:        "user5",
		Password:    "secret17832438",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	_, err = service.Register(context.Background(), userCreateReq)
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	userUpdate := &userweb.UserUpdateReq{
		ID:          user.ID,
		Email:       "user2_updated@mail.com",
		Name:        "user5",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default_updated.png",
	}

	userResponse, err := service.Update(context.Background(), userUpdate)
	assert.NotNil(t, err)
	assert.Nil(t, userResponse)
	assert.Equal(t, "username has been taken", err.Error())
}

func TestUpdateServiceEmailConflict(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)

	userCreateReq := &userweb.UserCreateReq{
		Email:       "user5@mail.com",
		Name:        "user5",
		Password:    "secret17832438",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default.png",
	}

	_, err = service.Register(context.Background(), userCreateReq)
	assert.Nil(t, err)

	user, err := userCreate()
	assert.Nil(t, err)

	userUpdate := &userweb.UserUpdateReq{
		ID:          user.ID,
		Email:       "user5@mail.com",
		Name:        "user1_updated",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default_updated.png",
	}

	userResponse, err := service.Update(context.Background(), userUpdate)
	assert.NotNil(t, err)
	assert.Nil(t, userResponse)
	assert.Equal(t, "email has been taken", err.Error())
}

func TestUpdateServiceEmailFailed(t *testing.T) {
	err := userTruncate()
	assert.Nil(t, err)
	user, err := userCreate()
	assert.Nil(t, err)

	userUpdate := &userweb.UserUpdateReq{
		ID:          user.ID,
		Email:       "",
		Name:        "user2_update",
		PhoneNumber: "33300998822",
		RoleId:      "a97e962c9f9e458ebe05e4d7b60a70f3",
		ImageUrl:    "https://image/default_update.png",
	}

	_, err = service.Update(context.Background(), userUpdate)
	assert.NotNil(t, err)
	assert.Equal(t, "Email is a required field", err.Error())
}
