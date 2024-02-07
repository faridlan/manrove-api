package domain

import "gorm.io/gorm"

type Finance struct {
	ID          string         `gorm:"primarykey;column:id;<-:create"`
	Date        int64          `gorm:"column:date"`
	IsDebit     bool           `gorm:"column:is_debit"`
	UserId      string         `gorm:"column:user_id"`
	Description string         `gorm:"column:description"`
	ImageUrl    string         `gorm:"column:image_url"`
	CreatedAt   int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *Finance) TableName() string {
	return "finance"
}
