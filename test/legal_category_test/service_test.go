package legalcategorytest

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/nostracode/mangrove-api/model/domain"
	legalcategoryweb "github.com/nostracode/mangrove-api/model/web/legal-category-web"
	legalcategoryservice "github.com/nostracode/mangrove-api/service/legal_category_service"
	"github.com/stretchr/testify/assert"
)

var validate = validator.New()
var service = legalcategoryservice.NewLegalCategoryService(repo, db, validate)

func TestFindAllLegalCategoryService(t *testing.T) {

	Truncate()
	Create()
	legalCategorys, err := service.FindAll(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, 1, len(legalCategorys))

}

func TestFindByIdLegalCategoryService(t *testing.T) {

	Truncate()
	legalCategoryResponse, _ := Create()
	legalCategory, err := service.FindById(context.Background(), legalCategoryResponse.ID)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1", legalCategory.Name)

}

func TestFindAllLegalCategoryServiceFailed(t *testing.T) {
	_, err := service.FindById(context.Background(), "salah")
	assert.NotNil(t, err)
	assert.Equal(t, "legal category not found", err.Error())
}

func TestCreateLegalCategoryService(t *testing.T) {
	Truncate()
	legalCategory := &legalcategoryweb.LegalCategoryCreateReq{
		Name: "legal_cat_1",
	}

	legalCategoryResponse, err := service.Create(context.Background(), legalCategory)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1", legalCategoryResponse.Name)
}

// func TestCreatelegalCategoryServiceConflict(t *testing.T) {
// 	Truncate()
// 	legalCategoryResponse, _ := Create()
// 	user := &legalcategoryweb.LegalCategoryCreateReq{
// 		Name: legalCategoryResponse.Name,
// 	}

// 	_, err := service.Create(context.Background(), user)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "legalCategory name already create", err.Error())
// }

func TestCreateLegalCategoryServiceFailed(t *testing.T) {

	legalCategory := &legalcategoryweb.LegalCategoryCreateReq{
		Name: "",
	}

	_, err := service.Create(context.Background(), legalCategory)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is a required field", err.Error())

}

func TestUpdateLegalCategoryService(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()

	legalCategory := &legalcategoryweb.LegalCategoryUpdateReq{
		ID:   legalCategoryResponse.ID,
		Name: "legal_cat_1_updated",
	}

	response, err := service.Update(context.Background(), legalCategory)
	assert.Nil(t, err)
	assert.Equal(t, "legal_cat_1_updated", response.Name)
}

func TestUpdateLegalCategoryServiceFailed(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()

	legalCategory := &legalcategoryweb.LegalCategoryUpdateReq{
		ID:   legalCategoryResponse.ID,
		Name: "",
	}

	_, err := service.Update(context.Background(), legalCategory)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is a required field", err.Error())
}

func TestUpdateLegalCategoryServiceNorFound(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()

	legalCategory := &legalcategoryweb.LegalCategoryUpdateReq{
		ID:   "salah",
		Name: legalCategoryResponse.Name,
	}

	_, err := service.Update(context.Background(), legalCategory)
	assert.NotNil(t, err)
	assert.Equal(t, "legal category not found", err.Error())
}

func TestDeleteLegalCategoryService(t *testing.T) {

	Truncate()
	legalCategoryResponse, _ := Create()

	err := service.Delete(context.Background(), legalCategoryResponse.ID)
	assert.Nil(t, err)

	_, err = service.FindById(context.Background(), legalCategoryResponse.ID)
	assert.NotNil(t, err)

}

func TestDeletelegalCategoryServiceFailed(t *testing.T) {
	Truncate()
	Create()

	legalCategoryResponse := domain.LegalCategory{
		ID: "salah",
	}

	err := service.Delete(context.Background(), legalCategoryResponse.ID)
	assert.NotNil(t, err)
	assert.Equal(t, "legal category not found", err.Error())
}
