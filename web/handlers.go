package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	cname, e1 := r.Cookie("username")
	cid, e2 := r.Cookie("session")
	if e1 != nil || e2 != nil {
		p := &HomePage{Name:"dollarKiller"}
		t, e := template.ParseFiles("../template/home.html")//编译
		if e != nil {
			log.Printf("Parseing template home.html error:%s",e)
			return
		}
		t.Execute(w,p)//放入
		return
	}

	if len(cname.Value) != 0 && len(cid.Value) != 0 {
		http.Redirect(w,r,"/userhome",http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	cookie, e1 := r.Cookie("username")
	_, e2 := r.Cookie("session")

	if e1 != nil || e2 != nil {
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var p *UserPage
	if len(cookie.Value) != 0 {
		p = &UserPage{Name:cookie.Value}
	}else if len(fname) != 0 {
		p = &UserPage{Name:fname}
	}
	t, e := template.ParseFiles("../template/userhome.html")
	if e != nil {
		log.Printf("Parsing userHome.html error:%s",e)
		return
	}
	t.Execute(w,p)
}

func apiHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	if r.Method != http.MethodPost {
		res, _ := json.Marshal(ErrorRequestNotRecognized)
		w.WriteHeader(400)
		w.Write(res)
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res,apibody);err != nil {
		res, _ := json.Marshal(ErrorRequestBodyParseFailed)
		w.WriteHeader(400)
		w.Write(res)
		return
	}

	request(apibody,w,r)
	defer r.Body.Close()
}