package config

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
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
	err := os.MkdirAll(logDir, os.ModePerm) // 创建多级目录
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return
	}

	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: timeFormat}
	consoleWriter.FormatTimestamp = func(i interface{}) string {
		t := "<nil>"
		switch tt := i.(type) {
		case string:
			ts, err := time.Parse(zerolog.TimeFieldFormat, tt)
			if err != nil {
				t = tt
			} else {
				t = ts.Format(consoleWriter.TimeFormat)
			}
		case json.Number:
			i, err := tt.Int64()
			if err != nil {
				t = tt.String()
			} else {
				var sec, nsec int64 = i, 0
				switch zerolog.TimeFieldFormat {
				case zerolog.TimeFormatUnixMs:
					nsec = int64(time.Duration(i) * time.Millisecond)
					sec = 0
				case zerolog.TimeFormatUnixMicro:
					nsec = int64(time.Duration(i) * time.Microsecond)
					sec = 0
				}
				ts := time.Unix(sec, nsec)
				t = ts.Format(consoleWriter.TimeFormat)
			}
		}
		return fmt.Sprintf("%s", t)
	}
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

	level, _ := zerolog.ParseLevel(Viper.Log.Level) // 日志等级转换

	logContext := zerolog.New(multi).With().Timestamp()

	if Viper.AppDebug {
		logContext = logContext.Caller()
	}

	Log = logContext.Logger().Level(level)

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
