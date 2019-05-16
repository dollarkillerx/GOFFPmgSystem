package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	w.Write([]byte("Hello Golang"))
}

func LoginUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	userName := p.ByName("user_name")
	w.Write([]byte("youName:" + userName))
}