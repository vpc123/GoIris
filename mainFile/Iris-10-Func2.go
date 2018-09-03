package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<h1>404</h1>")
	})

	app.Macros().Int.RegisterFunc("isPair", func() func(string) bool {
		return func(paramavalue string) bool {
			eval, _ := strconv.Atoi(paramavalue)
			return eval%2 == 0
		}
	})

	app.Get("/{id:int isPair()}", func(ctx iris.Context) {
		title := ctx.Params().Get("id")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}
