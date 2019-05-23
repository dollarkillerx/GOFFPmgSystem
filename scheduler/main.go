package main

import (
	"GOFFPmgSystem/scheduler/taskruner"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id",vidDelRecHandler)
	return router
}

func main() {
	go taskruner.Start()
	handlers := RegisterHandlers()
	http.ListenAndServe(":9002",handlers)
}