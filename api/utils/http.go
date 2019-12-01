package utils

import "net/http"

import "encoding/json"

func WriteAsJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	WriteAsJson(w, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}
