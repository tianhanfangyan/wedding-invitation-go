package models

import "wedding-invitation-go/utils"

type Hotel struct {
}

type HotelInfoResponse struct {
	HotelName    string  `json:"hotel_name"`    // 酒店名称
	HotelAddress string  `json:"hotel_address"` // 酒店地址
	HotelLat     float64 `json:"hotel_lat"`     // 酒店经度
	HotelLng     float64 `json:"hotel_lng"`     // 酒店纬度
}

// 查询新人酒店信息
func (h *Hotel) GetHotelInfoByUserId(userId int64) (HotelInfoResponse, error) {
	var result HotelInfoResponse

	db := utils.GetDB().Model(new(Hotels)).Where("id=?", userId).Scan(&result)
	if err := db.Error; err != nil {
		return result, err
	}

	return result, nil
}
