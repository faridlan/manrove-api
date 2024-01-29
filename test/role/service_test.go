package role

import (
	"context"
	"testing"

	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
	roleservice "github.com/nostracode/mangrove-api/service/role_service"
	"github.com/stretchr/testify/assert"
)

var service = roleservice.NewRoleService(repo, db)

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
	role, err := service.FindById(context.Background(), roleResponse.UID)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", role.Name)

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

func TestUpdateRoleService(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()

	role := &roleweb.RoleUpdateReq{
		UID:  roleResponse.UID,
		Name: "super_admin_updated",
	}

	response, err := service.Update(context.Background(), role)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin_updated", response.Name)
}

func TestDeleteRoleService(t *testing.T) {

	Truncate()
	roleResponse, _ := Create()

	err := service.Delete(context.Background(), roleResponse.UID)
	assert.Nil(t, err)

	_, err = service.FindById(context.Background(), roleResponse.UID)
	assert.NotNil(t, err)

}
