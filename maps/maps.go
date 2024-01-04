package maps

import (
	"fmt"
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
)

func StartMap(id int32, name string, line int32, mod *MapMod, args ...interface{}) *common.Pid {
	input := make(chan interface{}, 1000)
	output := make(chan interface{})
	pid := common.Pid{In: input, Out: output}
	go InitMap(&pid, id, name, line, mod, args)
	return &pid
}

func InitMap(pid *common.Pid, id int32, name string, line int32, mod *MapMod, args []interface{}) {
	common.Register(name, pid) // 注册到线程通道map内
	common.RegisterTimer(name, pid, 100)
	state := MakeMapState(pid, id, name, line, mod, args)
	state.mapHandleMsg(pid) // 消息监听

	fmt.Println("init map done", name)
	if id == 1 {
		fmt.Println("test find path")

		start := PNode{X: 45, Y: 10}
		end := PNode{X: 47, Y: 20}
		path := FindPath(state, &start, &end)
		fmt.Println("Find path : ", path)
	}

}

func MakeMapState(pid *common.Pid, id int32, name string, line int32, mod *MapMod, args []interface{}) *MapState {
	conf, _ := config.GetMapConfig(id)
	now2 := common.Now2()
	areas := InitAoi(conf)
	state := MapState{
		MapID:        id,
		Name:         name,
		Line:         line,
		Mod:          mod,
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

// MapLoop 地图轮询
func (state *MapState) MapLoop() {
	// todo 地图轮询
}
