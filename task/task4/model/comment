package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserId  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserId" json:"user,omitempty"`
	PostId  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostId" json:"post,omitempty"`
}
