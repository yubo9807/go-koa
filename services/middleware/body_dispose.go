package middleware

import (
	"fmt"
	koa "hicky/pkg/koa/onion"
)

func BodyDispose(ctx *koa.Context, next koa.Next) {

	fmt.Println("返回数据处理")
	next()

}
