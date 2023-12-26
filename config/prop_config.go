package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"rpgMap/global"
)

var propConfigs []global.Prop

func ReadProps() {
	filePath := "./config/json/props.json"
	jsonConfig, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading map config file:", err)
		return
	}

	err = json.Unmarshal(jsonConfig, &propConfigs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//showAllProps()
}

func showAllProps() {
	fmt.Println("show prop config")
	for _, prop := range propConfigs {
		fmt.Println("Index:", prop.Index)
		fmt.Println("MaxHp:", prop.MaxHp)
		fmt.Println("Attack:", prop.Attack)
		fmt.Println("Defense:", prop.Defense)
		fmt.Println("Speed:", prop.Speed)
		fmt.Println("----------")
	}
}

// GetPropConfig 返回指定index的配置
func GetPropConfig(index int32) (config global.Prop, err error) {
	for _, conf := range propConfigs {
		if index == conf.Index {
			config = conf
			return
		}
	}
	err = errors.New(fmt.Sprintf("not found prop %d", index))
	return
}
