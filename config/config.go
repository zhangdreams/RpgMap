package config

import (
	"fmt"
	"rpgMap/common"
)

func InitConfig() {
	fmt.Println("read config")
	startTime := common.Now2()
	ReadMap()
	ReadMonster()
	ReadBuffs()
	ReadProps()
	ReadSkills()
	endTime := common.Now2()
	fmt.Println("read config used ", endTime-startTime, "ms")
}
