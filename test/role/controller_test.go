package role

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	rolecontroller "github.com/nostracode/mangrove-api/controller/role_controller"
	"github.com/stretchr/testify/assert"
)

var controller = rolecontroller.NewRoleController(service)

func TestCreateRoleController(t *testing.T) {
	Truncate()
	app := fiber.New()
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

func TestFindByIdRoleController(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	app := fiber.New()
	app.Get("/roles/:roleId", controller.FindById)

	request := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s", roleResponse.UID), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))
}

func TestFindAllController(t *testing.T) {
	Truncate()
	Create()
	app := fiber.New()
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
	app := fiber.New()
	app.Delete("/roles/:roleId", controller.Delete)

	request := httptest.NewRequest("DELETE", fmt.Sprintf("/roles/%s", rolesResponse.UID), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateController(t *testing.T) {
	Truncate()
	roleResponse, _ := Create()
	app := fiber.New()
	app.Put("/roles/:roleId", controller.Update)

	body := strings.NewReader(
		`{
				"name":"super_admin_updated"
		}`,
	)

	request := httptest.NewRequest("PUT", fmt.Sprintf("/roles/%s", roleResponse.UID), body)
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
