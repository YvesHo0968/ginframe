package common

import (
	"fmt"
	"ginFrame/config"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123"))
}

func TestUniqueId(t *testing.T) {
	fmt.Println(UniqueId())
}

func TestBase64Encoded(t *testing.T) {
	fmt.Println(Base64Encoded("hello"))
}

func TestBase64Decode(t *testing.T) {
	fmt.Println(Base64Decode("aGVsbG8="))
}

func TestMd5File(t *testing.T) {
	fmt.Println(Md5File("/Volumes/DATA/镜像/CentOS-7-x86_64-Minimal-2009.iso"))
}

func TestSha1File(t *testing.T) {
	fmt.Println(Sha1File("/Volumes/DATA/镜像/CentOS-7-x86_64-Minimal-2009.iso"))
}

func TestStrToLower(t *testing.T) {
	fmt.Println(StrToLower("Hello"))
}

func TestLcFirst(t *testing.T) {
	fmt.Println(UcWords("hello soed"))
}

func TestMdStrLen(t *testing.T) {
	fmt.Println(MdStrLen("hello中国"))
}

func TestStrToUpper(t *testing.T) {
	fmt.Println(StrToUpper("Hello"))
}

func TestRand(t *testing.T) {
	fmt.Println(Rand(1000, 9999))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(1.2))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(1.9))
}

func TestRound(t *testing.T) {
	fmt.Println(Round(1.5))
}

func TestDD(t *testing.T) {
	config.InitLog()

	log := config.Log

	log.WithFields(logrus.Fields{
		"status_code": "200",
	}).Info()
}

func TestContextual(t *testing.T) {
	config.Initll()

	config.Logger.Info().
		Str("website", "xx").
		Str("account", "account").
		Msg("开始登录...")

	logg := config.Logger.With().Caller().Str("foo", "bar").Logger()
	logg.Info().Msg("Hello wrold")
}

func TestContextualLogger(t *testing.T) {
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
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	//consoleWriter.FormatLevel = func(i interface{}) string {
	//	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	//}
	//consoleWriter.FormatMessage = func(i interface{}) string {
	//	return fmt.Sprintf("%s", i)
	//}
	//consoleWriter.FormatFieldName = func(i interface{}) string {
	//	return fmt.Sprintf("%s:", i)
	//}
	//consoleWriter.FormatFieldValue = func(i interface{}) string {
	//	return fmt.Sprintf("%s;", i)
	//}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)

	fmt.Println(multi)
	//log := zerolog.New(os.Stdout)
	log := zerolog.New(multi)
	log.Info().Str("content", "Hello world").Int("count", 3).Msg("TestContextualLogger")

	// 添加上下文 (文件名/行号/字符串)
	log = log.With().Timestamp().Caller().Str("foo", "bar").Logger()
	log.Info().Msg("Hello wrold")
}
