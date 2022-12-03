package middleware

import (
	koa "hicky/pkg/koa/onion"
	"hicky/pkg/koa/router"
	routes "hicky/routes/test"
)

func Routes(ctx *koa.Context, next koa.Next) {

	route := router.CreateRouter("/api")

	routes.Test(route)

	router.Routes(ctx, next)

}
