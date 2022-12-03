package middleware

import (
	"fmt"
	koa "hicky/pkg/koa/onion"
	"strconv"
	"time"
)

func Logs(ctx *koa.Context, next koa.Next) {

	date := time.Now()
	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())
	timeStr := year + "/" + month + "/" + day + " " + hour + ":" + minute + ":" + second

	method := ctx.Request.Method
	url := (*ctx.Request.URL).Path
	header := ctx.Request.Header
	body := ctx.Body

	fmt.Println(timeStr, method, url, header, body)

	next()

}
