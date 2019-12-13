package models

type ICrowdfundItem interface {
}

type crowdfundItem struct {
	ID    int
	Name  string
	Stock int
	Price float32
}
