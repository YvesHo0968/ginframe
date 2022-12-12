package common

import (
	"fmt"
	"ginFrame/config"
	"testing"
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

	log.Info().Str("foo", "bar").Msg("Hello World")
	log.Error().Str("foo", "bar").Msg("Hello World")
}
