package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/{title}", func(ctx iris.Context) {
		title, _ := ctx.Params().Get("title")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}
