package controllers

import (
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
	Controllers
}

func (h *HomeController) Get() mvc.Result {
	postList := h.Post.GetAllPost()
	siteName := h.Config.GetSiteName()
	return mvc.View{
		Name: "home.html",
		Data: map[string]interface{}{
            "Title": "Home",
			"Posts": postList,
			"Authenticated": h.IsLogged(),
            "SiteName": siteName,
		},
	}
}
