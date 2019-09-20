package controller

import (
	"github.com/kataras/iris/mvc"
	"jhgocms/service"
)

type AdminUserController struct {
	BaseController
	Service service.AdminUserService
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
	用户登录
 */
func (auc *AdminUserController) PostLogin() mvc.Result {
	return mvc.Response{
		Object:map[string]interface{}{
			"code":20000,
			"data":"zhangsan",
		},
	}
}

/**
	用户登录
 */
func (auc *AdminUserController) PostLogout() mvc.Result {
	return mvc.Response{
		Object:map[string]interface{}{
			"code":20000,
			"data":"success",
		},
	}
}

func (auc *AdminUserController) GetInfo() mvc.Result {
	return mvc.Response{
		Object:map[string]interface{}{
			"code":20000,
			"data":"zhangsan",
		},
	}
}