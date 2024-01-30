package userweb

type UserUpdateReq struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RoleId      string `json:"role_id,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}
