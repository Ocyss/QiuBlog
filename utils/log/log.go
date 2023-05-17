package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"qiublog/utils"
	"qiublog/utils/middleware"
)

func InitLogger() {

	formatter := log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableQuote:              true,
	}
	log.SetFormatter(&formatter)
	logConf := utils.Config.Log
	if utils.Dev {
		log.SetLevel(log.DebugLevel)
	} else {
		level, err := log.ParseLevel(logConf.Level)
		if err != nil {
			panic(fmt.Sprintf("错误的日志等级,可用的[panic,fatal,error,warn,info,debug,trace],%v", err))
		}
		log.SetLevel(level)
	}
	if utils.Debug {
		log.SetReportCaller(true)
	}
	if logConf.Enable {
		var w io.Writer = &lumberjack.Logger{
			Filename:   logConf.Filename,
			MaxSize:    logConf.MaxSize,
			MaxBackups: logConf.MaxBackups,
			MaxAge:     logConf.MaxAge,
			Compress:   logConf.Compress,
		}
		if utils.Dev || utils.Debug {
			w = io.MultiWriter(os.Stdout, w)
		}
		log.SetOutput(w)
	}
	gin.DebugPrintRouteFunc = middleware.LoggerDebug
	log.Debug("init logrus success!")
}
