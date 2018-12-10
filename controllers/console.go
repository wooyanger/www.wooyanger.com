package controllers

import (
	"github.com/kataras/iris/mvc"
	"www.wooyanger.com/models"
)

const (
	authFatalMsgKey  = "AuthFatalMsg"
	loginUrl  = "/console/login"
	consolePostUrl  = "/console/posts"
)

type ConsoleController struct {
	Controllers
	User			*models.User
	Tags			*models.Tag
}

func (c *ConsoleController) GetLogin() mvc.Result {
	if c.IsLogged() {
		return mvc.Response{
			Path: consolePostUrl,
			Code: 302,
		}
	}
	flashMsg := c.Session.GetFlash(authFatalMsgKey)
	return mvc.View{
		Name: "console/login.html",
		Data: map[string]interface{}{
			"Title": "管理后台",
			"AuthFatalMsg": flashMsg,
			"RequireConsolePlugin": true,
		},
	}
}

func (c *ConsoleController) GetLogout() {
	c.Session.Destroy()
	c.Ctx.Redirect(loginUrl, 302)
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
		c.Ctx.Redirect(consolePostUrl, 302)
		return
	}
}

func (c *ConsoleController) GetPosts() mvc.Result {
	if c.IsLogged() {
		posts := c.Post.GetAllPost()
		return mvc.View{
			Name: "console/index.html",
			Data: map[string]interface{}{
				"Title": "管理后台",
				"Posts": posts,
			},
		}
	}
	return mvc.Response{
		Path: loginUrl,
		Code: 302,
	}
}

func (c *ConsoleController) GetTags() mvc.Result {
	if c.IsLogged() {
		tags := c.Tags.GetAllTag()
		return mvc.View{
			Name: "console/tags.html",
			Data: map[string]interface{}{
				"Title": "管理后台",
				"Tags": tags,
			},
		}
	}
	return mvc.Response{
		Path: loginUrl,
		Code: 302,
	}
}

func (c *ConsoleController) GetSystem() mvc.Result {
	if c.IsLogged() {
		siteName := c.Config.GetSiteName()
		return mvc.View{
			Name: "console/system.html",
			Data: map[string]interface{}{
				"Title": "管理后台",
				"SiteName": siteName,
			},
		}
	}
	return mvc.Response{Path: loginUrl, Code: 302,}
}

func (c *ConsoleController) PostSystem() {
	if c.IsLogged() {
		siteName := c.Ctx.PostValue("site-name")
		c.Ctx.Writef(siteName)
		return
	}
	c.Ctx.Redirect(loginUrl, 302)
}