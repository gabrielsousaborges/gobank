package config

import (
	"encoding/json"
	"net/http"
)

var (
	PostgresConnection = ""
	Port               = 0
	SecretKey          []byte
)

func RespondWithError(w http.ResponseWriter, code, errorCode int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "code": errorCode})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
