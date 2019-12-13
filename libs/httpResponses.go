package libs

import (
	"encoding/json"
	"net/http"
)

func String(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(message))
}

func Json(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
