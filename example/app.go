package main

import (
	"infra/v1"
	_ "infra/v1/base"

	"github.com/sirupsen/logrus"
	"github.com/tietang/props/ini"
)

func main() {
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource("config.ini")
	app := infra.New(conf)
	app.Start()
	serverPort := conf.GetDefault("server.port", "18080")
	logrus.Info(serverPort)

}
