package common

import (
	"fmt"
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
