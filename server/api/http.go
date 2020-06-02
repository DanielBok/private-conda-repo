package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func toJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "api/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(object); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func readJson(r *http.Request, object interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		return err
	}
	return nil
}

func ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprint(w, "Okay"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
