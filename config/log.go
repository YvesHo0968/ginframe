package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func InitLog() {
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
	Log = zerolog.New(multi).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)
}
