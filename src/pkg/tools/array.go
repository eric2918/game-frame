package tools

import (
	"strings"

	"github.com/spf13/cast"
)

func StringToInt64(s string) (res []int64) {
	slice := strings.Split(s, ",")
	for _, str := range slice {
		res = append(res, cast.ToInt64(str))
	}
	return
}

func StringToSlice(s string) (res []string) {
	slice := strings.Split(s, ",")
	for _, str := range slice {
		res = append(res, str)
	}
	return
}
