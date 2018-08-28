# Iris基础起步第一天

#
## 安装
> go get -u github.com/kataras/iris

## 代码如下
1. package main
1. import "github.com/kataras/iris"
1. func main() {
1. 
1. 	app := iris.Default()
2. 	
1. 	app.Handle("GET", "/", func(ctxiris.Context){
1. 		ctx.HTML("HELLO World!")
1. 	})
2. 	
1. 	app.Run(iris.Addr(":8080"))
1. 	// 暴露开放的端口，使其可以被访问到
1. }


![](https://i.imgur.com/aAFfgEP.png)


### 访问网页  http://127.0.0.1:8080

![](https://i.imgur.com/1T0DHW9.png)

