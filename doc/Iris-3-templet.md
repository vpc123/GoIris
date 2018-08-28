# 静态页面定向访问

有时候，往往会出现页面定向问题的出现，比如网站有特殊的临时需求，暂时加入一个动态变化的页面或者是需求经常变更的情况下，总会有类似的需求产生，所以动态定向页面也是生产中特别需要的一种资源定向需求。


## 代码实例
#
>package main
>
>import (
>	"github.com/kataras/iris"
>)

>func main() {
>
>	app := iris.New()
>	
>	app.StaticWeb("/public", "./templet/Iris-3-templet.html")

>	app.Run(iris.Addr(":12345"))
>}


## 代码解析

app.StaticWeb("/public", "./templet/Iris-3-templet.html")

定义指向具体文件的模板形式,根据路由定义访问的网页页面。

![](https://i.imgur.com/yyfYvpX.png)



## 应用场景

需求经常变更的快速迭代变化需求