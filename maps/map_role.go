package maps

// / ----- functions ---------

func (role *MapRole) GetType() int8 {
	return ACTOR_ROLE
}
func (role *MapRole) GetID() int64 {
	return role.ID
}
func (role *MapRole) IsAlive() bool {
	return role.State == 1
}
func (role *MapRole) GetCamp() int16 {
	return role.Camp
}
func (role *MapRole) GetBuffs() map[int32]*MapBuff {
	return role.Buffs
}
func (role *MapRole) SetBuffs(buffs map[int32]*MapBuff) {
	role.Buffs = buffs
}

// / ----- functions ---------
