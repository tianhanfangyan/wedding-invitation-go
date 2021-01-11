package models

// 登录请求值
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 新人请求值
type NewComerInfoRequest struct {
	UserId       int64   `json:"user_id"`                          // 用户ID
	HeName       string  `json:"he_name" binding:"required"`       // 新郎姓名
	SheName      string  `json:"she_name" binding:"required"`      // 新娘姓名
	WeddingDate  string  `json:"wedding_date" binding:"required"`  // 婚礼日期
	WeddingLunar string  `json:"wedding_lunar" binding:"required"` // 农历日期
	ShareMsg     string  `json:"share_msg" binding:"required"`     // 分享文字
	HotelName    string  `json:"hotel_name" binding:"required"`    // 酒店名称
	HotelAddress string  `json:"hotel_address" binding:"required"` // 酒店地址
	HotelLat     float64 `json:"hotel_lat" binding:"required"`     // 酒店经度
	HotelLng     float64 `json:"hotel_lng" binding:"required"`     // 酒店纬度
}

// 新人祝福评论请求值
type NewComerCommentRequest struct {
	UserId   int64  `json:"user_id"`   // 用户ID
	BlessMsg string `json:"bless_msg"` // 祝福文字
	NickName string `json:"nick_name"` // 微信别名
	WxFace   string `json:"wx_face"`   // 微信头像
}
