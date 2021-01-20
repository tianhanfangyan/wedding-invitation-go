package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

func GetNewComerIndexImageByUserID(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("user_id"))
	utils.GetLog().Sugar().Infof("userId: %d", userId)

	if userId <= 0 {
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	image := new(models.Image)

	results := image.GetIndexImagesByUserId(userId)

	models.RetSuccessWithDetail(results, "获取新人图片成功", c)
}
