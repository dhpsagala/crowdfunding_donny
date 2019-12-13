package libs

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func writeResponse(httpStatus int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	if data != nil {
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		dataType := reflect.TypeOf(data)
		if dataType.Implements(errorInterface) {
			json.NewEncoder(w).Encode(data.(error).Error())
			return
		}
	}
	json.NewEncoder(w).Encode(data)
}

func OK(w http.ResponseWriter, data interface{}) {
	writeResponse(http.StatusOK, w, data)
}

func NotFound(w http.ResponseWriter, data interface{}) {
	writeResponse(http.StatusNotFound, w, data)
}

func BadRequest(w http.ResponseWriter, data interface{}) {
	writeResponse(http.StatusBadRequest, w, data)
}

func InternalServerError(w http.ResponseWriter, data interface{}) {
	writeResponse(http.StatusInternalServerError, w, data)
}

func Unauthorized(w http.ResponseWriter, data interface{}) {
	writeResponse(http.StatusUnauthorized, w, data)
}
