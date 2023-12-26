package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"rpgMap/tool"
)

var monsterConfigs []MonsterConfig

func (b *MonsterConfig) UnmarshalJSON(data []byte) error {
	type Alias MonsterConfig
	aux := &struct {
		Skills string `json:"Skill"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert Skill string to []int
	b.Skills = tool.ConvertToIntArray(aux.Skills)
	return nil
}

func ReadMonster() {
	filePath := "./config/json/monsters.json"
	jsonConfig, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading map config file:", err)
		return
	}

	err = json.Unmarshal(jsonConfig, &monsterConfigs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	//showAllMonsters()
}

func showAllMonsters() {
	fmt.Println("show monster config")
	for _, monster := range monsterConfigs {
		fmt.Println("ID:", monster.ID)
		fmt.Println("Name:", monster.Name)
		fmt.Println("Level:", monster.Level)
		fmt.Println("PropID:", monster.PropID)
		fmt.Println("PatrolDistance:", monster.PatrolDistance)
		fmt.Println("PursueDistance:", monster.PursueDistance)
		fmt.Println("AttackDistance:", monster.AttackDistance)
		fmt.Println("RebornTime:", monster.RebornTime)
		fmt.Println("Skills:", monster.Skills)
		fmt.Println("----------")
	}
}

// GetMonsterConfig 返回指定monsterID的配置
// monsterID 怪物ID
func GetMonsterConfig(monsterID int32) (config MonsterConfig, err error) {
	for _, conf := range monsterConfigs {
		if monsterID == conf.ID {
			config = conf
			return
		}
	}
	err = errors.New(fmt.Sprintf("not found monster %d", monsterID))
	return
}
