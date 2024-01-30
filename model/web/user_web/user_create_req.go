package userweb

type UserCreateReq struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RoleId      string `json:"role_id,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}
