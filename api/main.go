package main

import (
	"github.com/julienschmidt/httprouter" //rest API快速
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}


func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000",r)
}
