package domain

import "gorm.io/gorm"

type User struct {
	UID         string         `gorm:"primarykey;column:uid;<-:create"`
	Email       string         `gorm:"column:email"`
	Name        string         `gorm:"column:name"`
	PhoneNumber string         `gorm:"column:phone_number"`
	RoleId      string         `gorm:"column:role_id"`
	ImageUrl    string         `gorm:"column:image_url"`
	FirstVisit  bool           `gorm:"column:first_visit"`
	CreatedAt   int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "role"
}