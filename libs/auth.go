package libs

import (
	"net/http"
	"strings"

	"github.com/dhpsagala/crowdfunding_donny/models"
)

func AuthUser(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Authorization"]) == 0 {
			Unauthorized(w, "Unauthorized")
			return
		}

		tokenParts := strings.Split(r.Header["Authorization"][0], " ")

		if len(tokenParts) != 2 {
			BadRequest(w, "Invalid authorization header format")
			return
		}

		if valid, err := models.ValidateToken(tokenParts[1]); err != nil {
			InternalServerError(w, err)
			return
		} else {
			if !valid {
				Unauthorized(w, "Unauthorized")
				return
			}
		}
		next(w, r)
	})
}
