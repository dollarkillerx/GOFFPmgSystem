package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandle struct {
	r *httprouter.Router
	l *ConnLimiter //加入流控模块
}

func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler {
	handle := middleWareHandle{}
	handle.r = r
	handle.l = NewConnLimiter(cc)
	return handle
}

// 流控核心
func (m middleWareHandle) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	// 从桶中获得令牌
	if !m.l.GetConn() {
		sendErrorResponse(w,http.StatusTooManyRequests,"Too many requests")
		return
	}
	m.r.ServeHTTP(w,r)
	defer func() {
		//当链接结束 令牌返回令牌桶中
		m.l.ReleaseConn()
	}()
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router
}

func main() {
	handler := RegisterHandler()
	// 接替
	wareHandler := NewMiddleWareHandler(handler, 2)
	http.ListenAndServe(":9003",wareHandler)
}