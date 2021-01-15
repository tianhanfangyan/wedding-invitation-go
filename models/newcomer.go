package models

import (
	"wedding-invitation-go/utils"
)

type NewComer struct {
}

type NewComerInfoResponse struct {
	HeName       string `json:"he_name"`       // 新郎姓名
	SheName      string `json:"she_name"`      // 新娘姓名
	WeddingDate  string `json:"wedding_date"`  // 婚礼日期
	WeddingLunar string `json:"wedding_lunar"` // 农历日期
	ShareMsg     string `json:"share_msg"`     // 分享文字
}

// 查询新人的信息
func (n *NewComer) GetNewComerInfoByUserId(userId int64) (NewComerInfoResponse, error) {
	var result NewComerInfoResponse

	db := utils.GetDB().Model(new(Users)).Where("id=?", userId).Scan(&result)
	if err := db.Error; err != nil {
		utils.GetLog().Sugar().Errorf("GetNewComerInfoByUserId fialed: %s", err)
		return result, err
	}

	return result, nil
}
