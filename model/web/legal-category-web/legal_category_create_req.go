package legalcategoryweb

type LegalCategoryCreateReq struct {
	Name string `json:"name,omitempty" validate:"required"`
}
