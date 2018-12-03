package controllers

import (
	"github.com/kataras/iris/mvc"
	"www.wooyanger.com/models"
)

const authFatalMsgKey  = "AuthFatalMsg"
const loginUrl  = "/console/login"
const consoleUrl  = "/console/index"

type ConsoleController struct {
	Controllers
	User			*models.User
}

func (c *ConsoleController) GetLogin() mvc.Result {
	if c.IsLogged() {
		return mvc.Response{
			Path: consoleUrl,
			Code: 302,
		}
	}
	return mvc.View{
		Name: "console/login.html",
		Data: map[string]interface{}{
			"Title": "管理后台",
		},
	}
}

func (c *ConsoleController) PostLogin() {
	username := c.Ctx.FormValue("username")
	password := c.Ctx.FormValue("password")
	user, err := models.UserLogin(username, password)
	if err != nil {
		c.Session.SetFlash(authFatalMsgKey, err)
		c.Ctx.Redirect(loginUrl, 302)
		return
	} else {
		c.LoginUser(user.Id)
		c.Ctx.Redirect(consoleUrl, 302)
		return
	}
}

func (c *ConsoleController) GetIndex() mvc.Result {
	if c.IsLogged() {
		return mvc.View{
			Name: "console/index.html",
		}
	}
	return mvc.Response{
		Path: loginUrl,
		Code: 302,
	}
}