package financeweb

type FinanceUpdateReq struct {
	ID          string `json:"id,omitempty" validate:"required"`
	Date        int64  `json:"date,omitempty" validate:"required"`
	IsDebit     bool   `json:"is_debit,omitempty" validate:"required"`
	UserId      string `json:"user_id,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	ImageUrl    string `json:"image_url,omitempty" validate:"required"`
}
