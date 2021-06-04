package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Category  int
	Title     string
	Tags      string
	Content   string
	Author    string
	Status    int
	Intro     string
	OuterLink string
}

func (Article) TableName() string {
	return "article"
}
