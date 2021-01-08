package main

import (
	"flag"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"wedding-invitation-go/routers"
	"wedding-invitation-go/utils"
)

var (
	env = flag.String("env", "dev", "dev, test, prod")
)

func main() {
	flag.Parse()

	// 环境初始化
	if err := utils.SetEnv(*env); err != nil {
		log.Println(err)
		return
	}
	log.Println("env init")

	// 配置初始化
	if err := utils.InitConfig("./conf/"); err != nil {
		log.Println(err)
		return
	}
	log.Println("config init")

	// jwtSecret初始化
	utils.InitJwtSecret()
	log.Println("jwtSecret init")

	// 日志初始化
	if err := utils.InitLogger(); err != nil {
		log.Println(err)
		return
	}
	log.Println("logger init")

	// 数据库初始化
	if err := utils.InitDB(); err != nil {
		log.Println(err)
		return
	}
	log.Println("db init")

	// 设置gin运行模式
	gin.SetMode(utils.GetConf().String("app::runmode"))

	// 初始化gin路由
	router := routers.InitRouters()

	// 添加公共中间件
	router.Use(gin.Recovery())
	router.Use(ginzap.Ginzap(utils.GetLog(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(utils.GetLog(), true))

	// ping
	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
