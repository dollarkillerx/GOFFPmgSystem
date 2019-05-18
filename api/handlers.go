package main

import (
	"GOFFPmgSystem/api/dbops"
	"GOFFPmgSystem/api/defs"
	"GOFFPmgSystem/api/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	//读取body
	bytes, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(bytes, ubody);err != nil {
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if 	err := dbops.AddUserCredential(ubody.Username, ubody.Pwd);err != nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}
	if marshal, e := json.Marshal(su);e != nil {
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(w,string(marshal),201)
	}
}

func LoginUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	userName := p.ByName("user_name")
	w.Write([]byte("youName:" + userName))
}