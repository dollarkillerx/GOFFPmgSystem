package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/login",LoginUser)
	return router
}

func main() {
	route := RegisterHandlers()

	http.ListenAndServe(":8580",route)
}