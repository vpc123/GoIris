# Iris模板文件指定功能
#
很多时候，网站开发的大型项目需要使用mvc框架使得整个项目易于管理使用，维护和后期升级的作用，今天我们就是开始讲解一下如何使用iris 中mvc的使用情况以及用法介绍。
## 代码详情




package main

import (
	"github.com/kataras/iris"
)

func main() {
	

- app := iris.New()
	

- // 定义静态模板文件
	


- app.RegisterView(iris.HTML("./templet", ".html"))
	

- // app.RegisterView(iris.Pug("./Pug", ".pug"))
	

- app.Get("/", func(ctx iris.Context) {
	

- 	ctx.ViewData("Title", "pppppppppppp")
	

- 	ctx.View("index.html")
	

- })

	

- app.Run(iris.Addr(":12345"))
}


### 我们通过函数注册模板文件的路径，使得访问可以快速定位文件的位置
>app.RegisterView(iris.HTML("./templet", ".html"))

### 通过控制器路由选择定位模板文件

ctx.ViewData("Title", "pppppppppppp")



## HTML  文件  index.html

>这个世界呀!
>{{.Title}}


### 页面显示图片

![](https://i.imgur.com/ZETG2oJ.png)

## 总结
2018/8/30	系统学习模板文件的使用，对iris  MVC
框架有更加深入的认知。为下一步学习打下坚实的基础！
