package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Posts    []Post    `gorm:"foreignKey:UserId" json:"posts,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserId" json:"comments,omitempty"`
}

