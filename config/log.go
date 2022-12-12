package config

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var Logger zerolog.Logger

func Initll() {
	timeFormat := "2006-01-02 15:04:05"
	zerolog.TimeFieldFormat = timeFormat

	// 创建log目录
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return
	}

	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s;", i)
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	Logger = zerolog.New(multi).With().Timestamp().Logger()
}

func InitLog() {
	//src, err := os.OpenFile("./logger.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	src, err := os.OpenFile("./logger.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
	}

	Log = logrus.New()

	// 添加文件名和方法信息
	Log.SetReportCaller(true)

	// 设置输出
	Log.Out = src

	// 设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		"./logger.log.%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName("./logger.log"),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)
}
