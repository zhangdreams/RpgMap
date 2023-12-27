package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapConfigs []MapConfig

func (m *MapConfig) UnmarshalJSON(data []byte) error {
	type Alias MapConfig
	aux := &struct {
		UnWalkStr string `json:"UnWalkStr"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Split UnWalkStr into substrings
	coords := strings.Split(strings.Trim(aux.UnWalkStr, "()"), "),(")

	// Parse each substring into a ConfigPos struct
	m.UnWalk = make([]ConfigPos, len(coords))
	for i, str := range coords {
		parts := strings.Split(str, ",")
		if len(parts) != 2 {
			return fmt.Errorf("invalid coordinate format: %s", str)
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("failed to parse X coordinate: %v", err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("failed to parse Y coordinate: %v", err)
		}
		m.UnWalk[i] = ConfigPos{X: int32(x), Y: int32(y)}
	}

	return nil
}

func ReadMap() {
	// 读取配置文件
	filePath := "./config/json/maps.json"
	jsonConfig, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading map config file:", err)
		return
	}

	// 解析 JSON 配置
	err = json.Unmarshal(jsonConfig, &mapConfigs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//showAllMaps()
}

// 输出解析结果
func showAllMaps() {
	fmt.Println("show map config")
	for _, config := range mapConfigs {
		fmt.Println("MapID:", config.MapID)
		fmt.Println("Name:", config.Name)
		fmt.Println("Width:", config.Width)
		fmt.Println("Height:", config.Height)
		fmt.Println("BornX:", config.BornX)
		fmt.Println("BornY:", config.BornY)
		fmt.Println("MaxNum:", config.MaxNum)
		fmt.Println("UnWalkStr:", config.UnWalkStr)
		fmt.Println("unWalk pos :")
		for _, pos := range config.UnWalk {
			fmt.Println("(", pos.X, pos.Y, ")")
		}
		fmt.Println("----------")
	}
}

// GetMapIDs 返回所有地图ID
func GetMapIDs() (IDs []int32) {
	for _, config := range mapConfigs {
		IDs = append(IDs, config.MapID)
	}
	return
}

// GetMapConfig 返回一个指定的地图ID
// mapID 地图ID
func GetMapConfig(mapID int32) (config MapConfig, err error) {
	for _, conf := range mapConfigs {
		if mapID == conf.MapID {
			config = conf
			return
		}
	}
	err = errors.New(fmt.Sprintf("not found map %d", mapID))
	return
}

func (pos ConfigPos) InSlice(posList []ConfigPos) bool {
	for _, p := range posList {
		if p.X == pos.X && p.Y == pos.Y {
			return true
		}
	}
	return false
}
