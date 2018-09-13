package main

import (
	"github.com/kataras/iris"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, value)

		ctx.Writef("cookie added: %s = %s", name, value)
	})
	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name) // <--
		// If you want more than the value then:
		// cookie, err := ctx.Request().Cookie(name)
		// if err != nil {
		//  handle error.
		// }

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))

		ctx.Writef("cookie %s removed", name)
	})

	return app

}

func main() {
	app := newApp()

	app.Run(iris.Addr(":8080"))
}
