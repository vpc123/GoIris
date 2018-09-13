package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	app := iris.New()
	irisMiddleware := iris.FromStd(negronilikeTestMiddleware)
	app.Use(irisMiddleware)
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Home </h1>")

	})
	app.Get("/ok", func(ctx iris.Context) {
		ctx.Writef("Hello world！")

	})
	app.Run(iris.Addr(":8080"))

}

func negronilikeTestMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/ok" && r.Method == "GET" {
		w.Write([]byte("OK. "))
		next(w, r) // go to the next route's handler
		return
	}
	// else print an error and do not forward to the route's handler.
	w.WriteHeader(iris.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
