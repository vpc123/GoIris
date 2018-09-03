# Iris的Post方法提交数据

##### 生产过程中，我们总会遇到的问题或者使用方法不是get网页内容，总会有post请求数据的方式提交数据请求，然后进行页面数据展示或者提交。		


>package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态模板文件
	app.RegisterView(iris.HTML("./templet", ".html"))
	// app.RegisterView(iris.Pug("./Pug", ".pug"))
	app.Any("/", func(ctx iris.Context) {
		title := ctx.FormValue("title")
		// title,name := ctx.FormValues("title","name")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":12345"))
}




## 分析代码
app.Any("/", func(ctx iris.Context) {
		title := ctx.FormValue("title")
		// title,name := ctx.FormValues("title","name")
		ctx.ViewData("Title", title)
		ctx.View("index.html")
	})

#### 我们使用句柄语句:指定方法为Any() ,所以无论Get()或者Post()方法都可以进行请求访问，通过title := ctx.FormValue("title")，我们获取网页提交的title数据字段的数据内容，并且定义为局部字段进行数据解析。之后通过ctx.ViewData("Title", title)指定内部字段使用方式传参方法。通过使用ctx.View("index.html")，我们可以将参数信息将参数内容传到网页进行排版显示，和数据使用。

## 总结
##### 传统的网页数据总会使用最多的两种请求方式就是post和get方法，使用这种方法我们基本可以处理所有的网页请求，另外的数据库连接仅仅只是请求的扩展。所以熟练掌握这两种请求方式，会使得在网站管理和开发的过程显得非常轻松！