package common

import "fmt"

// GetMapName 返回地图名
func GetMapName(id int32, arg ...interface{}) string {
	line := 0
	if len(arg) > 0 {
		line = arg[0].(int)
	}
	return fmt.Sprintf("normal_map_%d_%d", id, line)
}
