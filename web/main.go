package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/",homeHandler)
	router.POST("/",homeHandler)

	router.GET("/userhome",userHomeHandler)
	router.POST("/userhome",userHomeHandler)

	router.POST("/api",apiHandler)

	router.ServeFiles("/statics/*filepath",http.Dir("../template"))

	return router
}

func main() {
	handler := RegisterHandler()
	http.ListenAndServe(":9004",handler)
}