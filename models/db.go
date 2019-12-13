package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	db *gorm.DB
)

func InitDb() {
	db_file := os.Getenv("DB_FILE")
	var err error

	if db, err = gorm.Open("sqlite3", db_file); err != nil {
		fmt.Print(err)
	} else {
		dbMigration()
		dummyCrowdfundingItems()
	}
}

func dbMigration() {
	db.AutoMigrate(&buyedItem{}, &crowdfundItem{}, &user{})
}

func dummyCrowdfundingItems() {
	items := []*crowdfundItem{
		&crowdfundItem{
			Name:  "Air Guitar",
			Stock: 100,
			Price: 500.99,
		},
		&crowdfundItem{
			Name:  "Flying Car",
			Stock: 10,
			Price: 2000.50,
		},
		&crowdfundItem{
			Name:  "Transparent Phone",
			Stock: 1000,
			Price: 100.25,
		},
		&crowdfundItem{
			Name:  "Dual Screen Laptop",
			Stock: 50,
			Price: 125.50,
		},
	}

	tx := db.Begin()
	for _, item := range items {
		var xisItem crowdfundItem
		if tx.Where(item).First(&xisItem).RecordNotFound() {
			if tx.Save(&item).Error != nil {
				tx.Rollback()
				fmt.Println(tx.Error)
				return
			}
		}
	}
	tx.Commit()
}
