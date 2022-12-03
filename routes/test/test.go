package routes

import (
	"fmt"
	koa "hicky/pkg/koa/onion"
	"hicky/pkg/koa/router"
)

func Test(prentRoute router.Route) {

	route := router.CreateRouter("/test")
	prentRoute.Use(&route) // 先挂载路由，再添加路由中间件

	route.Get("", func(ctx *koa.Context, next koa.Next) {
		fmt.Println(ctx.Request.URL)
	})

	route.Get("/111", func(ctx *koa.Context, next koa.Next) {
		fmt.Println(ctx.Request.URL)
	})

}
