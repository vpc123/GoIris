# Iris 条件判断执行
# 
# 

## 源码

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	app.Macros().Int.RegisterFunc("isPair", func() func(string) bool {
		return func(paramavalue string) bool {
			eval, _ := strconv.Atoi(paramavalue)
			return eval%2 == 0
		}
	})
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/{id isPair()}", func(ctx iris.Context) {
		// app.Get("/{ent:int}", func(ctx iris.Context) {
		title, _ := ctx.Params().GetInt("id")
		ctx.ViewData("Title", title*12)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}



## 代码分析
匹配函数返回值，判断函数的可用性。
