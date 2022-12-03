/* 洋葱皮中间件 */
package koa

import (
	"net/http"
)

type State = map[string]interface{}
type Next = func()
type Onion struct {
	ctx Context
}

type Middleware = func(*Context, Next)

// 执行队列
var performQueue = []Middleware{}

// 创建一个洋葱皮模型
func CreateOnion(req *http.Request, res http.ResponseWriter) *Onion {
	o := &Onion{
		ctx: Context{
			State:    State{},
			Request:  req,
			Response: res,
			Body:     "Not found",
		},
	}
	return o
}

// 添加中间件
func (o *Onion) Use(fn Middleware) {
	performQueue = append(performQueue, fn)
}

// 添加完中间件之后回调
func (o *Onion) CallBack() Context {
	if len(performQueue) > 0 {
		o.compose(0, performQueue)
	}
	return o.ctx
}

// 执行中间件队列
func (o *Onion) compose(i int, middlewareList []Middleware) {
	fn := middlewareList[i]
	if i == len(middlewareList)-1 {
		fn(&o.ctx, func() {})
		return
	}
	fn(&o.ctx, func() {
		i++
		o.compose(i, middlewareList)
	})
}
