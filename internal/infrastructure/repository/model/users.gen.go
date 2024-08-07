// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID          string    `gorm:"column:id;type:varchar(100);primaryKey" json:"id"`
	UserName    string    `gorm:"column:user_name;type:varchar(100);not null;uniqueIndex:users_unique,priority:1" json:"user_name"`
	DisplayName string    `gorm:"column:display_name;type:varchar(100);not null" json:"display_name"`
	Biography   string    `gorm:"column:biography;type:text;not null" json:"biography"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	IconURL     string    `gorm:"column:icon_url;type:varchar(100);not null" json:"icon_url"`
	BgimageURL  string    `gorm:"column:bgimage_url;type:varchar(100);not null" json:"bgimage_url"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
