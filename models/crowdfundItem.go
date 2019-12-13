package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type ICrowdfundItem interface {
	Buy(qty int) (IBuyedItem, error)
}

type crowdfundItem struct {
	gorm.Model
	Name  string
	Stock int
	Price float32
}

func GetAvailableCroudfundItems() ([]ICrowdfundItem, error) {
	var items []crowdfundItem
	datas := []ICrowdfundItem{}
	if rows, err := db.Where("stock > ?", 0).Find(&items).Rows(); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var r crowdfundItem
			if err := db.ScanRows(rows, &r); err != nil {
				return nil, err
			} else {
				datas = append(datas, r)
			}
		}
	}
	return datas, nil
}

func GetCroudfundItem(id int) (ICrowdfundItem, error) {
	var p crowdfundItem
	err := db.First(&p, id).Error
	return &p, err
}

func (p crowdfundItem) Buy(qty int) (IBuyedItem, error) {
	if p.Stock == 0 || p.Stock < qty {
		return nil, errors.New("Not enough stock")
	}
	var err error
	var dbOri *gorm.DB = db
	db = db.New().Begin()
	p.Stock -= qty
	if err = db.Save(p).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	var bi IBuyedItem
	if bi, err = NewBuyedItem(p.ID, qty, p.Price); err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	db = dbOri
	return bi, nil
}
