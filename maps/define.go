package maps

import (
	"rpgMap/common"
	"rpgMap/config"
	"rpgMap/global"
)

type MapActor struct {
	ActorType int8
	ActorID   int64
	State     int8
	IsMove    bool
	BaseProp  *global.Prop
	TotalProp *global.Prop
	Buffs     map[int32]*MapBuff
}

// MapRole 地图中玩家信息，基本是需要用于同步的
// todo 后续可根据需求增加
type MapRole struct {
	ID     int64
	Name   string
	Level  int32
	HP     int32
	MaxHP  int32
	Pos    *global.Pos
	TarPos *global.Pos
	Camp   int16
}

// MapMonster 地图中怪物信息，基本是需要用于同步的
// todo 后续可根据需求增加
type MapMonster struct {
	ID     int64
	Name   string
	Level  int32
	HP     int32
	MaxHP  int32
	Pos    *global.Pos
	TarPos *global.Pos
	Camp   int16
	Path   []global.Point
}

type MapState struct {
	MapID        int32
	Name         string
	Line         int32
	CreateTime   int64
	LastTickTime int64
	Pid          *common.Pid
	Actors       map[global.ActorKey]*MapActor
	Config       config.MapConfig
	Areas        *map[global.Grid]*[]global.ActorKey
}

type MapBuff struct {
	BuffID  int32
	SrcType int8
	SrcID   int64
	EndTime int64
	Value   int32
}
