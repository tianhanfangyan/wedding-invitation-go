package utils

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var globalMiniProgram *miniprogram.MiniProgram

func InitWechatMiniProgram() {
	appId := GetConf().String("mini::app_id")
	appSecret := GetConf().String("mini::app_secret")

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     memory,
	}

	miniProgram := wc.GetMiniProgram(cfg)

	globalMiniProgram = miniProgram
}

func GetMiniProgram() *miniprogram.MiniProgram {
	return globalMiniProgram
}
