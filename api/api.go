package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	code    int
	message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{code: code, message: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal error", http.StatusInternalServerError)
	}
)
