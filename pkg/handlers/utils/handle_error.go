package utils_handlers

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func HandleError(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, map[string]string{
		"error": err.Error(),
	})
}
