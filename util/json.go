package util

import (
	"encoding/json"
	"net/http"
)

func WriteErrorHTTP(rw http.ResponseWriter, err error, status int) {
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]string{
		"error": err.Error(),
	})
}

func WriteJSONHTTP(rw http.ResponseWriter, js interface{}, status int) {
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(js)
}
