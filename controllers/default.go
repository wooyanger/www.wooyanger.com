package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"www.wooyanger.com/models"
)

const UserIdKey  = "UserId"

type Controllers struct {
	Ctx			iris.Context
	Post		*models.Post
	Session		*sessions.Session
}

func (c *Controllers) CurrentUserId() int64 {
	return c.Session.GetInt64Default(UserIdKey, 0)
}

func (c *Controllers) IsLogged() bool {
	return c.CurrentUserId() > 0
}