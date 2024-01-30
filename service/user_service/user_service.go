package userservice

import (
	"context"

	userweb "github.com/nostracode/mangrove-api/model/web/user_web"
)

type UserService interface {
	Register(ctx context.Context, request *userweb.UserCreateReq) (*userweb.UserResponse, error)
	Update(ctx context.Context, request *userweb.UserUpdateReq) (*userweb.UserResponse, error)
	Delete(ctx context.Context, userID string) error
	FindById(ctx context.Context, userID string) (*userweb.UserResponse, error)
	FindAll(ctx context.Context) ([]*userweb.UserResponse, error)
}
