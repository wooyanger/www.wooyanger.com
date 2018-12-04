package controllers

import "github.com/kataras/iris/mvc"

type PostController struct {
	Controllers
}

func (c *PostController) GetBy(id int64) mvc.Result {
	p := c.Post.Get(id)
	return mvc.View{
		Name: "post/view.html",
		Data: map[string]interface{}{
			"Title": "管理后台",
			"Post": p,
		},
	}
}