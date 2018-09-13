package main

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/kataras/iris"

	"github.com/getsentry/raven-go"
)

func irisRavenMiddleware(ctx iris.Context) {
	w, r := ctx.ResponseWriter(), ctx.Request()

	defer func() {
		if rval := recover(); rval != nil {
			debug.PrintStack()
			rvalStr := fmt.Sprint(rval)
			packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
			raven.Capture(packet, nil)
			w.WriteHeader(iris.StatusInternalServerError)
		}
	}

	ctx.Next()
}

// https://docs.sentry.io/clients/go/integrations/http/
func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func main() {
	app := iris.New()
	app.Use(irisRavenMiddleware)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hi")
	})

	app.Run(iris.Addr(":8080"))
}
