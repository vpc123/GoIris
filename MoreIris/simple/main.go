package main

import (
	"time"

	"fmt"

	"github.com/kataras/iris"

	"github.com/kataras/iris/cache"
)

var markdownContents = []byte(`## Hello Markdown

This is a sample of Markdown contents
*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

`)

func main() {
	app := iris.New()
	// 开启debug模式
	app.Logger().SetLevel("debug")
	app.Get("/", cache.Handler(10*time.Second), writeMarkdown)
	app.Run(iris.Addr(":8080"))
}
func writeMarkdown(ctx iris.Context) {

	fmt.Println("句柄语句被执行!")
	ctx.Markdown(markdownContents)

}
