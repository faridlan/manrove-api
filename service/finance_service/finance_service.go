package financeservice

import (
	"context"

	financeweb "github.com/nostracode/mangrove-api/model/web/finance_web"
)

type FinanceService interface {
	Create(ctx context.Context, request *financeweb.FinanceCreateReq) (*financeweb.FinanceResponse, error)
	Update(ctx context.Context, request *financeweb.FinanceUpdateReq) (*financeweb.FinanceResponse, error)
	Delete(ctx context.Context, financeId string) error
	FindById(ctx context.Context, financeId string) (*financeweb.FinanceResponse, error)
	FindAll(ctx context.Context) ([]*financeweb.FinanceResponse, error)
}
