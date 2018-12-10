package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"www.wooyanger.com/controllers"
	"www.wooyanger.com/pkg/setting"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	session := sessions.New(setting.SessionCfg)
	ViewEngine := iris.HTML("./templates", ".html")
	app.RegisterView(ViewEngine.Reload(true))
	app.StaticWeb("/public", "./public")
	mvc.New(app).Register(session.Start).Handle(new(controllers.HomeController))
	mvc.New(app.Party("/console")).Register(session.Start).Handle(new(controllers.ConsoleController))
	mvc.New(app.Party("/posts")).Register(session.Start).Handle(new(controllers.PostController))
	mvc.New(app.Party("/file")).Handle(new(controllers.FileController))
	app.Run(iris.Addr(setting.ListenAddr))
}