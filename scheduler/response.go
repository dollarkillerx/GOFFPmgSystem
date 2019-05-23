package main

import "net/http"

func sendResopnse(w http.ResponseWriter,sc int,resp string)  {
	w.WriteHeader(sc)
	w.Write([]byte(resp))
}