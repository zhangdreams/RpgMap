package maps

// / ----- functions ---------

func (monster *MapMonster) GetType() int8 {
	return ACTOR_ROLE
}
func (monster *MapMonster) GetID() int64 {
	return monster.ID
}
func (monster *MapMonster) IsAlive() bool {
	return monster.State == 1
}
func (monster *MapMonster) GetCamp() int16 {
	return monster.Camp
}
func (monster *MapMonster) GetBuffs() map[int32]*MapBuff {
	return monster.Buffs
}
func (monster *MapMonster) SetBuffs(buffs map[int32]*MapBuff) {
	monster.Buffs = buffs
}

// / ----- functions ---------
