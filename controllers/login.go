package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

func Login(c *gin.Context) {
	var args models.LoginRequest

	err := c.BindJSON(&args)
	if err != nil {
		utils.GetLog().Sugar().Errorf("Parse invalid args filed: %s", err)
		models.RetFailWithMessage("登录请求参数不规范", c)
		return
	}

	login := new(models.Login)
	var result models.LoginResponse

	result, err = login.LoginAuth(args.Username, args.Password)
	if err != nil {
		models.RetFailWithMessage("登录请求失败", c)
		return
	}

	models.RetSuccessWithDetail(result, "登录请求成功", c)
}
