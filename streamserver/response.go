package main

import "net/http"

func sendErrorResponse(w http.ResponseWriter,sc int,errMsg string) {
	w.WriteHeader(sc)
	w.Write([]byte(errMsg))
}