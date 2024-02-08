package model

import (
	"github.com/nostracode/mangrove-api/model/domain"
	legalcategoryweb "github.com/nostracode/mangrove-api/model/web/legal-category-web"
)

func ToLegalCategoryResponse(legal_category *domain.LegalCategory) *legalcategoryweb.LegalCategoryResponse {

	return &legalcategoryweb.LegalCategoryResponse{
		ID:        legal_category.ID,
		Name:      legal_category.Name,
		CreatedAt: legal_category.CreatedAt,
		UpdatedAt: legal_category.UpdatedAt,
	}

}

func ToLegalCategoryResponses(legal_categories []*domain.LegalCategory) []*legalcategoryweb.LegalCategoryResponse {
	legalCategoriesResponses := []*legalcategoryweb.LegalCategoryResponse{}
	for _, legalCategory := range legal_categories {
		legalCategoriesResponses = append(legalCategoriesResponses, ToLegalCategoryResponse(legalCategory))
	}

	return legalCategoriesResponses
}
