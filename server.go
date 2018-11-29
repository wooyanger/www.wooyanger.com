package main

import (
	"github.com/kataras/iris"
	"www.wooyanger.com/pkg/setting"
)

func main() {
	app := iris.New()
	ViewEngine := iris.HTML("./templates", ".html")
	app.RegisterView(ViewEngine.Reload(true))
	app.StaticWeb("/public", "./public")
	app.Get("/", func(ctx iris.Context){
		ctx.View("index.html")
	})
	app.Run(iris.Addr(setting.ListenAddr))
}