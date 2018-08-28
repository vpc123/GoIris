package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态网页访问
	app.StaticWeb("/public", "./templet/Iris-3-templet.html")

	app.Run(iris.Addr(":12345"))
}
