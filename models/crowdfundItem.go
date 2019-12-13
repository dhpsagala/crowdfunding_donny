package models

import "github.com/jinzhu/gorm"

type ICrowdfundItem interface {
}

type crowdfundItem struct {
	gorm.Model
	Name  string
	Stock int
	Price float32
}
