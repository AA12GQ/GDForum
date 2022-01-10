// Package logger 处理日志相关逻辑
package logger

import (
	"GDForum/app/app"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)
//全局Logger对象
var Logger *zap.Logger

func InitLogger(filename string,maxSize,maxBackup,maxAge int,compress bool,logType string,level string){
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level));err != nil{
		fmt.Println("日志初始化错误，日志级别设置有误，请修改config/log.go文件中的log.level配置项")
	}
	//初始化core
	core := zapcore.NewCore(getEncoder(),writeSyncer,logLevel)
	//初始化logger
	Logger = zap.New(core,
		zap.AddCaller(),						//调用文件和符号，内部使用runtime.Caller
		zap.AddCallerSkip(1),				//封装了一层，调用文件去除一层(runtime.Caller)
		zap.AddStacktrace(zap.ErrorLevel),		//Error时才会显示stacktrace
	)
	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)

}
//getEncoder设置日志存储格式
func getEncoder()zapcore.Encoder{
	//日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}
	//本地环境配置
	if app.IsLocal(){
		//终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		//本地设置内置的Console解码器（支持stacktrace换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	//线上环境使用JSON解码器
	return zapcore.NewJSONEncoder(encoderConfig)
}
// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
// getLogWriter 日志记录介质。GDForum 中使用了两种介质，os.Stdout 和文件
func getLogWriter(filename string,maxSize,maxBackup,maxAge int,compress bool,logType string)zapcore.WriteSyncer{
	//如果配置了按照日期记录日志文件
	if logType == "daily"{
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename,"logs.log",logname)
	}
	//滚动日志，详情见config/log.go
	// 滚动日志，详见 config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	//配置输出介质
	if app.IsLocal(){
		//本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(lumberJackLogger))
	}else{
		//生产环境只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}