package config

type BuffConfig struct {
	BuffID     int32  `json:"BuffID"`
	Name       string `json:"Name"`
	Type       int32  `json:"Type"`
	DurTime    int32  `json:"DurTime"`
	Level      int32  `json:"Level"`
	AddType    int32  `json:"AddType"`
	EffectType int32  `json:"EffectType"`
	Value      int32  `json:"Value"`
	Func       string `json:"Func"`
}

type MapConfig struct {
	MapID     int32       `json:"MapID"`
	Name      string      `json:"Name"`
	Width     int32       `json:"Width"`
	Height    int32       `json:"Height"`
	BornX     int32       `json:"BornX"`
	BornY     int32       `json:"BornY"`
	MaxNum    int32       `json:"MaxNum"`
	UnWalkStr string      `json:"UnWalkStr"`
	UnWalk    []ConfigPos `json:"-"`
}

type ConfigPos struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type MonsterConfig struct {
	ID             int32   `json:"ID"`
	Name           string  `json:"Name"`
	Level          int32   `json:"Level"`
	PropID         int32   `json:"PropID"`
	PatrolDistance int32   `json:"PatrolDistance"`
	PursueDistance int32   `json:"PursueDistance"`
	AttackDistance int32   `json:"Attack_distance"`
	Skills         []int32 `json:"-"`
	RebornTime     int32   `json:"Reborn_time"`
}

type SkillConfig struct {
	SkillID        int32   `json:"SkillID"`
	Name           string  `json:"Name"`
	AttackParam    float64 `json:"AttackParam"`
	AttackDistance int32   `json:"AttackDistance"`
	CD             int32   `json:"CD"`
	Type           int32   `json:"Type"`
	TotalWave      int32   `json:"TotalWave"`
	WaveInterval   []int32 `json:"-"`
	DamageType     int32   `json:"DamageType"`
	RangeParams    []int32 `json:"-"`
	TargetNum      int32   `json:"TargetNum"`
	SelfBuffs      []int32 `json:"-"`
	TargetBuffs    []int32 `json:"-"`
}
