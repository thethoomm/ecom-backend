package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ParseBody(w http.ResponseWriter, r *http.Request, to any) error {
	if r.Body == nil {
		return ErrEmptyBody
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(to); err != nil {
		return err
	}

	return nil
}
