package maps

import (
	"fmt"
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
)

func StartMap(args ...interface{}) *common.Pid {
	input := make(chan interface{}, 1000)
	output := make(chan interface{})
	pid := common.Pid{In: input, Out: output}
	go InitMap(&pid, args)
	return &pid
}

func InitMap(pid *common.Pid, args []interface{}) {
	name := args[1].(string)
	common.Register(name, pid) // 注册到线程通道map内
	common.RegisterTimer(name, pid, 100)

	fmt.Println("init map done", name)
	if args[0].(int32) == 1 {
		fmt.Println("test find path")
		state := MakeMapState(pid, args)
		start := PNode{X: 45, Y: 10}
		end := PNode{X: 47, Y: 20}
		path := FindPath(state, &start, &end)
		fmt.Println("Find path : ", path)
	}

}

func MakeMapState(pid *common.Pid, args []interface{}) *MapState {
	id := args[0].(int32)
	name := args[1].(string)
	line := args[2].(int32)
	conf, _ := config.GetMapConfig(args[0].(int32))
	now2 := common.Now2()
	areas := InitAoi(conf)
	state := MapState{
		MapID:        id,
		Name:         name,
		Line:         line,
		Pid:          pid,
		CreateTime:   now2,
		LastTickTime: now2,
		Config:       conf,
		Areas:        areas,
	}
	return &state
}

// GetMapInfo 返回一个地图内的对象
func (state *MapState) GetMapInfo(actorType int8, actorID int64) MapInfo {
	if actorType == ACTOR_ROLE {
		return state.Roles[actorID]
	}
	return state.Monsters[actorID]
}
func (state *MapState) GetMapInfoByKey(key global.ActorKey) MapInfo {
	return state.GetMapInfo(key.ActorType, key.ActorID)
}
