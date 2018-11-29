package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"www.wooyanger.com/controllers"
	"www.wooyanger.com/pkg/setting"
)

func main() {
	app := iris.New()
	ViewEngine := iris.HTML("./templates", ".html")
	app.RegisterView(ViewEngine.Reload(true))
	app.StaticWeb("/public", "./public")
	mvc.New(app).Handle(new(controllers.HomeController))
	app.Run(iris.Addr(setting.ListenAddr))
}