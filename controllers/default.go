package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"html/template"
	"strconv"
	"www.wooyanger.com/models"
)

// 定义常量
const UserIdKey  = "UserId"

// 定义主Controller结构
type Controllers struct {
	Ctx			iris.Context
	Post		*models.Post
	Config		*models.Config
	Session		*sessions.Session
}

// 获取当前用户ID
func (c *Controllers) CurrentUserId() int64 {
	return c.Session.GetInt64Default(UserIdKey, 0)
}

// 是否已经登录
func (c *Controllers) IsLogged() bool {
	return c.CurrentUserId() > 0
}

// 登录
func (c *Controllers) LoginUser(uid int64) {
	c.Session.Set(UserIdKey, uid)
}

// 字符串数字转整型
func (c *Controllers) StringToInt64(str string) int64 {
	d, _ := strconv.ParseInt(str, 10, 64)
	return d
}

// 不转义HTML标签
func (c *Controllers) HTML(str string) template.HTML {
	return template.HTML(str)
}