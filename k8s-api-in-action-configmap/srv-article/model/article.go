package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Userid  uint   `gorm:"type:int(10) unsigned;not null" json:"userid"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
}
