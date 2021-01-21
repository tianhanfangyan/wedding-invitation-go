package models

import (
	"wedding-invitation-go/utils"
)

type Comment struct {
	UserId int64 `json:"user_id"` // 用户ID
}

type CommentResponse struct {
	UserId   int64  `json:"user_id"`   // 用户ID
	BlessMsg string `json:"bless_msg"` // 祝福文字
	NickName string `json:"nick_name"` // 微信别名
	WxFace   string `json:"wx_face"`   // 微信头像
}

// 添加新人祝福评论
func (c *Comment) AddBlessCommentByUserId(args NewComerCommentRequest) error {
	comment := Comments{
		UserId:   args.UserId,
		BlessMsg: args.BlessMsg,
		OpenId:   args.OpenId,
	}

	db := utils.GetDB().Model(new(Comments)).Create(&comment)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// 查询新人祝福评论
func (c *Comment) GetBlessCommentsByUserId(pageNumber, pageSize int, maps map[string]interface{}) ([]CommentResponse, error) {
	var results []CommentResponse
	var err error

	if pageNumber > 0 && pageSize > 0 {
		err = utils.GetDB().Table("comments").Select("user_id, nick_name, wx_face, bless_msg").Joins("left join wx_users on comments.open_id=wx_users.open_id").Where(maps).Offset((pageNumber - 1) * pageSize).Limit(pageSize).Scan(&results).Error
	} else {
		err = utils.GetDB().Table("comments").Select("user_id, nick_name, wx_face, bless_msg").Joins("left join wx_users on comments.open_id=wx_users.open_id").Where(maps).Scan(&results).Error
	}

	if err != nil {
		return results, err
	}

	return results, nil
}

// 获取行数
func (c *Comment) GetBlessCommentsTotal(maps interface{}) (int, error) {
	var count int

	db := utils.GetDB().Model(new(Comments)).Where(maps).Count(&count)
	if err := db.Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 筛选条件
func (c *Comment) GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if c.UserId > 0 {
		maps["user_id"] = c.UserId
	}

	return maps
}
