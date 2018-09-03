# Iris使用正则方式指定提交参数
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
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/{id regexp([0-9][A-Z])}", func(ctx iris.Context) {
		title := ctx.Params().Get("id")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}



## 分析
app.Get("/{id regexp([0-9][A-Z])}", func(ctx iris.Context) {
		title := ctx.Params().Get("id")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

通过在参数后面指定正则规则进行匹配：
regexp([0-9][A-Z])   
**更多正则规则请参考，正则表达式的使用方法！**


## 总结

作者: 流火夏梦        /杭州/        2018/9/3 