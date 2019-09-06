package utils

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger

func InitLogger(mod string) {
	hook := lumberjack.Logger{
		Filename: "logs/app.log",
		// 每个日志文件保存的最大尺寸 单位：M
		MaxSize: 100,
		// 日志文件最多保存多少个备份
		MaxBackups: 3,
		// 文件最多保存多少天
		MaxAge: 1,
		// 是否压缩
		Compress: true,
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core

	if mod == gin.ReleaseMode {
		core = zapcore.NewCore(encoder, zapcore.AddSync(&hook), zap.InfoLevel)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	}

	zapLogger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCaller())
}

func GetLogger() *zap.Logger {
	return zapLogger
}