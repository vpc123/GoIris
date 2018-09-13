package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))

	app.Get("/", avent, func(ctx iris.Context) {
		ctx.HTML("<h1>你好!世界ing</h1>")
		ctx.Next()
	}, apres)
	app.Get("/test", func(ctx iris.Context) {
		ctx.HTML("<h1>你好!test</h1>")
		ctx.Next()
	})

	app.Run(iris.Addr(":12345"))
}

func avent(ctx iris.Context) {
	ctx.HTML("<h1>你好!版头</h1>")
	ctx.Next()

}
func apres(ctx iris.Context) {
	ctx.HTML("<h1>你好!版权。</h1>")

}
