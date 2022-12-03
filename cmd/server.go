package main

import (
	koa "hicky/pkg/koa/onion"
	"hicky/services/middleware"
	"log"
	"net/http"
	"strconv"
)

var app *koa.Onion

// 初始化添加应用中间件
func init() {

	// 请求头设置
	app.Use(middleware.Header)

	// 添加路由
	app.Use(middleware.Routes)

	// 返回 body 处理
	app.Use(middleware.BodyDispose)

	// 日志记录
	app.Use(middleware.Logs)

}

// 启动服务
func main() {
	http.HandleFunc("/", dispatch)
	listen(20020)
}

// 请求调度
func dispatch(res http.ResponseWriter, req *http.Request) {
	app = koa.CreateOnion(req, res)
	ctx := app.CallBack()
	res.Write([]byte(ctx.Body))
}

// 监听端口
func listen(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
