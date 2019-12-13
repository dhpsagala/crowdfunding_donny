package models

import "github.com/jinzhu/gorm"

type IBuyedItem interface {
}

type buyedItem struct {
	gorm.Model
	ItemID     int `gorm:"foreignkey:crowdfund_items(id)"`
	UserID     int `gorm:"foreignkey:users(id)"`
	Quantity   int
	TotalPrice float32
}
