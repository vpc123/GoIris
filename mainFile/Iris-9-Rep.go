package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/{id regexp([0-9][A-Z])}", func(ctx iris.Context) {
		title := ctx.Params().Get("id")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}
