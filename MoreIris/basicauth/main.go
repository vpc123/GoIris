package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()
	// 开启自动认证功能，键入用户名和密码用于登录验证
	authConfig := basicauth.Config{
		Users:   map[string]string{"vpc": "123", "qwe": "123"},
		Realm:   "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)
	// 重
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/admin")

	})

	needAuth := app.Party("/admin", authentication)
	{
		// http://localhost:8080/admin
		needAuth.Get("/", h)
		needAuth.Get("/profile", h)
		needAuth.Get("/settings", func(ctx iris.Context) {
			ctx.Text("Hello World vpc!")

		})

	}
	return app
}
func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef(ctx.Path(), username, password)
}

func main() {
	app := newApp()
	// open http://localhost:8080/admin
	app.Run(iris.Addr(":8080"))

}
