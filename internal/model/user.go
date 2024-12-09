package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:32" json:"username"`
	Password string `gorm:"size:128" json:"-"`
	Email    string `gorm:"size:128" json:"email"`
}
