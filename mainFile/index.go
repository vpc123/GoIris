package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("HELLO World!")
	})
	app.Run(iris.Addr(":8080"))
	// 暴露开放的端口，使其可以被访问到
}
