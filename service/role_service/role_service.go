package roleservice

import (
	"context"

	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
)

type RoleService interface {
	Create(ctx context.Context, request *roleweb.RoleCreateReq) (*roleweb.RoleResponse, error)
	Update(ctx context.Context, request *roleweb.RoleUpdateReq) (*roleweb.RoleResponse, error)
	Delete(ctx context.Context, roleId string) error
	FindById(ctx context.Context, roleId string) (*roleweb.RoleResponse, error)
	FindAll(ctx context.Context) ([]*roleweb.RoleResponse, error)
}
