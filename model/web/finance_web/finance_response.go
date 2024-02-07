package financeweb

type FinanceResponse struct {
	ID          string `json:"id,omitempty"`
	Date        int64  `json:"date,omitempty"`
	IsDebit     bool   `json:"is_debit,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}
