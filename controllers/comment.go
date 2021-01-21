package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

// 添加新人对应祝福评论
func AddBlessCommentsByUserId(c *gin.Context) {
	var args models.NewComerCommentRequest

	err := c.BindJSON(&args)
	if err != nil {
		models.RetFailWithMessage("请求参数不规范", c)
		return
	}

	comment := new(models.Comment)
	err = comment.AddBlessCommentByUserId(args)
	if err != nil {
		models.RetFailWithMessage("添加祝福评论失败", c)
		return
	}

	models.RetSuccessWithMessage("添加祝福评论成功", c)
}

// 获取新人对应所有祝福评论
func GetBlessCommentsByUserId(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("user_id"))
	pageNumber := utils.StringToInt(c.Query("page_number"))
	pageSize := utils.StringToInt(c.Query("page_size"))
	utils.GetLog().Sugar().Infof("userId: %d, pageNumber: %d, pageSize: %d", userId, pageNumber, pageSize)

	if userId <= 0 {
		models.RetFailWithMessage("登录请求参数不规范", c)
		return
	}

	pageNumber, pageSize = utils.GetPage(pageNumber, pageSize)

	comment := models.Comment{UserId: userId}
	result, err := comment.GetBlessCommentsByUserId(pageNumber, pageSize, comment.GetMaps())
	if err != nil {
		models.RetFailWithMessage("查询失败", c)
		return
	}

	totalCount, err := comment.GetBlessCommentsTotal(comment.GetMaps())
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
