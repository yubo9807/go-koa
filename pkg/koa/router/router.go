// 洋葱皮路由
package router

import (
	koa "hicky/pkg/koa/onion"
)

type Route struct {
	baseURL string
}
type Path = string

// 创建一个路由
// @param baseURL 基础路径
func CreateRouter(baseURL string) Route {
	return Route{
		baseURL: baseURL,
	}
}

var performQueue = make(map[string][]koa.Middleware)

// 执行相应的中间件
func Routes(ctx *koa.Context, next koa.Next) {
	url := *ctx.Request.URL
	method := ctx.Request.Method

	allQueue := performQueue["ALL"+url.Path]
	queue := performQueue[method+url.Path]

	if len(allQueue) > 0 {
		compose(0, allQueue, ctx)
	}
	if len(queue) > 0 {
		compose(0, queue, ctx)
	}
	next()
}

// 执行中间件队列
func compose(i int, middlewareList []koa.Middleware, ctx *koa.Context) {
	fn := middlewareList[i]
	if i == len(middlewareList)-1 {
		fn(ctx, func() {})
		return
	}
	fn(ctx, func() {
		i++
		compose(i, middlewareList, ctx)
	})
}

// 将子路由挂载到父路由下
func (r Route) Use(subRoute *Route) {
	subRoute.baseURL = r.baseURL + subRoute.baseURL
}

func AllRoutes() map[string][]koa.Middleware {
	newPerformQueue := make(map[string][]koa.Middleware)
	for key, value := range performQueue {
		newPerformQueue[key] = value
	}
	return newPerformQueue
}
