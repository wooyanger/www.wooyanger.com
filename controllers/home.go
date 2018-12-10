package controllers

import (
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
	Controllers
}

func (h *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home.html",
		Data: map[string]interface{}{
            "Title": "Home",
			"Posts": h.Post.GetAllPost(),
			"SiteName": h.Config.GetSiteName(),
            "IntroHeader": "",
            "IntroContent": h.Config.GetHomeIntroContent(),
			"Authenticated": h.IsLogged(),
		},
	}
}
