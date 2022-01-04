package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(11);not null;unique" json:"phone"`
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
}
