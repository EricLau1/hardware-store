package controllers

import (
	"fmt"
	"net/http"
)

func buildCreatedResponse(w http.ResponseWriter, location string) {
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func buildLocation(r *http.Request, id uint64) string {
	return fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, id)
}

func buildDeleteResponse(w http.ResponseWriter, entity_id uint64) {
	w.Header().Set("Entity", fmt.Sprint(entity_id))
	w.WriteHeader(http.StatusNoContent)
}
