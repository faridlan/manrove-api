package role

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/model/domain"
	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
	roleservice "github.com/nostracode/mangrove-api/service/role_service"
	"github.com/stretchr/testify/assert"
)

var validate = validator.New()
var service = roleservice.NewRoleService(repo, db, validate)

func TestFindAllRoleService(t *testing.T) {

	Truncate()
	Create()
	roles, err := service.FindAll(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, 1, len(roles))

}

func TestFindByIdRoleService(t *testing.T) {

	Truncate()
	roleResponse, _ := Create()
	role, err := service.FindById(context.Background(), roleResponse.ID)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", role.Name)

}

func TestFindAllRoleServiceFailed(t *testing.T) {
	_, err := service.FindById(context.Background(), "salah")
	assert.NotNil(t, err)
	assert.Equal(t, "role not found", err.Error())
}

func TestCreateRoleService(t *testing.T) {
	Truncate()
	role := &roleweb.RoleCreateReq{
		Name: "super_admin",
	}

	roleResponse, err := service.Create(context.Background(), role)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", roleResponse.Name)
}

func TestCreateRoleServiceConflict(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	user := &roleweb.RoleCreateReq{
		Name: roleResponse.Name,
	}

	_, err := service.Create(context.Background(), user)
	assert.NotNil(t, err)
	assert.Equal(t, "role name already create", err.Error())
}

func TestCreateRoleServiceFailed(t *testing.T) {

	role := &roleweb.RoleCreateReq{
		Name: "",
	}

	_, err := service.Create(context.Background(), role)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is a required field", err.Error())

}

func TestUpdateRoleService(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()

	role := &roleweb.RoleUpdateReq{
		ID:   roleResponse.ID,
		Name: "super_admin_updated",
	}

	response, err := service.Update(context.Background(), role)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin_updated", response.Name)
}

func TestUpdateRoleServiceFailed(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()

	role := &roleweb.RoleUpdateReq{
		ID:   roleResponse.ID,
		Name: "",
	}

	_, err := service.Update(context.Background(), role)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is a required field", err.Error())
}

func TestUpdateRoleServiceNorFound(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()

	role := &roleweb.RoleUpdateReq{
		ID:   "salah",
		Name: roleResponse.Name,
	}

	_, err := service.Update(context.Background(), role)
	assert.NotNil(t, err)
	assert.Equal(t, "role not found", err.Error())
}

func TestDeleteRoleService(t *testing.T) {

	Truncate()
	roleResponse, _ := Create()

	err := service.Delete(context.Background(), roleResponse.ID)
	assert.Nil(t, err)

	_, err = service.FindById(context.Background(), roleResponse.ID)
	assert.NotNil(t, err)

}

func TestDeleteRoleServiceFailed(t *testing.T) {
	Truncate()
	Create()

	roleResponse := domain.Role{
		ID: "salah",
	}

	err := service.Delete(context.Background(), roleResponse.ID)
	assert.NotNil(t, err)
	assert.Equal(t, "role not found", err.Error())
}
