package financetest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	financeweb "github.com/nostracode/mangrove-api/model/web/finance_web"
	financeservice "github.com/nostracode/mangrove-api/service/finance_service"
	"github.com/stretchr/testify/assert"
)

var validate = validator.New()
var service = financeservice.NewFinanceRepository(repo, db, validate)

func TestFindallFinanceService(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		_, err = financeCreate()
		assert.Nil(t, err)
	}

	financeResponse, err := service.FindAll(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 2, len(financeResponse))
}

func TestFindByIdFinanceServiceSuccess(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance, err := financeCreate()
	assert.Nil(t, err)

	financeResponse, err := service.FindById(context.Background(), finance.ID)
	assert.Nil(t, err)
	assert.Equal(t, finance.Date, financeResponse.Date)
	assert.Equal(t, finance.IsDebit, financeResponse.IsDebit)
	assert.Equal(t, finance.UserId, financeResponse.UserId)
	assert.Equal(t, finance.Description, financeResponse.Description)
	assert.Equal(t, finance.ImageUrl, financeResponse.ImageUrl)
}

func TestFindByIdFinanceServiceFailed(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	_, err = financeCreate()
	assert.Nil(t, err)

	_, err = service.FindById(context.Background(), "salah")
	assert.NotNil(t, err)
	assert.Equal(t, "finance not found", err.Error())
}

func TestCreateFinanceServiceSuccess(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance := &financeweb.FinanceCreateReq{
		Date:        time.Now().UnixMilli(),
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance",
		ImageUrl:    "https://image.com/jasdfljaf.png",
	}

	financeResponse, err := service.Create(context.Background(), finance)
	defer fmt.Println(err)
	assert.Nil(t, err)
	assert.Equal(t, finance.Date, financeResponse.Date)
	assert.Equal(t, finance.IsDebit, financeResponse.IsDebit)
	assert.Equal(t, finance.UserId, financeResponse.UserId)
	assert.Equal(t, finance.Description, financeResponse.Description)
	assert.Equal(t, finance.ImageUrl, financeResponse.ImageUrl)
}

func TestCreateFinanceServiceFailed(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	financeCreate := &financeweb.FinanceCreateReq{
		Date:        0,
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance",
		ImageUrl:    "https://image.com/jasdfljaf.png",
	}

	_, err = service.Create(context.Background(), financeCreate)
	assert.NotNil(t, err)
	assert.Equal(t, "Date is a required field", err.Error())
}

func TestUpdateFinanceServiceSuccess(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)
	finance, err := financeCreate()
	assert.Nil(t, err)

	financeUpdate := &financeweb.FinanceUpdateReq{
		ID:          finance.ID,
		Date:        time.Now().UnixMilli(),
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance updated",
		ImageUrl:    "https://image.com/jasdfljaf_updated.png",
	}

	financeResponse, err := service.Update(context.Background(), financeUpdate)
	assert.Nil(t, err)
	assert.Equal(t, financeUpdate.Date, financeResponse.Date)
	assert.Equal(t, financeUpdate.IsDebit, financeResponse.IsDebit)
	assert.Equal(t, financeUpdate.UserId, financeResponse.UserId)
	assert.Equal(t, financeUpdate.Description, financeResponse.Description)
	assert.Equal(t, financeUpdate.ImageUrl, financeResponse.ImageUrl)
}

func TestUpdateFinanceServiceFailed(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)
	finance, err := financeCreate()
	assert.Nil(t, err)

	financeUpdate := &financeweb.FinanceUpdateReq{
		ID:          finance.ID,
		Date:        0,
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance updated",
		ImageUrl:    "https://image.com/jasdfljaf_updated.png",
	}

	_, err = service.Update(context.Background(), financeUpdate)
	assert.NotNil(t, err)
	assert.Equal(t, "Date is a required field", err.Error())
}

func TestUpdateFinanceServiceNotFound(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)
	_, err = financeCreate()
	assert.Nil(t, err)

	financeUpdate := &financeweb.FinanceUpdateReq{
		ID:          "salah",
		Date:        time.Now().UnixMilli(),
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance updated",
		ImageUrl:    "https://image.com/jasdfljaf_updated.png",
	}

	_, err = service.Update(context.Background(), financeUpdate)
	assert.NotNil(t, err)
	assert.Equal(t, "finance not found", err.Error())
}

func TestDeleteFinanceServiceSuccess(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance, err := financeCreate()
	assert.Nil(t, err)

	err = service.Delete(context.Background(), finance.ID)
	assert.Nil(t, err)

	_, err = service.FindById(context.Background(), finance.ID)
	assert.NotNil(t, err)
	assert.Equal(t, "finance not found", err.Error())
}

func TestDeleteFinanceServiceNotFound(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	_, err = financeCreate()
	assert.Nil(t, err)

	err = service.Delete(context.Background(), "salah")
	assert.NotNil(t, err)
	assert.Equal(t, "finance not found", err.Error())
}
