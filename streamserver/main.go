package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router
}

func main() {
	handler := RegisterHandler()

	http.ListenAndServe(":9000",handler)
}