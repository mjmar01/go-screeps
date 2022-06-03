package resources

import "syscall/js"

type StructureInvaderCore struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureInvaderCore) iRef() js.Value {
	return s.ref
}

func (s *StructureInvaderCore) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureInvaderCore{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureInvaderCore) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureInvaderCore) x() int {
	return s.Pos().x()
}

func (s *StructureInvaderCore) y() int {
	return s.Pos().y()
}

func (s *StructureInvaderCore) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureInvaderCore) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureInvaderCore) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureInvaderCore) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureInvaderCore) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureInvaderCore) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureInvaderCore) StructureType() CStructure {
	return STRUCTURE_INVADER_CORE
}

func (s *StructureInvaderCore) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureInvaderCore) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureInvaderCore) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureInvaderCore) My() bool {
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureInvaderCore) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureInvaderCore) Level() int {
	return jsGet(s, "level", getInt).(int)
}

func (s *StructureInvaderCore) TicksToDeploy() int {
	return jsGet(s, "ticksToDeploy", getInt).(int)
}

func (s *StructureInvaderCore) Spawning() *Spawning {
	return jsGet(s, "spawning", func(ref js.Value, property string) interface{} {
		return (&Spawning{}).deRef(ref.Get(property))
	}).(*Spawning)
}
