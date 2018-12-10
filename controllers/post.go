package controllers

import (
	"fmt"
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

func getTagId(str []string) []int64 {
	var tags []int64
	for s := range str {
		d, _ := strconv.ParseInt(str[s], 10, 64)
		tags = append(tags, d)
	}
	return tags
}

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
	if err != nil {
		p.Session.SetFlash(CreateFatalMsgKey, err)
	}
	for _, tag := range getTagId(tags) {
		newPostMapTag.Tid = tag
		newPostMapTag.Pid = pid
		if err := models.CreatePostMapTag(newPostMapTag); err != nil {
			p.Session.SetFlash(CreateFatalMsgKey, err)
		}
	}
	p.Ctx.Redirect(fmt.Sprintf("/posts/%d", pid))
}