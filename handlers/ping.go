package handlers

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	HTTPResponse(w, http.StatusOK, []byte("pong"))
}
