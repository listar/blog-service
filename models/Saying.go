package models

import "github.com/jinzhu/gorm"

type Saying struct {
	gorm.Model
	Content string
	Author  string
	Status  int
}

func (Saying) TableName() string {
	return "saying"
}
