package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status int, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "status": status})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"data": data, "status": status})
}