# Iris文件模板下载功能
#
有时候，网站需要频繁向用户提供文件下载功能，或视频，或图片，或者其他可以通过网络传输的资源类型。
## 文件资源下载目录
![](https://i.imgur.com/w6ibL8p.png)


## 代码层次

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 定义静态网页访问

	app.Get("/down", func(ctx iris.Context) {
		ctx.ContentType("发送图片文件成功!")

		ctx.SendFile("./download/1.png", "downpng.png")

	})

	app.Run(iris.Addr(":12345"))
}


## 代码分析

使用SendFile（）函数调用，定义发送新文件的地址以及地址名称

>ctx.SendFile("./download/1.png", "downpng.png") 

## 总结:
###### 针对网页文件的传输，我们使用ctx中的函数方法，可以定义传输文档路径和请求文件的接受者所接受的文件名称。