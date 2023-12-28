package tool

import (
	"fmt"
	"math"
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

// Ceil 向上取整
func Ceil(v float32) int32 {
	return int32(math.Ceil(float64(v)))
}

// Floor 向下取整
func Floor(v float32) int32 {
	return int32(math.Floor(float64(v)))
}

// Round 四舍五入
func Round(v float32) int32 {
	return int32(math.Round(float64(v)))
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
