package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any ) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	res, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Write(res)
	return nil
}
