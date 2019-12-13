package controllers

import (
	"encoding/json"
	"github.com/dhpsagala/crowdfunding_donny/libs"
	"github.com/dhpsagala/crowdfunding_donny/models"
	"github.com/dhpsagala/crowdfunding_donny/models/views"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	regUser := &views.RegisterUser{}
	if err := json.NewDecoder(r.Body).Decode(regUser); err != nil {
		libs.OK(w, err.Error())
		return
	}
	if u, err := models.CreateUser(regUser); err != nil {
		libs.BadRequest(w, err)
	} else {
		libs.OK(w, u)
	}
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	logUser := &views.LoginUser{}
	if err := json.NewDecoder(r.Body).Decode(logUser); err != nil {
		libs.OK(w, err.Error())
		return
	}
	if u, err := models.AuthenticateUser(logUser); err != nil {
		libs.BadRequest(w, err)
	} else {
		libs.OK(w, u.GetToken())
	}
}

func UserTransaction(w http.ResponseWriter, r *http.Request) {
	if d, err := models.GetUserTransaction(); err != nil {
		libs.InternalServerError(w, err)
	} else {
		libs.OK(w, d)
	}
}

func UserExpense(w http.ResponseWriter, r *http.Request) {
	if d, err := models.GetUserExpense(); err != nil {
		libs.InternalServerError(w, err)
	} else {
		libs.OK(w, d)
	}
}
