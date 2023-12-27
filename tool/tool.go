package tool

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToIntArray(str string) []int32 {
	if str == "" {
		return make([]int32, 0)
	}
	parts := strings.Split(str, ",")
	result := make([]int32, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return nil
		}
		result[i] = int32(num)
	}
	return result
}

// IF 一个简单的判断
func IF(t bool, p1, p2 interface{}) interface{} {
	if t {
		return p1
	} else {
		return p2
	}
}
