package role

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/model/domain"
	rolerepo "github.com/nostracode/mangrove-api/repository/role_repo"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {

	dsn := "host=localhost user=nullhakim password=NullHakimNostra123 dbname=mangrove_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db

}

var db = OpenConnection()
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

func TestCreateUserRepo(t *testing.T) {
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

func TestFindAllRepo(t *testing.T) {
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

func TestFindByIdRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()
	fmt.Println(roleResponse.UID)

	roleResponse, err = repo.FindByUID(context.Background(), db, roleResponse.UID)
	assert.Nil(t, err)
	assert.Equal(t, "super_admin", roleResponse.Name)

}

func TestDeleteRepo(t *testing.T) {

	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()

	err = repo.Delete(context.Background(), db, roleResponse)
	assert.Nil(t, err)

	_, err = repo.FindByUID(context.Background(), db, roleResponse.UID)

	assert.NotNil(t, err)

}

func TestUpdateRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	roleResponse, _ := Create()

	role, err := repo.FindByUID(context.Background(), db, roleResponse.UID)

	assert.Nil(t, err)

	role.Name = "role_updated"

	roleResponse, err = repo.Update(context.Background(), db, role)
	assert.Nil(t, err)
	assert.Equal(t, "role_updated", roleResponse.Name)
}
