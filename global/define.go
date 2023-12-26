package global

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
