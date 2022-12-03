package router

import koa "hicky/pkg/koa/onion"

// 所有请求方式都会经过
func (r Route) All(path Path, middlewareList ...koa.Middleware) {
	key := "ALL" + r.baseURL + path
	performQueue[key] = middlewareList
}

// GET 请求中间件
func (r Route) Get(path Path, middlewareList ...koa.Middleware) {
	key := "GET" + r.baseURL + path
	performQueue[key] = middlewareList
}

// POST 请求中间件
func (r Route) Post(path Path, middlewareList ...koa.Middleware) {
	key := "POST" + r.baseURL + path
	performQueue[key] = middlewareList
}

// DELETE 请求中间件
func (r Route) Delete(path Path, middlewareList ...koa.Middleware) {
	key := "DELETE" + r.baseURL + path
	performQueue[key] = middlewareList
}

// PUT 请求中间件
func (r Route) Put(path Path, middlewareList ...koa.Middleware) {
	key := "PUT" + r.baseURL + path
	performQueue[key] = middlewareList
}

// OPTIONS 请求中间件
func (r Route) Options(path Path, middlewareList ...koa.Middleware) {
	key := "OPTIONS" + r.baseURL + path
	performQueue[key] = middlewareList
}
