package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

// 获取新人对应酒店的位置信息
func GetHotelLocationByUserId(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("user_id"))
	utils.GetLog().Sugar().Infof("userId: %d", userId)

	if userId <= 0 {
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	hotel := new(models.Hotel)
	result, err := hotel.GetHotelInfoByUserId(userId)
	if err != nil {
		models.RetFailWithMessage("查询酒店信息失败", c)
		return
	}

	models.RetSuccessWithDetail(result, "查询酒店信息成功", c)
}
