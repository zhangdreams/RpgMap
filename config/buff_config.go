package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var buffConfigs []BuffConfig

func ReadBuffs() {
	filePath := "./config/json/buffs.json"
	jsonConfig, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading map config file:", err)
		return
	}

	err = json.Unmarshal(jsonConfig, &buffConfigs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//showAllBuffs()
}

func showAllBuffs() {
	fmt.Println("show buff config")
	for _, buff := range buffConfigs {
		fmt.Println("BuffID:", buff.BuffID)
		fmt.Println("Name:", buff.Name)
		fmt.Println("Type:", buff.Type)
		fmt.Println("DurTime:", buff.DurTime)
		fmt.Println("Level:", buff.Level)
		fmt.Println("AddType:", buff.AddType)
		fmt.Println("EffectType:", buff.EffectType)
		fmt.Println("Value:", buff.Value)
		fmt.Println("Func:", buff.Func)
		fmt.Println("----------")
	}
}

// GetBuffConfig 返回指定buffID的配置
func GetBuffConfig(buffID int32) (config BuffConfig, err error) {
	for _, conf := range buffConfigs {
		if buffID == conf.BuffID {
			config = conf
			return
		}
	}
	err = errors.New(fmt.Sprintf("not found buff %d", buffID))
	return
}
