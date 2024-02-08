package financetest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	financecontroller "github.com/nostracode/mangrove-api/controller/finance_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/stretchr/testify/assert"
)

var controller = financecontroller.NewFinanceController(service)
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

func TestCreateFinanceControllerSuccess(t *testing.T) {
	financeTruncate()
	app.Post("/finances", controller.Create)

	body := strings.NewReader(
		`{
			"date":       33300998822,
			"is_debit":   true,
			"user_id":    "6331deffc6a84f4e88a9881c0a839725",
			"description": "this is the first user finance",
			"image_url":    "https://image/default.png"
		}`,
	)

	request := httptest.NewRequest("POST", "/finances", body)
	ResponseTest(t, request, 200)
}

func TestCreateFinanceControllerFailed(t *testing.T) {
	financeTruncate()
	app.Post("/finances", controller.Create)

	body := strings.NewReader(
		`{
			"is_debit":   true,
			"user_id":    "6331deffc6a84f4e88a9881c0a839725",
			"description": "this is the first user finance",
			"image_url":    "https://image/default.png"
		}`,
	)

	request := httptest.NewRequest("POST", "/finances", body)
	ResponseTest(t, request, 400)
}

func TestFindByIdFinanceControllerSuccess(t *testing.T) {
	financeTruncate()
	financeResponse, _ := financeCreate()
	app.Get("/finances/:financeId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/finances/%s", financeResponse.ID), nil)
	ResponseTest(t, request, 200)
}

func TestFindByIdFinanceControllerFailed(t *testing.T) {
	financeTruncate()
	financeCreate()
	app.Get("/finances/:financeId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/finances/%s", "salah"), nil)
	ResponseTest(t, request, 404)
}

func TestFindAllFinanceController(t *testing.T) {
	financeTruncate()
	for i := 1; i <= 2; i++ {
		financeCreate()
	}
	app.Get("/finances", controller.FindAll)

	request := httptest.NewRequest("GET", "/finances", nil)
	ResponseTest(t, request, 200)
}

func TestDeleteFinanceControllerSuccess(t *testing.T) {
	financeTruncate()
	financeResponse, _ := financeCreate()
	app.Delete("/finances/:financeId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/finances/%s", financeResponse.ID), nil)
	ResponseTest(t, request, 200)
}

func TestDeleteFinanceControllerFailed(t *testing.T) {
	financeTruncate()
	financeCreate()
	app.Delete("/finances/:financeId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/finances/%s", "salah"), nil)
	ResponseTest(t, request, 404)
}

func TestUpdateFinanceControllerSuccess(t *testing.T) {
	financeTruncate()
	financeResponse, _ := financeCreate()
	app.Put("/finances/:financeId", controller.Update)

	body := strings.NewReader(
		`{
			"date":       3330099882243,
			"is_debit":   true,
			"user_id":    "6331deffc6a84f4e88a9881c0a839725",
			"description": "this is the first user finance updated",
			"image_url":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/finances/%s", financeResponse.ID), body)
	ResponseTest(t, request, 200)

}

func TestUpdateFinanceControllerFailed(t *testing.T) {
	financeTruncate()
	financeResponse, _ := financeCreate()
	app.Put("/finances/:financeId", controller.Update)

	body := strings.NewReader(
		`{
			"is_debit":   true,
			"user_id":    "6331deffc6a84f4e88a9881c0a839725",
			"description": "this is the first user finance updated",
			"image_url":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/finances/%s", financeResponse.ID), body)
	ResponseTest(t, request, 400)
}

func TestUpdateFinanceControllerNotFound(t *testing.T) {
	financeTruncate()
	financeCreate()
	app.Put("/finances/:financeId", controller.Update)

	body := strings.NewReader(
		`{
			"date":       3330099882243,
			"is_debit":   true,
			"user_id":    "6331deffc6a84f4e88a9881c0a839725",
			"description": "this is the first user finance updated",
			"iamge_url":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/finances/%s", "salah"), body)
	ResponseTest(t, request, 404)
}
