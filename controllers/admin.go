package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

// 添加新人信息
func AddNewComerInfo(c *gin.Context) {
	var args models.NewComerInfoRequest

	err := c.BindJSON(&args)
	if err != nil {
		utils.GetLog().Sugar().Errorf("Parse invalid args filed: %s", err)
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	admin := new(models.Admin)
	err = admin.AddNewComerInfo(args)
	if err != nil {
		models.RetFailWithMessage("添加新人信息失败", c)
		return
	}

	models.RetSuccessWithMessage("添加新人信息成功", c)
}

// 删除新人信息
func DeleteNewComerInfo(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("user_id"))
	utils.GetLog().Sugar().Infof("userId: %d", userId)

	if userId <= 0 {
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	admin := new(models.Admin)
	err := admin.DeleteNewComerInfo(userId)
	if err != nil {
		models.RetFailWithMessage("删除新人信息失败", c)
		return
	}

	models.RetSuccessWithMessage("删除新人信息成功", c)
}

// 获取所有新人信息
func GetNewComerInfo(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("user_id"))
	pageNumber := utils.StringToInt(c.Query("page_number"))
	pageSize := utils.StringToInt(c.Query("page_size"))
	utils.GetLog().Sugar().Infof("userId: %d, pageNumber: %d, pageSize: %d", userId, pageNumber, pageSize)

	pageNumber, pageSize = utils.GetPage(pageNumber, pageSize)

	admin := models.Admin{UserId: userId}
	result, err := admin.GetNewComerInfo(pageNumber, pageSize, admin.GetMaps())
	if err != nil {
		models.RetFailWithMessage("获取所有新人信息失败", c)
		return
	}

	totalCount, err := admin.GetAllNewComerTotal(admin.GetMaps())
	if err != nil {
		models.RetFailWithMessage("查询失败", c)
		return
	}

	totalPage := totalCount / pageSize
	if totalCount > totalPage*pageSize {
		totalPage += 1
	}

	models.RetSuccessWithData(models.Page{PageNumber: int64(pageNumber), PageSize: int64(pageSize), TotalCount: int64(totalCount), TotalPage: int64(totalPage), Data: result}, c)
}

// 修改新人信息
func ModifyNewComerInfo(c *gin.Context) {
	var args models.NewComerInfoRequest

	err := c.BindJSON(&args)
	if err != nil {
		utils.GetLog().Sugar().Errorf("Parse invalid args filed: %s", err)
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	admin := new(models.Admin)
	err = admin.ModifyNewComerInfo(args)
	if err != nil {
		models.RetFailWithMessage("修改新人信息失败", c)
		return
	}

	models.RetSuccessWithMessage("修改新人信息成功", c)
}
