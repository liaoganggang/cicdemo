package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志配置信息
type LogConfig struct {
	Filename   string //日志文件路径
	Maxsize    int    //分割文件大小
	Maxbackups int    //备份多少次
	Compress   bool   //是否压缩
}

func init() {
	filename := "logs/deploy.log"
	logg := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,
		MaxBackups: 7,
		Compress:   true,
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(logg)
	logrus.SetOutput(os.Stdout)

}

var Debug = logrus.Debug
var Error = logrus.Error
