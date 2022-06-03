package resources

import "syscall/js"

type StructureKeeperLair struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureKeeperLair) iRef() js.Value {
	return s.ref
}

func (s *StructureKeeperLair) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureKeeperLair{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureKeeperLair) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureKeeperLair) x() int {
	return s.Pos().x()
}

func (s *StructureKeeperLair) y() int {
	return s.Pos().y()
}

func (s *StructureKeeperLair) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureKeeperLair) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureKeeperLair) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureKeeperLair) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureKeeperLair) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureKeeperLair) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureKeeperLair) StructureType() CStructure {
	return STRUCTURE_KEEPER_LAIR
}

func (s *StructureKeeperLair) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureKeeperLair) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureKeeperLair) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureKeeperLair) My() bool {
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureKeeperLair) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureKeeperLair) TicksToSpawn() int {
	return jsGet(s, "ticksToSpawn", getInt).(int)
}
