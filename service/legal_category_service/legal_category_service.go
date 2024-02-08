package legalcategoryservice

import (
	"context"

	legalcategoryweb "github.com/nostracode/mangrove-api/model/web/legal-category-web"
)

type LegalCategoryService interface {
	Create(ctx context.Context, request *legalcategoryweb.LegalCategoryCreateReq) (*legalcategoryweb.LegalCategoryResponse, error)
	Update(ctx context.Context, request *legalcategoryweb.LegalCategoryUpdateReq) (*legalcategoryweb.LegalCategoryResponse, error)
	Delete(ctx context.Context, legalCategoryId string) error
	FindById(ctx context.Context, legalCategoryId string) (*legalcategoryweb.LegalCategoryResponse, error)
	FindAll(ctx context.Context) ([]*legalcategoryweb.LegalCategoryResponse, error)
}
