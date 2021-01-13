package models

import (
	"wedding-invitation-go/utils"
)

type Admin struct {
	UserId int64 `json:"user_id"` // 用户ID
}

type NewComerResponse struct {
	Id           int64   `json:"user_id"`       // 用户ID
	HeName       string  `json:"he_name"`       // 新郎姓名
	SheName      string  `json:"she_name"`      // 新娘姓名
	WeddingDate  string  `json:"wedding_date"`  // 婚礼日期
	WeddingLunar string  `json:"wedding_lunar"` // 农历日期
	ShareMsg     string  `json:"share_msg"`     // 分享文字
	HotelName    string  `json:"hotel_name"`    // 酒店名称
	HotelAddress string  `json:"hotel_address"` // 酒店地址
	HotelLat     float64 `json:"hotel_lat"`     // 酒店经度
	HotelLng     float64 `json:"hotel_lng"`     // 酒店纬度
}

// 添加
func (a *Admin) AddNewComerInfo(args NewComerInfoRequest) error {
	user := Users{
		HeName:       args.HeName,
		SheName:      args.SheName,
		WeddingDate:  args.WeddingDate,
		WeddingLunar: args.WeddingLunar,
		ShareMsg:     args.ShareMsg,
	}

	db := utils.GetDB().Model(new(Users)).Create(&user) // 通过数据的指针来创建
	if err := db.Error; err != nil {
		return err
	}

	hotel := Hotels{
		HotelAddress: args.HotelAddress,
		HotelName:    args.HotelName,
		HotelLat:     args.HotelLat,
		HotelLng:     args.HotelLng,
		UserId:       user.Id,
	}

	db = utils.GetDB().Model(new(Hotels)).Create(&hotel) // 通过数据的指针来创建
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// 删除
func (a *Admin) DeleteNewComerInfo(userId int64) error {
	// 删除新人信息
	db := utils.GetDB().Where("id=?", userId).Delete(&Users{})
	if err := db.Error; err != nil {
		return err
	}

	// 删除酒店信息
	db = utils.GetDB().Where("user_id=?", userId).Delete(&Hotels{})
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// 获取全部
func (a *Admin) GetNewComerInfo(pageNumber, pageSize int, maps map[string]interface{}) ([]NewComerResponse, error) {
	var results []NewComerResponse
	var err error

	if pageNumber > 0 && pageSize > 0 {
		err = utils.GetDB().Table("users").Select("users.id, he_name, she_name, wedding_date, wedding_lunar, share_msg, hotel_name, hotel_address, hotel_lat, hotel_lng").Joins("join hotels on users.id=hotels.user_id").Where(maps).Offset((pageNumber - 1) * pageSize).Limit(pageSize).Scan(&results).Error
	} else {
		err = utils.GetDB().Table("users").Select("users.id, he_name, she_name, wedding_date, wedding_lunar, share_msg, hotel_name, hotel_address, hotel_lat, hotel_lng").Joins("join hotels on users.id=hotels.user_id").Where(maps).Scan(&results).Error
	}

	if err != nil {
		return results, err
	}

	return results, nil
}

// 获取行数
func (a *Admin) GetAllNewComerTotal(maps interface{}) (int, error) {
	var count int

	db := utils.GetDB().Model(new(Users)).Where(maps).Count(&count)
	if err := db.Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 修改
func (a *Admin) ModifyNewComerInfo(args NewComerInfoRequest) error {
	user := Users{
		HeName:       args.HeName,
		SheName:      args.SheName,
		WeddingDate:  args.WeddingDate,
		WeddingLunar: args.WeddingLunar,
		ShareMsg:     args.ShareMsg,
	}

	db := utils.GetDB().Model(new(Users)).Where("id=?", args.UserId).Updates(user)
	if err := db.Error; err != nil {
		return err
	}

	hotel := Hotels{
		HotelAddress: args.HotelAddress,
		HotelName:    args.HotelName,
		HotelLat:     args.HotelLat,
		HotelLng:     args.HotelLng,
		UserId:       user.Id,
	}

	db = utils.GetDB().Model(new(Hotels)).Where("user_id=?", args.UserId).Updates(hotel)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

// 筛选条件
func (a *Admin) GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if a.UserId > 0 {
		maps["id"] = a.UserId
	}

	return maps
}
