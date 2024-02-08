package legalcategoryweb

type LegalCategoryUpdateReq struct {
	ID   string `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}
