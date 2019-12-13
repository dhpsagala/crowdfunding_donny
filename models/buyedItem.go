package models

type IBuyedItem interface {
}

type buyedItem struct {
	ItemID     int
	UserID     int
	Quantity   int
	TotalPrice float32
}
