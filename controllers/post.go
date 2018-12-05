package controllers

import (
	"github.com/kataras/iris/mvc"
	"www.wooyanger.com/models"
)

type PostController struct {
	Controllers
}

func (p *PostController) GetBy(id int64) mvc.Result {
	post := p.Post.GetPostById(id)
	return mvc.View{
		Name: "post/view.html",
		Data: map[string]interface{}{
			"Title": "管理后台",
			"Post": post,
		},
	}
}

func (p *PostController) GetNew() mvc.Result {
	if p.IsLogged() {
		return mvc.View{
			Name: "post/new.html",
			Data: map[string]interface{}{
				"Title": "",
				"Authenticated": true,
				"RequireQuillPlugin": true,
			},
		}
	} else {
		return mvc.Response{
			Path: loginUrl,
			Code: 302,
		}
	}
}

func (p *PostController) PostNew() {
	newPost := &models.Post{}
	newPost.Title = p.Ctx.FormValue("title")
	newPost.Content = p.Ctx.FormValue("content")
	if err := models.CreatePost(newPost); err != nil {
		return
	}

}