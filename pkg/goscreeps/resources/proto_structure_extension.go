package resources

import "syscall/js"

type StructureExtension struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureExtension) iRef() js.Value {
	return s.ref
}

func (s *StructureExtension) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureExtension{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureExtension) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureExtension) x() int {
	return s.Pos().x()
}

func (s *StructureExtension) y() int {
	return s.Pos().y()
}

func (s *StructureExtension) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureExtension) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureExtension) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureExtension) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureExtension) My() bool {
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureExtension) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureExtension) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureExtension) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureExtension) StructureType() CStructure {
	return STRUCTURE_EXTENSION
}

func (s *StructureExtension) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureExtension) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureExtension) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureExtension) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureExtension) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}
