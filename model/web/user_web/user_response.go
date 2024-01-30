package userweb

import "gorm.io/gorm"

type UserResponse struct {
	ID          string         `json:"id,omitempty"`
	Email       string         `json:"email,omitempty"`
	Name        string         `json:"name,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	RoleId      string         `json:"role_id,omitempty"`
	ImageUrl    string         `json:"image_url,omitempty"`
	FirstVisit  bool           `json:"first_visit,omitempty"`
	CreatedAt   int64          `json:"created_at,omitempty"`
	UpdatedAt   int64          `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}
