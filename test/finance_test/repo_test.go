package financetest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nostracode/mangrove-api/config/db/conn"
	"github.com/nostracode/mangrove-api/helper"
	"github.com/nostracode/mangrove-api/model/domain"
	financerepo "github.com/nostracode/mangrove-api/repository/finance_repo"
	"github.com/stretchr/testify/assert"
)

var db = conn.NewDatabase()
var repo = financerepo.NewFinanceRepository()

func financeTruncate() error {
	err := db.Exec("TRUNCATE finance CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

func financeCreate() (*domain.Finance, error) {

	finance := &domain.Finance{
		Date:        time.Now().UnixMilli(),
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance",
		ImageUrl:    "https://image.com/jasdfljaf.png",
	}

	financeResponse, err := repo.Save(context.Background(), db, finance)
	if err != nil {
		return nil, err
	}

	return financeResponse, nil
}

func TestCreateFinanceRepo(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance := &domain.Finance{
		Date:        time.Now().UnixMilli(),
		IsDebit:     true,
		UserId:      "6331deffc6a84f4e88a9881c0a839725",
		Description: "This is the first user finance",
		ImageUrl:    "https://image.com/jasdfljaf.png",
	}

	financeResponse, err := repo.Save(context.Background(), db, finance)
	helper.PanicIfError(err)

	fmt.Println(financeResponse)

	assert.Nil(t, err)
	assert.Equal(t, "This is the first user finance", financeResponse.Description)
}

func TestFindAllFinanceRepo(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		_, err = financeCreate()
		assert.Nil(t, err)
	}

	financeResponse, err := repo.FindAll(context.Background(), db)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(financeResponse))

}

func TestFindByIdFinanceRepo(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)
	finance, err := financeCreate()
	assert.Nil(t, err)

	financeResponse, err := repo.FindById(context.Background(), db, finance.ID)
	helper.PanicIfError(err)

	fmt.Println(financeResponse)

	assert.Nil(t, err)
	assert.Equal(t, "This is the first user finance", financeResponse.Description)
}

func TestDeleteUserRepo(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance, err := financeCreate()
	assert.Nil(t, err)

	err = repo.Delete(context.Background(), db, finance)
	assert.Nil(t, err)

	_, err = repo.FindById(context.Background(), db, finance.ID)
	assert.NotNil(t, err)

}

func TestUpdateUserRepo(t *testing.T) {
	err := financeTruncate()
	assert.Nil(t, err)

	finance, err := financeCreate()
	assert.Nil(t, err)

	finance, err = repo.FindById(context.Background(), db, finance.ID)
	assert.Nil(t, err)

	finance.Date = time.Now().UnixMilli()
	finance.IsDebit = true
	finance.UserId = "6331deffc6a84f4e88a9881c0a839725"
	finance.Description = "https://image.com/jasdfljaf.png"
	finance.ImageUrl = "https://image.com/jasdfljaf_updated.png"

	financeResponse, err := repo.Update(context.Background(), db, finance)
	assert.Nil(t, err)

	assert.Equal(t, finance.Description, financeResponse.Description)
}
