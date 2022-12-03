# Go Koa

## 介绍

> 模仿 NodeJS koa 所写的 Golang 洋葱皮式中间件。
> Golang 版的 Koa

## 使用

### 洋葱皮中间件

```go
package main

import (
	"fmt"
	koa "hicky/pkg/koa/onion"
	"log"
	"net/http"
	"strconv"
)

var app *koa.Onion

func init() {

	app.Use(func(ctx *koa.Context, next koa.Next) {
		fmt.Println("------ 1")
		next()
		fmt.Println("------ 4")
	})

	app.Use(func(ctx *koa.Context, next koa.Next) {
		fmt.Println("------ 2")
	})

	app.Use(func(ctx *koa.Context, n koa.Next) {
		fmt.Println("------ 3")
	})

}

func main() {
	http.HandleFunc("/", dispatch)
	listen(9090)
}

// 请求调度
func dispatch(res http.ResponseWriter, req *http.Request) {
	app = koa.CreateOnion(req)
	ctx := app.CallBack()
	body := ctx.Body
	res.Write([]byte(body))
}

// 监听端口
func listen(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

#### 输出

```
------ 1
------ 2
------ 4
```

### 洋葱皮路由

```go
func init() {

	route := router.CreateRouter("/api")

	route.All("/test", func(ctx *koa.Context, next koa.Next) {
		fmt.Println("route 中间件（不区分请求方法）", ctx.Request.Method, ctx.Request.URL)
	})

	route.Get("/test",
		func(ctx *koa.Context, next koa.Next) {
			fmt.Println("route 中间件1", ctx.Request.Method, ctx.Request.URL)
			next()
		},
		func(ctx *koa.Context, n koa.Next) {
			fmt.Println("route 中间件2", ctx.Request.Method, ctx.Request.URL)
		},
	)

	app.Use(router.Routes)

	app.Use(func(ctx *koa.Context, next koa.Next) {
		fmt.Println("app 中间件2", ctx.Request.URL)
	})

}
```

#### 输出

```
route 中间件（不区分请求方法） GET /api/test
route 中间件1 GET /api/test
route 中间件2 GET /api/test
app 中间件2 /api/test
```

### 路由拼接

```go
func init() {
	route := router.CreateRouter("/api")
	fileRoute := router.CreateRouter("/file")
	route.Routes(&fileRoute)

  // 访问路径 /api/test
	route.Get("/test", func(ctx *koa.Context, next koa.Next) {
		fmt.Println("test")
	})

	// 访问路径 /api/file/upload
	fileRoute.Get("/upload", func(ctx *koa.Context, next koa.Next) {
		fmt.Println("文件上传")
	})

	app.Use(router.Routes)
}
```