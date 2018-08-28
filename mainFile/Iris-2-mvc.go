package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hello world！</h1>")

	})
	app.Get("/json", func(ctx iris.Context) {
		liste := map[string]int8{"A": 1, "B": 2, "C": 3}
		ctx.JSON(liste)

	})

	app.Run(iris.Addr(":12345"))
}
