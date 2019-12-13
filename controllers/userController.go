package controllers

import (
	"github.com/dhpsagala/crowdfunding_donny/libs"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	libs.String(w, "CreateUser")
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	libs.String(w, "AuthenticateUser")
}
