package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

type Model struct {
	ID int `gorm:"primark_key" json:"id"`
	CreateOn int `gorm:"creat_on" json:"create_on"`
	
}
