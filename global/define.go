package global

import "rpgMap/tool"

// Prop 属性
type Prop struct {
	Index   int32 `json:"Index"`
	MaxHp   int32 `json:"MaxHp"`
	Attack  int32 `json:"Attack"`
	Defense int32 `json:"defense"`
	Speed   int32 `json:"speed"`
}

// Pos 位置坐标
type Pos struct {
	X   float32
	Y   float32
	dir int16
}

// Point 地图点坐标
type Point struct {
	X int32
	Y int32
}

// Grid aoi map的key
type Grid struct {
	X int32
	Y int32
}

type ActorKey struct {
	ActorType int8
	ActorID   int64
}

// 地图格子的大小
const (
	GridLength int32 = 16
	GridHeight int32 = 16
)

// MapAreaKeyCache 缓存一个地图的格子数据
// 后续创建的时候直接读取，就不用再初始化一遍了
var MapAreaKeyCache = make(map[int32][]Grid)

var MapAreaNeighborsCache = make(map[int32]map[Grid][]Grid)

func GetGrid(gridX, gridY int32) Grid {
	return Grid{X: int32(gridX / GridLength), Y: int32(gridY / GridHeight)}
}
func GetGridByPos(pos Pos) Grid {
	gridX := int32(pos.X) / GridLength
	gridY := int32(pos.Y) / GridHeight
	return GetGrid(gridX, gridY)
}

func GetGridNeighbors(gridX, gridY, maxWidth, maxHeight int32) []Grid {
	xMin := tool.MaxInt32(gridX-1, 0)
	xMax := tool.MinInt32(gridX+1, maxWidth)
	yMin := tool.MaxInt32(gridY-1, 0)
	yMax := tool.MinInt32(gridY+1, maxHeight)
	//areas := make([]Grid, (xMax-xMin+1)*(yMax-yMin+1))
	var areas []Grid
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			grid := GetGrid(i, j)
			areas = append(areas, grid)
		}
	}
	return areas
}
