package common

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Success(c *gin.Context, data interface{}) {
	type SuccessData struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	d := SuccessData{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}

	c.JSON(http.StatusOK, d)
}

// Uuid 获取uuid
func Uuid() string {
	return uuid.New().String()
}

// Path 获取进程工作目录
func Path() string {
	path, _ := os.Getwd()

	return path
}

// FilePath 获取运行的目录
func FilePath() string {
	_, filename, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(filename))

	return root
}

// Time 获取系统时间戳
func Time() int64 {
	return time.Now().Unix()
}

// Date 时间戳转日期 Date("Y-m-d H:i:s")
func Date(format string, times ...int64) string {
	// DateFormat pattern rules.
	var datePatterns = []string{
		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
		"y", "06", // A two digit representation of a year   Examples: 99 or 03

		// month
		"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
		"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
		"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
		"F", "January", // A full textual representation of a month, such as January or March   January through December

		// day
		"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
		"j", "2", // Day of the month without leading zeros 1 to 31

		// week
		"D", "Mon", // A textual representation of a day, three letters Mon through Sun
		"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

		// time
		"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
		"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
		"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
		"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

		"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
		"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

		"i", "04", // Minutes with leading zeros    00 to 59
		"s", "05", // Seconds, with leading zeros   00 through 59

		// time zone
		"T", "MST",
		"P", "-07:00",
		"O", "-0700",

		// RFC 2822
		"r", time.RFC1123Z,
	}

	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)

	var t int64
	if len(times) > 0 {
		t = times[0]
	} else {
		t = Time()
	}

	return time.Unix(t, 0).Format(format)
}

// Datetotime 日期转时间戳 Datetotime("2006-01-02 15:04:05", "2022-01-01 11:00:00")
func Datetotime(format, date string) (int64, error) {
	t, err := time.ParseInLocation(format, date, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func Strtotime(str string) int64 {

	uintToSeconds := map[string]int64{"minute": 60, "hour": 3600, "day": 86400, "week": 604800, "year": ((365 * 86400) + 86400)}

	accumulator := time.Now().Unix()

	var delta int64
	plus := true
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "in ") {
		str = strings.Replace(str, "in ", "", 1)
	}

	if strings.Index(str, " ago") > 0 {
		str = strings.Replace(str, " ago", "", 1)
		plus = false
	}

	if strings.Index(str, "+") >= 0 {
		str = strings.Replace(str, "+", "", 1)
	}

	if strings.Index(str, "-") >= 0 {
		str = strings.Replace(str, "-", "", 1)
		plus = false
	}

	noteValMap := make(map[string]int64, 10)

	re := regexp.MustCompile(`\d+\s+(minute|hour|day|week|year)`)

	parts := re.FindAllStringSubmatch(str, -1)

	for i, _ := range parts {
		strArray := strings.Split(parts[i][0], " ")
		v, _ := strconv.Atoi(strArray[0])
		noteValMap[parts[i][1]] = int64(v)
	}

	delta = 0
	for k, v := range noteValMap {

		delta += uintToSeconds[k] * v
	}

	if plus {
		accumulator += delta
	} else {
		accumulator -= delta
	}

	return accumulator
}

// Md5 生成32位md5字串 Md5("123")
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 生成sha1字串 sha1("123")
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// UniqueId 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(Base64Encoded(string(b)))
}

// Base64Encoded Base64加密 Base64Encoded("hello")
func Base64Encoded(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// Base64Decode Base64解密 Base64Decode("aGVsbG8=")
func Base64Decode(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)

	return string(decode)
}
