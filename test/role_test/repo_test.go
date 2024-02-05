package roletest

import (
	"context"
	"fmt"
	"testing"

	"github.com/nostracode/mangrove-api/config/db/conn"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/model/domain"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	"github.com/stretchr/testify/assert"
)

var db = conn.NewDatabase()
var repo = rolerepo.NewRoleRepository()

func Truncate() error {
	err := db.Exec("TRUNCATE role CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

func Create() (*domain.Role, error) {
	role := &domain.Role{
		Name: "super_admin",
	}

	roleResponse, err := repo.Save(context.Background(), db, role)
	return roleResponse, err
}

func TestCreateRoleRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	role := &domain.Role{
		Name: "super_admin",
	}

	roleResponse, err := repo.Save(context.Background(), db, role)
	// fmt.Println(roleResponse.ID)
	fmt.Println(roleResponse)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", roleResponse.Name)
}

func TestFindAllRoleRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)

	_, err = Create()
	helper.PanicIfError(err)

	roles, err := repo.FindAll(context.Background(), db)
	assert.Nil(t, err)

	for _, role := range roles {
		fmt.Println(role)
	}

	assert.Equal(t, 1, len(roles))
}

func TestFindByIdRoleRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()
	fmt.Println(roleResponse.ID)

	roleResponse, err = repo.FindByID(context.Background(), db, roleResponse.ID)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", roleResponse.Name)

}

func TestDeleteRoleRepo(t *testing.T) {

	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()

	err = repo.Delete(context.Background(), db, roleResponse)
	assert.Nil(t, err)

	_, err = repo.FindByID(context.Background(), db, roleResponse.ID)

	assert.NotNil(t, err)

}

func TestUpdateRoleRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()

	role, err := repo.FindByID(context.Background(), db, roleResponse.ID)

	assert.Nil(t, err)

	role.Name = "role_updated"

	roleResponse, err = repo.Update(context.Background(), db, role)
	assert.Nil(t, err)
	assert.Equal(t, "role_updated", roleResponse.Name)
}
