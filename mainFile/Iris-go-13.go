package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var sessconf = sessions.Config{Cookie: "Liu"}
var sess = sessions.New(sessconf)

func main() {

	app = iris.New()

	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))

	app.Get("/login", func(ctx iris.Context) {
		sess.Start(ctx).Set("authoring", true)
		ctx.HTML("login！")
	})
	app.Get("/logout", func(ctx iris.Context) {
		sess.Start(ctx).Set("authoring", false)
		ctx.HTML("logout！")
	})
	app.Get("/membre", func(ctx iris.Context) {
		authoring, _ := sess.Start(ctx).GetBoolean(authoring)
		if !authoring {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.HTML("Nope !!!!!")
		} else {
			ctx.HTML("Falid!")
		}
	})

	app.Run(iris.Addr(":12345"))
}
