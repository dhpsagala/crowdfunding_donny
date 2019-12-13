package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhpsagala/crowdfunding_donny/libs"
	"github.com/dhpsagala/crowdfunding_donny/models"
	"github.com/dhpsagala/crowdfunding_donny/models/views"
	"github.com/gorilla/mux"
)

func ListOfAvailableItems(w http.ResponseWriter, r *http.Request) {
	if items, err := models.GetAvailableCroudfundItems(); err != nil {
		libs.InternalServerError(w, err)
	} else {
		libs.OK(w, items)
	}
}

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	buyProduct := &views.BuyProduct{}
	if err := json.NewDecoder(r.Body).Decode(buyProduct); err != nil {
		libs.BadRequest(w, err.Error())
		return
	}
	if v, err := strconv.Atoi(queries["id"]); err != nil {
		libs.BadRequest(w, err.Error())
		return
	} else {
		if i, err := models.GetCroudfundItem(v); err != nil {
			libs.BadRequest(w, err.Error())
			return
		} else {
			if ib, err := i.Buy(buyProduct.Quantity); err != nil {
				libs.BadRequest(w, err)
				return
			} else {
				libs.OK(w, ib)
			}
		}
	}
}
