package models

import "github.com/jinzhu/gorm"

type Poetry struct {
	gorm.Model
	Title string
	Content string
	Author string
	Remark string
	Status    int
}

func (Poetry) TableName() string {
	return "poetry"
}
