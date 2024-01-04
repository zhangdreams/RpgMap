package mod

import "rpgMap/maps"

func init() {
	maps.ModMaps["mod_common"] = modCommon
}

var modCommon = &maps.MapMod{
	Init: func(state *maps.MapState, mapID int32, args ...interface{}) {
		// todo
	},
	RoleEnter: func(state *maps.MapState, roleID int64) {
		// todo
	},
	RoleExit: func(state *maps.MapState, roleID int64, offline bool) {
		// todo
	},
	RoleConnect: func(state *maps.MapState, roleID int64) {
		// todo
	},
	RoleDead: func(state *maps.MapState, roleID int64, killerType int8, killerID int64) {
		// todo
	},
	MonsterDead: func(state *maps.MapState, monsterID int64, killerType int8, killerID int64) {
		// todo
	},
	RoleRelive: func(state *maps.MapState, roleID int64) {
		// todo
	},
	Handle: func(state *maps.MapState, msg interface{}) {
		// todo
	},
}
