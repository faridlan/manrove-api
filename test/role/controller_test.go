package role

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	rolecontroller "github.com/nostracode/mangrove-api/controller/role_controller"
	"github.com/nostracode/mangrove-api/exception"
	"github.com/stretchr/testify/assert"
)

var controller = rolecontroller.NewRoleController(service)
var app = fiber.New(
	fiber.Config{
		ErrorHandler: exception.ExceptionError,
	},
)

func TestCreateRoleController(t *testing.T) {
	Truncate()
	app.Post("/roles", controller.Create)

	body := strings.NewReader(
		`{
				"name":"super_admin"
		}`,
	)

	request := httptest.NewRequest("POST", "/roles", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestCreateRoleControllerFailed(t *testing.T) {
	Truncate()

	app.Post("/roles", controller.Create)

	body := strings.NewReader(
		`{
				"name":""
		}`,
	)

	request := httptest.NewRequest("POST", "/roles", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateRoleControllerConflict(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()

	app.Post("/roles", controller.Create)

	body := strings.NewReader(
		fmt.Sprintf(`{
			"name":"%s"
	}`, roleResponse.Name),
	)

	request := httptest.NewRequest("POST", "/roles", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 409, response.StatusCode)
}

func TestFindByIdRoleController(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	app.Get("/roles/:roleId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s", roleResponse.ID), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestFindByIdRoleControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Get("/roles/:roleId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s", "salah"), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 404, response.StatusCode)
}

func TestFindAllController(t *testing.T) {
	Truncate()
	Create()
	app.Get("/roles", controller.FindAll)

	request := httptest.NewRequest("GET", "/roles", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestDeleteController(t *testing.T) {
	Truncate()
	rolesResponse, _ := Create()
	app.Delete("/roles/:roleId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/roles/%s", rolesResponse.ID), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Delete("/roles/:roleId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/roles/%s", "salah"), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateController(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	app.Put("/roles/:roleId", controller.Update)

	body := strings.NewReader(
		`{
				"name":"super_admin_updated"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/roles/%s", roleResponse.ID), body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestUpdateControllerFailed(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	app.Put("/roles/:roleId", controller.Update)

	body := strings.NewReader(
		`{
				"name":""
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/roles/%s", roleResponse.ID), body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateControllerNotFound(t *testing.T) {
	Truncate()
	Create()
	app.Put("/roles/:roleId", controller.Update)

	body := strings.NewReader(
		`{
				"name":"updated_hehe"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/roles/%s", "salah"), body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 404, response.StatusCode)
}
