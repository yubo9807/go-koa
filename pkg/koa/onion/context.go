package koa

import "net/http"

// 上下文对象，不够用自己往里添加
type Context struct {
	State    State
	Body     string
	Request  *http.Request
	Response http.ResponseWriter
}
