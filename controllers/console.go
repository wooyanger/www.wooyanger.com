package controllers

import "github.com/kataras/iris/mvc"

type ConsoleController struct {
	Controllers
}

func (c *ConsoleController) Get() mvc.Result {
	return mvc.View{
		Name: "console/login.html",
	}
}