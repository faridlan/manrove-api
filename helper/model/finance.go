package model

import (
	"github.com/nostracode/mangrove-api/model/domain"
	financeweb "github.com/nostracode/mangrove-api/model/web/finance_web"
)

func ToFinanceResponse(finance *domain.Finance) *financeweb.FinanceResponse {

	return &financeweb.FinanceResponse{
		ID:          finance.ID,
		Date:        finance.Date,
		IsDebit:     finance.IsDebit,
		UserId:      finance.UserId,
		Description: finance.Description,
		ImageUrl:    finance.ImageUrl,
		CreatedAt:   finance.CreatedAt,
		UpdatedAt:   finance.UpdatedAt,
	}

}

func ToFinanceResponses(finances []*domain.Finance) []*financeweb.FinanceResponse {

	financeResponses := []*financeweb.FinanceResponse{}
	for _, finance := range finances {
		financeResponses = append(financeResponses, ToFinanceResponse(finance))
	}

	return financeResponses

}
