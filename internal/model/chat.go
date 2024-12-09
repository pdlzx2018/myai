package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	UserID    uint   `gorm:"index" json:"user_id"`
	Message   string `gorm:"type:text" json:"message"`
	Response  string `gorm:"type:text" json:"response"`
	ModelName string `gorm:"size:32" json:"model_name"`
}
