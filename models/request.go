package models

// 登录请求值
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 新人请求值
type NewComerInfoRequest struct {
	UserId       int64  `json:"user_id"`       // 用户ID
	HeName       string `json:"he_name"`       // 新郎姓名
	SheName      string `json:"she_name"`      // 新娘姓名
	WeddingDate  string `json:"wedding_date"`  // 婚礼日期
	WeddingLunar string `json:"wedding_lunar"` // 农历日期
	ShareMsg     string `json:"share_msg"`     // 分享文字
	HotelName    string `json:"hotel_name"`    // 酒店名称
	HotelAddress string `json:"hotel_address"` // 酒店地址
	HotelLat     int64  `json:"hotel_lat"`     // 酒店经度
	HotelLng     int64  `json:"hotel_lng"`     // 酒店纬度
}

// 新人祝福评论请求值
type NewComerCommentRequest struct {
	UserId   int64  `json:"user_id"`   // 用户ID
	BlessMsg string `json:"bless_msg"` // 祝福文字
	NickName string `json:"nick_name"` // 微信别名
	WxFace   string `json:"wx_face"`   // 微信头像
}
