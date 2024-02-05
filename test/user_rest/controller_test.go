package userrest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	usercontroller "github.com/nostracode/mangrove-api/controller/user_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/stretchr/testify/assert"
)

var controller = usercontroller.NewUserController(service)
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

func TestRegisterUserController(t *testing.T) {
	userTruncate()
	app.Post("/users", controller.Register)

	body := strings.NewReader(
		`{
			"email":       "user1@mail.com",
			"name":        "user1",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default.png"
		}`,
	)

	request := httptest.NewRequest("POST", "/users", body)
	ResponseTest(t, request, 200)
}

func TestRegisterUserControllerFailed(t *testing.T) {
	userTruncate()
	app.Post("/users", controller.Register)

	body := strings.NewReader(
		`{
			"name":        "user1",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default.png"
		}`,
	)

	request := httptest.NewRequest("POST", "/users", body)
	ResponseTest(t, request, 400)
}

func TestRegisterUserControllerConflict(t *testing.T) {
	userTruncate()
	app.Post("/users", controller.Register)
	userCreate()

	body := strings.NewReader(
		`{
					"email":       "user2@mail.com",
					"name":        "user1",
					"password":    "secret1234565",
					"phone_number": "33300998822",
					"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
					"ImageUrl":    "https://image/default.png"
				}`,
	)

	request := httptest.NewRequest("POST", "/users", body)
	ResponseTest(t, request, 409)
}

func TestFindByIdUserControllerSuccess(t *testing.T) {
	userTruncate()
	userResponse, _ := userCreate()
	app.Get("/users/:userId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", userResponse.ID), nil)
	ResponseTest(t, request, 200)
}

func TestFindByIdUserControllerNotFound(t *testing.T) {
	userTruncate()
	userCreate()
	app.Get("/users/:userId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", "salah"), nil)
	ResponseTest(t, request, 404)
}

func TestFindAllUserController(t *testing.T) {
	userTruncate()
	for i := 1; i <= 2; i++ {
		userCreate()
	}
	app.Get("/users", controller.FindAll)

	request := httptest.NewRequest("GET", "/users", nil)
	ResponseTest(t, request, 200)
}

func TestDeleteUserControllerSuccess(t *testing.T) {
	userTruncate()
	userResponse, _ := userCreate()
	app.Delete("/users/:userId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%s", userResponse.ID), nil)
	ResponseTest(t, request, 200)
}

func TestDeleteUserControllerNotFound(t *testing.T) {
	userTruncate()
	userCreate()
	app.Delete("/users/:userId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%s", "salah"), nil)
	ResponseTest(t, request, 404)
}

func TestUpdateUserController(t *testing.T) {
	userTruncate()
	userResponse, _ := userCreate()
	app.Put("/users/:userId", controller.Update)

	body := strings.NewReader(
		`{
			"email":       "user1_updated@mail.com",
			"name":        "user1_updated",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", userResponse.ID), body)
	ResponseTest(t, request, 200)

}

func TestUpdateUserControllerFailed(t *testing.T) {
	userTruncate()
	userResponse, _ := userCreate()
	app.Put("/users/:userId", controller.Update)

	body := strings.NewReader(
		`{
			"name":        "user1_updated",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", userResponse.ID), body)
	ResponseTest(t, request, 400)

}

func TestUpdateUserControllerConflict(t *testing.T) {
	userTruncate()
	userResponse, _ := userCreate()
	app.Put("/users/:userId", controller.Update)

	body := strings.NewReader(
		`{
			"email":       "user1_updated@mail.com",
			"name":        "user1",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", userResponse.ID), body)
	ResponseTest(t, request, 409)

}

func TestUpdateUserControllerNotFound(t *testing.T) {
	userTruncate()
	userCreate()
	app.Put("/users/:userId", controller.Update)

	body := strings.NewReader(
		`{
			"email":       "user1_updated@mail.com",
			"name":        "user1",
			"password":    "secret1234565",
			"phone_number": "33300998822",
			"role_id":      "a97e962c9f9e458ebe05e4d7b60a70f3",
			"ImageUrl":    "https://image/default_updated.png"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", "salah"), body)
	ResponseTest(t, request, 404)
}
