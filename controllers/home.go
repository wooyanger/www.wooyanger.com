package controllers

import "github.com/kataras/iris/mvc"

type HomeController struct {
	Controllers
}

func (h *HomeController) Get() mvc.Result {
	postList := h.Post.GetAll()
	return mvc.View{
		Name: "home.html",
		Data: map[string]interface{}{
			"Posts": postList,
		},
	}
}