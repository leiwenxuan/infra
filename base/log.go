package base

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	// 定义日志格式
	formatter := &prefixed.TextFormatter{}
	//开启完整时间戳输出和时间戳格式
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	// 控制台高亮显示
	formatter.ForceFormatting = true
	formatter.ForceColors = true
	formatter.DisableColors = true
	//设置高亮显示的色彩样式
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	log.SetFormatter(formatter)
	//日志级别，通过环境变量来设置
	// 后期可以变更到配置中来设置
	level := os.Getenv("log.debug")
	log.Info("log level %s", level)
	if level == "true" {
		log.SetLevel(log.DebugLevel)
	}
	//log.SetLevel(log.DebugLevel)

	//log.Info("测试")
	//log.Debug("测试")
	//日志文件和滚动配置
	//github.com/lestrrat/go-file-rotatelogs

	logFileSettings()
	log.Info("日志初始化成功")
}

func logFileSettings() {

	//配置日志输出目录
	logPath, _ := filepath.Abs("./logs")
	log.Infof("log dir: %s", logPath)
	logFileName := "infra"
	//日志文件最大保存时间，24小时
	maxAge := time.Hour * 24
	//日志切割时间间隔,1小时一个
	rotationTime := time.Hour * 1
	_ = os.MkdirAll(logPath, os.ModePerm)

	baseLogPath := path.Join(logPath, logFileName)
	//设置滚动日志输出
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", err)
	}
	log.SetOutput(writer)

}
