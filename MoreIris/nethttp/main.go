package main

import (
	"fmt"
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	app := iris.New()
	irisMiddleware := iris.FromStd(nativeTestMiddleware)
	app.Use(irisMiddleware)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("hOME")

	})
	app.Get("/ok", func(ctx iris.Context) {
		ctx.HTML("<b>Hello world!</b>")

	})
	app.Run(iris.Addr(":8080"))

}

func nativeTestMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:" + r.URL.Path)
}
