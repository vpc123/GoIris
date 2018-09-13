package main

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

type myXML struct {
	Result string "xml"
}

func main() {
	app := iris.New()
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})

	app.Use(irisyaag.New())
	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"result": "Hello World"})
	})
	app.Get("/plain", func(ctx iris.Context) {
		ctx.Text("Hello World!")
	})
	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(myXML{Result: "Hello World！"})
	})
	app.Get("/complex", func(ctx iris.Context) {
		value := ctx.URLParam("key")
		ctx.JSON(iris.Map{"value": value})
	})
	app.Run(iris.Addr(":8080"))

}
