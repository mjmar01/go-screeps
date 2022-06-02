package resources

import "syscall/js"

type StructureContainer struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureContainer) iRef() js.Value {
	return s.ref
}

func (s *StructureContainer) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureContainer{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureContainer) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureContainer) x() int {
	return s.Pos().x()
}

func (s *StructureContainer) y() int {
	return s.Pos().y()
}

func (s *StructureContainer) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureContainer) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureContainer) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureContainer) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureContainer) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureContainer) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureContainer) StructureType() CStructure {
	return STRUCTURE_CONTAINER
}

func (s *StructureContainer) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureContainer) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureContainer) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureContainer) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureContainer) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureContainer) TicksToDecay() int {
	return jsGet(s, "ticksToDecay", getInt).(int)
}
