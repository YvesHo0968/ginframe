package common

import (
	"errors"
	"fmt"
	"ginFrame/config"
	"github.com/rs/zerolog"
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

func TestSleep(t *testing.T) {
	Sleep(1)
}

func TestUsleep(t *testing.T) {
	Usleep(1000000)
}

func TestGetHostName(t *testing.T) {
	fmt.Println(GetHostName())
}

func TestGetOS(t *testing.T) {
	fmt.Println(GetOS())
}

func TestGetArch(t *testing.T) {
	fmt.Println(GetArch())
}

func TestGetArchBit(t *testing.T) {
	fmt.Println(GetArchBit())
}

func TestGetCpuCores(t *testing.T) {
	fmt.Println(GetCpuCores())
}

func TestLog(t *testing.T) {
	config.InitLog()

	log := config.Log

	log.Info().Str("foo", "bar").Msg("Hello World")
	log.Error().Str("foo", "bar").Msg("Hello World")

	log.Info().
		Str("foo", "bar").
		Dict("dict", zerolog.Dict().
			Str("bar", "baz").
			Int("n", 1),
		).Msg("hello world")

	err := errors.New("A repo man spends his life getting into tense situations")

	fmt.Println(err)
	service := "myservice"

	log.Fatal().
		Err(err).
		Str("service", service).
		Msgf("Cannot start %s", service)
}

func TestStructToJson(t *testing.T) {
	//DataMap := map[string]int{"a": 1, "b": 2, "c": 3}
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		Sex  string `json:"sex,omitempty"`
	}
	sList := []S{{
		Name: "小红",
		Age:  18,
		Sex:  "女",
	}, {
		Name: "小李",
		Age:  20,
		Sex:  "男",
	}, {
		Name: "夏龙",
		Age:  18,
		Sex:  "女",
	}}
	fmt.Println(JsonEncode(sList))
}

func TestJsonToStruct(t *testing.T) {
	str := `[{"name":"小红","age":18,"sex":"女"},{"name":"小李","age":20,"sex":"男"},{"name":"夏龙","age":18,"sex":"女"}]`

	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		Sex  string `json:"sex,omitempty"`
	}

	var data []S
	JsonDecode(str, &data)
	fmt.Println(data)
}

func TestSerialize(t *testing.T) {
	data := map[string]interface{}{"php": "世界上最好的语言"}
	
	//data := []int{1, 3}

	fmt.Println(Serialize(data))
}

func TestUnSerialize(t *testing.T) {
	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`
	out, _ := UnSerialize(str)
	m := out.(map[string]interface{})
	fmt.Println(m["php"])
}
