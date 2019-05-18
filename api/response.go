package main

import (
	"GOFFPmgSystem/api/defs"
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErroResponse) {
	w.WriteHeader(errResp.HttpSC)
	bytes, _ := json.Marshal(&errResp.Error)
	w.Write(bytes)
}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int) {
	w.WriteHeader(sc)
	w.Write([]byte(resp))
}