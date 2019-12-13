package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/dhpsagala/crowdfunding_donny/libs/utils"
	v "github.com/dhpsagala/crowdfunding_donny/models/views"
	"github.com/jinzhu/gorm"
)

type IUser interface {
	Insert() error
	GetPassword() string
	UpdateToken(token string) error
	GetToken() string
}

type user struct {
	gorm.Model
	Email    string
	Password string
	Token    string
}

func GetUserByEmail(email string) (IUser, error) {
	var u user
	err := db.Where(&user{Email: email}).First(&u).Error
	return &u, err
}

func CreateUser(regUser *v.RegisterUser) (IUser, error) {
	var u IUser = nil
	var err error
	if err = regUser.Validate(); err == nil {
		hPassword := utils.HashString(regUser.Password, "")
		u = &user{Email: regUser.Email, Password: hPassword}
		if err = u.Insert(); err != nil {
			return nil, err
		}
	}
	return u, err
}

func AuthenticateUser(logUser *v.LoginUser) (IUser, error) {
	var u IUser = nil
	var err error
	if err = logUser.Validate(); err == nil {
		if u, err = GetUserByEmail(logUser.Email); err == nil {
			hPassword := utils.HashString(logUser.Password, "")
			if hPassword != u.GetPassword() {
				return nil, errors.New("Invalid email or password")
			}
			token := generateToken("")
			if err = u.UpdateToken(token); err != nil {
				return nil, err
			}
		}
	}
	return u, err
}

func ValidateToken(token string) (bool, error) {
	var u user
	if err := db.Where(&user{Token: token}).First(&u).Error; err != nil {
		return false, err
	}
	uP := &u
	if uP == nil {
		return false, nil
	}
	return true, nil
}

func (u *user) Insert() error {
	return db.Save(u).Error
}

func (u *user) GetPassword() string {
	return u.Password
}

func (u *user) UpdateToken(token string) error {
	u.Token = token
	return db.Save(u).Error
}

func (u *user) GetToken() string {
	return u.Token
}

func generateToken(hash string) string {
	var timestamp string = strconv.Itoa(int(time.Now().Unix()))
	return utils.HashString(timestamp, hash)
}
