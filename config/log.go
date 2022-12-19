package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func InitLog() {
	timeFormat := "2006-01-02 15:04:05"
	zerolog.TimeFieldFormat = timeFormat

	// 设置字段别名
	//zerolog.TimestampFieldName = "t"
	//zerolog.LevelFieldName = "l"
	//zerolog.MessageFieldName = "m"

	// 创建log目录
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return
	}

	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.Stamp}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s ---- ", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s;", i)
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)

	level, _ := zerolog.ParseLevel(viper.GetString("log.level"))

	if viper.GetBool("app_debug") {
		Log = zerolog.New(multi).With().Timestamp().Caller().Logger().Level(level)
	} else {
		Log = zerolog.New(multi).With().Timestamp().Logger().Level(level)
	}

	// 日志采样 每隔多少条输出一次
	//Log = Log.Sample(&zerolog.BasicSampler{N: 10})

	// 日志采样
	//Log = Log.Sample(&zerolog.LevelSampler{
	//	DebugSampler: &zerolog.BurstSampler{
	//		Burst:       5,
	//		Period:      time.Second,
	//		NextSampler: &zerolog.BasicSampler{N: 100},
	//	},
	//})

	//for i := 0; i < 20; i++ {
	//	Log.Info().Msg("will be logged every 10 message")
	//}
}
