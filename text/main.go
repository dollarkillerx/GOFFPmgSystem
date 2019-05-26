package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
创建路由注册
 */
func RegisterRouter() *httprouter.Router {
	router := httprouter.New() //得到router实例

	//注册路由
	// 路由后面可以跟.GET .POST .DELETE .PUT 等  (其他的怎么查看呢? 你按住ctrl 点赞 New()进入方法 看下面的代码 里面就是详情啊!)
	// 第一个你发觉参数有些不同的吗?  对 多了一个params httprouter.Params 这个用来接收路由中的参数  等会来讲

	// 我们先看GET POST都指向同一个路径 but返回不同的内容,  怎么测试呢?  百度postman  这个是测试基础啊!百度学习吧骚年 非常简单的

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("this is get method"))
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("this is post method"))
	})

	return router
}

func main() {
	router := RegisterRouter() //注册路由

	fmt.Println("server is runing ...")

	//这里发现改变了什么吗?
	//对第二次参数变成了router
	err := http.ListenAndServe(":8085", router) //第一个是地址(ip:端口 ip可以省略 监听本机全部端口) 第二个是handler

	if err != nil {
		fmt.Println("server error:",err.Error())
	}
}
