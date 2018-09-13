package main

import (
	"time"

	"github.com/kataras/iris"
)

const refreshEvery = 10 * time.Second

func main() {
	app := iris.New()
	// 加入高速缓存中
	app.Use(iris.Cache304(refreshEvery))
	// same as:
	// app.Use(func(ctx iris.Context) {
	// 	now := time.Now()
	// 	if modified, err := ctx.CheckIfModifiedSince(now.Add(-refresh)); !modified && err == nil {
	// 		ctx.WriteNotModified()
	// 		return
	// 	}

	// 	ctx.SetLastModified(now)

	// 	ctx.Next()
	// })
	app.Get("/", greet)
	app.Run(iris.Addr(":8080"))

}

func greet(ctx iris.Context) {
	ctx.Header("X-Custom", "my custom header")
	ctx.Writef("Hello World!", time.Now())

}
