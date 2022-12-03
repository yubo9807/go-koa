package middleware

import koa "hicky/pkg/koa/onion"

func Header(ctx *koa.Context, next koa.Next) {

	ctx.Request.Header.Set("Access-Control-Allow-Origin", "http://hpyyb.cn")
	next()

}
