package controllers

import (
	"github.com/kataras/iris/mvc"
	"strconv"
	"time"
	"www.wooyanger.com/models"
)

type PostController struct {
	Controllers
	Tags			models.Tag
}

const (
	CreateFatalMsgKey = "CreateFatalMsg"
)

func (p *PostController) GetBy(id int64) mvc.Result {
	post := p.Post.GetPostById(id)
	return mvc.View{
		Name: "post/view.html",
		Data: map[string]interface{}{
			"Title": post.Title,
			"IntroHeader": post.Title,
			"Post": post,
		},
	}
}

func (p *PostController) GetNew() mvc.Result {
	if p.IsLogged() {
		tags := p.Tags.GetAllTag()
		return mvc.View{
			Name: "post/new.html",
			Data: map[string]interface{}{
				"Title": "",
				"Tags": tags,
				"Authenticated": true,
				"RequireQuillPlugin": true,
				"RequireConsolePlugin": true,
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
	newPostMapTag := &models.PostMapTag{}
	tags := p.Ctx.PostValues("post-tag")
	newPost.Title = p.Ctx.PostValue("post-title")
	newPost.Content = p.Ctx.PostValue("post-content")
	newPost.CreateAt, newPost.UpdateAt, newPost.Uid = time.Now(), time.Now(), p.CurrentUserId()
	pid, err := models.CreatePost(newPost)
	p.Ctx.Writef("%s, %s", pid, err)
	if err != nil {
		p.Session.SetFlash(CreateFatalMsgKey, err)
	}
	for tag := range tags {
		newPostMapTag.Tid, err = strconv.ParseInt(tags[tag], 10, 64)
		newPostMapTag.Pid = pid
	}
}