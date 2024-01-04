package maps

import (
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
)

// 实例类型
const (
	ACTOR_ROLE    = 1 // 玩家
	ACTOR_MONSTER = 2
)

type MapActor struct {
	ActorType int8
	ActorID   int64
	State     int8
	IsMove    bool
	BaseProp  *global.Prop
	TotalProp *global.Prop
}

// MapRole 地图中玩家信息，基本是需要用于同步的
// todo 后续可根据需求增加
type MapRole struct {
	ID     int64
	Name   string
	Level  int32
	HP     int32
	MaxHP  int32
	State  int8
	Pos    *global.Pos
	TarPos *global.Pos
	Camp   int16
	Buffs  map[int32]*MapBuff
}

// MapMonster 地图中怪物信息，基本是需要用于同步的
// todo 后续可根据需求增加
type MapMonster struct {
	ID     int64
	Name   string
	Level  int32
	HP     int32
	MaxHP  int32
	State  int8
	Pos    *global.Pos
	TarPos *global.Pos
	Camp   int16
	Path   []global.Point
	Buffs  map[int32]*MapBuff
}

type MapState struct {
	MapID        int32
	Name         string
	Line         int32
	Mod          *MapMod
	CreateTime   int64
	LastTickTime int64
	Pid          *common.Pid
	Actors       map[global.ActorKey]*MapActor
	Roles        map[int64]*MapRole
	Monsters     map[int64]*MapMonster
	Config       config.MapConfig
	Areas        *map[global.Grid]*[]global.ActorKey
}

type MapInfo interface {
	GetType() int8
	GetID() int64
	IsAlive() bool
	GetCamp() int16
	GetBuffs() map[int32]*MapBuff
	SetBuffs(map[int32]*MapBuff)
}

type MapBuff struct {
	BuffID  int32
	SrcType int8
	SrcID   int64
	EndTime int64
	Value   int32
}

type MapMod struct {
	// 玩法地图初始化
	Init func(state *MapState, mapID int32, args ...interface{})
	// 玩家进入地图
	RoleEnter func(state *MapState, roleID int64)
	// 玩家离开地图 offline表示是否为是否下线
	RoleExit func(state *MapState, roleID int64, offline bool)
	// 玩家重连进游戏
	RoleConnect func(state *MapState, roleID int64)
	// 玩家死亡回调
	RoleDead func(state *MapState, roleID int64, killerType int8, killerID int64)
	// 怪物死亡回调
	MonsterDead func(state *MapState, monsterID int64, killerType int8, killerID int64)
	// 玩家复活回调
	RoleRelive func(state *MapState, roleID int64)
	// 指定模块的消息处理
	Handle func(state *MapState, msg interface{})
}
