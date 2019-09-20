package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type BaseController struct {
	Ctx iris.Context
	Session *sessions.Session
}