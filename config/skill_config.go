package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"rpgMap/tool"
)

var skillConfigs []SkillConfig

func (s *SkillConfig) UnmarshalJSON(data []byte) error {
	type Alias SkillConfig
	aux := &struct {
		WaveInterval string `json:"WaveInterval"`
		RangeParams  string `json:"RangeParams"`
		SelfBuffs    string `json:"SelfBuffs"`
		TargetBuffs  string `json:"TargetBuffs"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert comma-separated string fields to []int
	s.WaveInterval = tool.ConvertToIntArray(aux.WaveInterval)
	s.RangeParams = tool.ConvertToIntArray(aux.RangeParams)
	s.SelfBuffs = tool.ConvertToIntArray(aux.SelfBuffs)
	s.TargetBuffs = tool.ConvertToIntArray(aux.TargetBuffs)

	return nil
}

func ReadSkills() {
	filePath := "./config/json/skills.json"
	jsonConfig, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading map config file:", err)
		return
	}

	err = json.Unmarshal(jsonConfig, &skillConfigs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//showAllSkills()
}

func showAllSkills() {
	fmt.Println("show skill config")
	for _, skill := range skillConfigs {
		fmt.Println("SkillID:", skill.SkillID)
		fmt.Println("Name:", skill.Name)
		fmt.Println("AttackParam:", skill.AttackParam)
		fmt.Println("AttackDistance:", skill.AttackDistance)
		fmt.Println("CD:", skill.CD)
		fmt.Println("Type:", skill.Type)
		fmt.Println("TotalWave:", skill.TotalWave)
		fmt.Println("WaveInterval:", skill.WaveInterval)
		fmt.Println("DamageType:", skill.DamageType)
		fmt.Println("RangeParams:", skill.RangeParams)
		fmt.Println("TargetNum:", skill.TargetNum)
		fmt.Println("SelfBuffs:", skill.SelfBuffs)
		fmt.Println("TargetBuffs:", skill.TargetBuffs)
		fmt.Println("----------")
	}
}

// GetSkillConfig 返回指定skillID的配置
func GetSkillConfig(skillID int32) (config SkillConfig, err error) {
	for _, conf := range skillConfigs {
		if skillID == conf.SkillID {
			config = conf
			return
		}
	}
	err = errors.New(fmt.Sprintf("not found skill %d", skillID))
	return
}
