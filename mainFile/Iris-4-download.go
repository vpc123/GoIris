package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态网页访问

	app.Get("/down", func(ctx iris.Context) {
		ctx.ContentType("发送图片文件成功!")

		ctx.SendFile("./download/1.png", "downpng.png")

	})

	app.Run(iris.Addr(":12345"))
}
