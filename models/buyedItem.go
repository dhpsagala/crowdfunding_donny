package models

import "github.com/jinzhu/gorm"

import "errors"

type IBuyedItem interface {
	Insert() error
}

type buyedItem struct {
	gorm.Model
	ItemID     uint `gorm:"foreignkey:crowdfund_items(id)"`
	UserID     uint `gorm:"foreignkey:users(id)"`
	Quantity   int
	TotalPrice float32
}

func NewBuyedItem(itemId uint, qty int, price float32) (IBuyedItem, error) {
	if CurrUser == nil {
		return nil, errors.New("User not found")
	}
	p := &buyedItem{
		ItemID:     itemId,
		UserID:     CurrUser.GetID(),
		Quantity:   qty,
		TotalPrice: float32(qty) * price,
	}
	if err := p.Insert(); err != nil {
		return nil, err
	}
	return p, nil
}

func GetUserTransaction() ([]IBuyedItem, error) {
	var items []buyedItem
	datas := []IBuyedItem{}
	if rows, err := db.Where("user_id == ?", CurrUser.GetID()).Find(&items).Rows(); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var r buyedItem
			if err := db.ScanRows(rows, &r); err != nil {
				return nil, err
			} else {
				datas = append(datas, r)
			}
		}
	}
	return datas, nil
}

func GetUserExpense() (float32, error) {
	query := `
		SELECT SUM(buyed_items.total_price) AS expense FROM buyed_items
		GROUP BY buyed_items.user_id
		HAVING buyed_items.user_id = ?;
	`
	var temp struct {
		Expense float32
	}
	if err := db.Raw(query, CurrUser.GetID()).Scan(&temp).Error; err != nil {
		return 0, err
	} else {
		return temp.Expense, nil
	}
}

func (p buyedItem) Insert() error {
	return db.Save(p).Error
}
