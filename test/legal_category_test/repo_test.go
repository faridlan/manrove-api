package legalcategorytest

import (
	"context"
	"fmt"
	"testing"

	"github.com/nostracode/mangrove-api/config/db/conn"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/model/domain"
	legalcategoryrepo "github.com/nostracode/mangrove-api/repository/legal_category_repo"
	"github.com/stretchr/testify/assert"
)

var db = conn.NewDatabase()
var repo = legalcategoryrepo.NewLegalCategoryRepository()

func Truncate() error {
	err := db.Exec("TRUNCATE legal_category CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

func Create() (*domain.LegalCategory, error) {
	legalCategory := &domain.LegalCategory{
		Name: "legal_cat_1",
	}

	legalCategoryRes, err := repo.Save(context.Background(), db, legalCategory)
	return legalCategoryRes, err
}

func TestCreateLegalCategoryRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	legalCategory := &domain.LegalCategory{
		Name: "legal_cat_1",
	}

	legalCategoryResponse, err := repo.Save(context.Background(), db, legalCategory)
	// fmt.Println(legalCategoryResponse.ID)
	fmt.Println(legalCategoryResponse)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1", legalCategoryResponse.Name)
}

func TestFindAllLegalCategoryRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)

	for i := 1; i <= 2; i++ {
		_, err = Create()
		helper.PanicIfError(err)
	}

	legalCategorys, err := repo.FindAll(context.Background(), db)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(legalCategorys))
}

func TestFindByIdLegalCategoryRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	legalCategoryResponse, _ := Create()
	fmt.Println(legalCategoryResponse.ID)

	legalCategoryResponse, err = repo.FindById(context.Background(), db, legalCategoryResponse.ID)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1", legalCategoryResponse.Name)

}

func TestDeleteLegalCategoryRepo(t *testing.T) {

	err := Truncate()
	helper.PanicIfError(err)
	legalCategoryResponse, _ := Create()

	err = repo.Delete(context.Background(), db, legalCategoryResponse)
	assert.Nil(t, err)

	_, err = repo.FindById(context.Background(), db, legalCategoryResponse.ID)

	assert.NotNil(t, err)

}

func TestUpdateLegalCategoryRepo(t *testing.T) {
	err := Truncate()
	helper.PanicIfError(err)
	legalCategoryResponse, _ := Create()

	legalCategory, err := repo.FindById(context.Background(), db, legalCategoryResponse.ID)

	assert.Nil(t, err)

	legalCategory.Name = "legal_cat_1_updated"

	legalCategoryResponse, err = repo.Update(context.Background(), db, legalCategory)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1_updated", legalCategoryResponse.Name)
}
