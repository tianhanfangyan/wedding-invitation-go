package models

// 用户验证表
type Auths struct {
	Id       int64  `gorm:"AUTO_INCREMENT;primary_key"` // ID
	Username string `gorm:"type:varchar(20)"`           // 用户名
	Password string `gorm:"type:varchar(20)"`           // 密码
}

// 新人信息表
type Users struct {
	Id           int64      `gorm:"AUTO_INCREMENT;primary_key"` // 用户ID
	HeName       string     `gorm:"type:varchar(10)"`           // 新郎姓名
	SheName      string     `gorm:"type:varchar(10)"`           // 新娘姓名
	WeddingDate  string     `gorm:"type:varchar(20)"`           // 婚礼日期
	WeddingLunar string     `gorm:"type:varchar(20)"`           // 农历日期
	ShareMsg     string     `gorm:"type:varchar(255)"`          // 分享文字
	Hotels       Hotels     `gorm:"foreignKey:UsersId"`         // has one
	Comments     []Comments `gorm:"foreignKey:UsersId"`         // 一对多
}

// 新人酒店表
type Hotels struct {
	Id           int64   `gorm:"AUTO_INCREMENT;primary_key"` // 自增ID
	HotelName    string  `gorm:"type:varchar(50)"`           // 酒店名称
	HotelAddress string  `gorm:"type:varchar(50)"`           // 酒店地址
	HotelLat     float64 `gorm:"type:not null"`              // 酒店经度
	HotelLng     float64 `gorm:"type:not null"`              // 酒店纬度
	UserId       int64   `gorm:"type:not null"`              // 外键
}

// 新人祝福评论表
type Comments struct {
	Id       int64  `gorm:"AUTO_INCREMENT;primary_key"` // 自增ID
	BlessMsg string `gorm:"type:text"`                  // 祝福文字
	NickName string `gorm:"type:text"`                  // 微信别名
	WxFace   string `gorm:"type:text"`                  // 微信头像
	UserId   int64  `gorm:"type:not null"`              // 外键
}
