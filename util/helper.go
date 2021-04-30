package util

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func IntegerToString(i int64) string {
	s := strconv.Itoa(int(i))
	return s
}

func StringToInteger(txt string) int {
	i, _ := strconv.Atoi(txt)
	return int(i)
}
