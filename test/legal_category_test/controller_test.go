package legalcategorytest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	legalcategorycontroller "github.com/nostracode/mangrove-api/controller/legal_category_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/stretchr/testify/assert"
)

var controller = legalcategorycontroller.NewLegalCategoryController(service)
var app = fiber.New(
	fiber.Config{
		ErrorHandler: exception.ExceptionError,
	},
)

func ResponseTest(t *testing.T, request *http.Request, statusCode int) {
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	helper.PanicIfError(err)
	assert.Nil(t, err)
	assert.Equal(t, statusCode, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestCreatelegalCategoryController(t *testing.T) {
	Truncate()
	app.Post("/legal/categories", controller.Create)

	body := strings.NewReader(
		`{
				"name":"legal_cat_1"
		}`,
	)

	request := httptest.NewRequest("POST", "/legal/categories", body)
	ResponseTest(t, request, 200)

}

func TestCreatelegalCategoryControllerFailed(t *testing.T) {
	Truncate()

	app.Post("/legal/categories", controller.Create)

	body := strings.NewReader(
		`{
				"name":""
		}`,
	)

	request := httptest.NewRequest("POST", "/legal/categories", body)
	ResponseTest(t, request, 400)

}

// func TestCreatelegalCategoryControllerConflict(t *testing.T) {
// 	Truncate()
// 	legalCategoryResponse, _ := Create()

// 	app.Post("/legal/categories", controller.Create)

// 	body := strings.NewReader(
// 		fmt.Sprintf(`{
// 			"name":"%s"
// 	}`, legalCategoryResponse.Name),
// 	)

// 	request := httptest.NewRequest("POST", "/legal/categories", body)
// 	request.Header.Set("content-type", "application/json")
// 	response, err := app.Test(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Nil(t, err)
// 	assert.Equal(t, 409, response.StatusCode)
// }

func TestFindByIdlegalCategoryController(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()
	app.Get("/legal/categories/:legalCategoryId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/legal/categories/%s", legalCategoryResponse.ID), nil)
	ResponseTest(t, request, 200)

}

func TestFindByIdlegalCategoryControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Get("/legal/categories/:legalCategoryId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/legal/categories/%s", "salah"), nil)
	ResponseTest(t, request, 404)

}

func TestFindAlllegalCategoryController(t *testing.T) {
	Truncate()
	Create()
	app.Get("/legal/categories", controller.FindAll)

	request := httptest.NewRequest("GET", "/legal/categories", nil)
	ResponseTest(t, request, 200)

}

func TestDeletelegalCategoryController(t *testing.T) {
	Truncate()
	legalCategorysResponse, _ := Create()
	app.Delete("/legalCategorys/:legalCategoryId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/legalCategorys/%s", legalCategorysResponse.ID), nil)
	ResponseTest(t, request, 200)

}

func TestDeletelegalCategoryControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Delete("/legalCategorys/:legalCategoryId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/legalCategorys/%s", "salah"), nil)
	ResponseTest(t, request, 404)

}

func TestUpdatelegalCategoryController(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()
	app.Put("/legalCategorys/:legalCategoryId", controller.Update)

	body := strings.NewReader(
		`{
				"name":"legal_cat_1_updated"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/legalCategorys/%s", legalCategoryResponse.ID), body)
	ResponseTest(t, request, 200)
}

func TestUpdatelegalCategoryControllerFailed(t *testing.T) {
	Truncate()
	legalCategoryResponse, _ := Create()
	app.Put("/legalCategorys/:legalCategoryId", controller.Update)

	body := strings.NewReader(
		`{
				"name":""
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/legalCategorys/%s", legalCategoryResponse.ID), body)
	ResponseTest(t, request, 400)

}

func TestUpdatelegalCategoryControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Put("/legalCategorys/:legalCategoryId", controller.Update)

	body := strings.NewReader(
		`{
				"name":"updated_hehe"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/legalCategorys/%s", "salah"), body)
	ResponseTest(t, request, 404)

}
