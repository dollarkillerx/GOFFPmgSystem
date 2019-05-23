package main

import (
	"GOFFPmgSystem/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResopnse(w,400,"vido id should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResopnse(w,500,"Internal server error")
		return
	}
	sendResopnse(w,200,"")
	return
}