package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
)

var defaultConfig *Config

type Config struct {
	conf        config.Configer
	env         string
	absConfPath string
	filePath    string
}

func NewConfig(confPath string, confType string) (*Config, error) {
	c := new(Config)
	var err error

	c.absConfPath, err = filepath.Abs(confPath)
	if err != nil {
		return nil, err
	}

	c.env = fmt.Sprintf("%s", GetEnv())
	c.filePath = c.absConfPath + "/" + c.env + ".conf"

	c.conf, err = config.NewConfig(confType, c.filePath)
	if err != nil {
		return nil, err
	}

	c.merge()
	return c, nil
}

func (c *Config) merge() (err error) {
	baseDir := filepath.Dir(c.filePath)

	envConfFile, err := os.Open(baseDir + "/" + c.env + ".conf")
	if err != nil {
		log.Println("配置文件missing")
		os.Exit(2)
	}
	defer envConfFile.Close()

	appConfFile, err := os.Create(baseDir + "/app.conf")
	if err != nil {
		log.Println("无法创建app.conf")
		os.Exit(2)
	}
	defer appConfFile.Close()

	io.Copy(appConfFile, envConfFile)
	return nil
}

func (c *Config) GetConf() config.Configer {
	return c.conf
}

func InitConfig(confPath string) (err error) {
	defaultConfig, err = NewConfig(confPath, "ini")
	if err != nil {
		return err
	}

	return
}

func GetConf() config.Configer {
	return defaultConfig.GetConf()
}
