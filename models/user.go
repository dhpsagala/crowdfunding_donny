package models

type IUser interface {
}

type user struct {
	ID int
	Email    string
	Password string
}
