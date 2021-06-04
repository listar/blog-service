package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string
	Avatar string
	Role int
	Password string
}

func (User) TableName() string {
	return "users"
}
