# Iris控制器使用步骤之解析json数据
#
## 前情回顾
1. package main
1. 
1. import (
1. 	"github.com/kataras/iris"
1. )
1. 
1. func main() {
1. 	app := iris.New()
1. 	app.Handle("GET", "/", func(ctx iris.Context) {
1.		ctx.HTML("<\h1>hello world！<\ /h1>")
1. 
1. 	})
1. 	app.Run(iris.Addr(":12345"))
1. }package main


## 网页服务传输过程中总会需要网页解析json数据，Json数据在网页解析中的位置占据很大的一块，所以以下代码详细分析了json在iris中的使用

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/json", func(ctx iris.Context) {
		liste := map[string]int8{"A": 1, "B": 2, "C": 3}
		ctx.JSON(liste)

	})

	app.Run(iris.Addr(":12345"))
}


#### 访问网页返回所需要的数据类型
![](https://i.imgur.com/2Oqlp2A.png)