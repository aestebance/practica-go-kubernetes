package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSONResponse Creation response in JSON
func JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(PrettyPrint(body))
}

// PrettyPrint is used to print the response pretty
func PrettyPrint(b []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return nil
	}
	return out.Bytes()
}
