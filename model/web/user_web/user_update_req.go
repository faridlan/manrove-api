package userweb

type UserUpdateReq struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Name        string `json:"name,omitempty" validate:"required,gte=5"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	RoleId      string `json:"role_id,omitempty" validate:"required"`
	ImageUrl    string `json:"image_url,omitempty"`
}
