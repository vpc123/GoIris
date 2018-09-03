# Iris指定参数名称还有参数类型
# 
## 应用场景
我们总会产生这样或者那样的需求，我们总需要对文章或者文档归档分类。所以划分文档归类域就显得十分必要。但是有时需要限制文档提交的参数还有类型就可以完美的处理好这些功能。下面我们就会开始介绍关于post参数的必要条件还有参数限制。

## 代码源码

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Get("/{ent}", func(ctx iris.Context) {
		// app.Get("/{ent:int}", func(ctx iris.Context) {
		title, _ := ctx.Params().GetInt("ent")
		ctx.ViewData("Title", title*12)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}


## 源码框架参数解析
	app.Get("/{ent}", func(ctx iris.Context) {
		// app.Get("/{ent:int}", func(ctx iris.Context) {
		title, _ := ctx.Params().GetInt("ent")
		ctx.ViewData("Title", title*12)
		ctx.View("index.html")
	})

#  
# 

我们在get请求中提交请求参数ent，之后在页面中通过函数方法得到具体参数通过键值对的方式使用指定参数对应。在前台进行页面的指定。


## 分析总结
# 

我们在学习过程中要注意这些潜在的细节问题，在处理问题时尽可能关系底层，一个框架就是由底层方法实现的模式。学习不可以不注重基础，只追求新颖。