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
	OpenId   string `json:"open_id"`   // 微信openid
	BlessMsg string `json:"bless_msg"` // 祝福文字
}

// 微信用户信息请求值
type WxUserInfoRequest struct {
	OpenId   string `json:"open_id"`   // 微信openid
	UnionId  string `json:"union_id"`  // 微信开放平台union_id
	NickName string `json:"nick_name"` // 微信别名
	WxFace   string `json:"wx_face"`   // 微信头像
	Gender   int    `json:"gender"`    // 性别
	Province string `json:"province"`  // 省份
	City     string `json:"city"`      // 城市
	Mobile   string `json:"mobile"`    // 手机号码
}

// 微信用户加密数据请求值
type WxEncryptedDataRequest struct {
	SessionKey    string
	EncryptedData string
	Iv            string
}
