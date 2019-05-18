package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler  {
	handler := middleWareHandler{}
	handler.r = r
	return handler
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w,r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/login",LoginUser)
	return router
}

func main() {
	route := RegisterHandlers()
	handler := NewMiddleWareHandler(route)
	http.ListenAndServe(":8580",handler)
}

//main->middleware->defs(message,err)->handles->dbops->response
