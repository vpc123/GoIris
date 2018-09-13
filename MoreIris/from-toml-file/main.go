package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello vpc!</b>")

	})

	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./configs/iris.yaml")))
}
