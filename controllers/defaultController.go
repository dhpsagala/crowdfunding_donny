package controllers

import (
	"github.com/dhpsagala/crowdfunding_donny/libs"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	libs.String(w, "Hello")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	libs.String(w, "Healthy")
}
