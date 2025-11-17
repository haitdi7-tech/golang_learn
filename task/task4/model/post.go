package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"titel"`
	Content  string    `gorm:"not null" json:"content"`
	UserId   uint      `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserId" json:"comments,omitempty"`
}
