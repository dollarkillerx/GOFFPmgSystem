package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init()  {
	httpClient = &http.Client{}
}

func request(body *ApiBody,w http.ResponseWriter,r *http.Request)  {
	var resp *http.Response
	var err error

	switch body.Method {
		case http.MethodGet:
			newRequest, _ := http.NewRequest("GET", body.Url, nil)
			newRequest.Header = r.Header
			resp, err = httpClient.Do(newRequest)
			if err != nil {
				log.Printf(err.Error())
				return
			}
			normalResponse(w,resp)
	case http.MethodPost:
		newRequest, _ := http.NewRequest("GET", body.Url, bytes.NewBuffer([]byte(body.ReqBody)))
		newRequest.Header = r.Header
		resp, err = httpClient.Do(newRequest)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		normalResponse(w,resp)
	case http.MethodDelete:
		newRequest, _ := http.NewRequest("Delete", body.Url, nil)
		newRequest.Header = r.Header
		resp, err = httpClient.Do(newRequest)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		normalResponse(w,resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: BadRequest"))
		return
	}

}

func normalResponse(w http.ResponseWriter,r *http.Response) {
	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		marshal, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		w.Write(marshal)
		return
	}
	w.WriteHeader(r.StatusCode)
	w.Write(bytes)
}