package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := VIDOE_DIR + vid

	file, e := os.Open(vl)
	if e != nil {
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}
	w.Header().Set("Content-Type","video/mp4")
	// 当返回二进制流文件
	http.ServeContent(w,r,"",time.Now(),file)

	defer file.Close()

}

func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//检测文件
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);err != nil{
		log.Printf("Error when try")
		sendErrorResponse(w,http.StatusBadRequest,"File is to big")
		return
	}

	file, _, e := r.FormFile("file") //<form name="file">
	if e != nil {
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		log.Printf("Read file error:%v",e)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}
	fn := p.ByName("vid-id")
	err := ioutil.WriteFile(VIDOE_DIR + fn,bytes,0666)

	if err != nil {
		log.Printf("Write file error:%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("uploaded successfully"))
}