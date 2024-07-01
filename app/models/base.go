package models

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	Name string
	Age  int
}
