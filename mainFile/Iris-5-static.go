package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Title", "pppppppppppp")
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}
