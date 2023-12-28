package maps

import (
	"rpgMap/config"
	"rpgMap/global"
)

func InitAoi(conf config.MapConfig) *map[global.Grid]*[]global.ActorKey {
	areaKeys := global.MapAreaKeyCache[conf.MapID] // 直接从缓存中取
	var areas = make(map[global.Grid]*[]global.ActorKey)
	for _, keys := range areaKeys {
		areas[keys] = &[]global.ActorKey{}
	}
	return &areas
}

// GetGridActorsByKey 根据一个格子key返回格子内对象
func GetGridActorsByKey(state *MapState, grid global.Grid) *[]global.ActorKey {
	areas := *state.Areas
	return areas[grid]
}

// GetGridActorsByPos 根据一个坐标返回一个格子内的对象
func GetGridActorsByPos(state *MapState, pos global.Pos) *[]global.ActorKey {
	grid := global.GetGridByPos(pos)
	return GetGridActorsByKey(state, grid)
}

// GetGridActors 返回一个格子内的对象
func GetGridActors(state *MapState, gridX, gridY int32) *[]global.ActorKey {
	grid := global.GetGrid(gridX, gridY)
	return GetGridActorsByKey(state, grid)
}

// GetAoiActors 返回一个格子所在九宫格内的对象
func GetAoiActors(state *MapState, gridX, gridY int32) (actorKeys []global.ActorKey) {
	gridNeighborMaps := global.MapAreaNeighborsCache[state.MapID]
	gridKey := global.GetGrid(gridX, gridY)
	gridNeighbors := gridNeighborMaps[gridKey]
	for _, grid := range gridNeighbors {
		aKeys := GetGridActorsByKey(state, grid)
		actorKeys = append(actorKeys, *aKeys...)
	}
	return
}

// GetAoiActorsByPos 根据坐标返回所在九宫格内的对象
func GetAoiActorsByPos(state *MapState, pos global.Pos) (actorKeys []global.ActorKey) {
	gridX := int32(pos.X) / global.GridLength
	gridY := int32(pos.Y) / global.GridHeight
	return GetAoiActors(state, gridX, gridY)
}

// EnterArea 进入一个格子
func EnterArea(state *MapState, grid global.Grid, key global.ActorKey) {
	keys := GetGridActorsByKey(state, grid)
	*keys = append(*keys, key)
}

// ExitArea 离开一个格子
func ExitArea(state *MapState, grid global.Grid, key global.ActorKey) {
	keys := GetGridActorsByKey(state, grid)
	for i, v := range *keys {
		if v == key {
			(*keys)[i] = (*keys)[len(*keys)-1] // 这里不用在乎顺序
			*keys = (*keys)[:len(*keys)-1]
		}
	}
}

func DoUpPos(state *MapState, key global.ActorKey, pos1, pos2 global.Pos) {
	grid1 := global.GetGridByPos(pos1)
	grid2 := global.GetGridByPos(pos2)
	if grid1 != grid2 {
		ExitArea(state, grid1, key)
		EnterArea(state, grid2, key)
		// todo 视野刷新
	}
}
