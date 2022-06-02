package resources

import "syscall/js"

type StructureExtractor struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureExtractor) iRef() js.Value {
	return s.ref
}

func (s *StructureExtractor) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureExtractor{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureExtractor) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureExtractor) x() int {
	return s.Pos().x()
}

func (s *StructureExtractor) y() int {
	return s.Pos().y()
}

func (s *StructureExtractor) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureExtractor) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureExtractor) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureExtractor) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureExtractor) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureExtractor) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureExtractor) StructureType() CStructure {
	return STRUCTURE_EXTRACTOR
}

func (s *StructureExtractor) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureExtractor) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureExtractor) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureExtractor) My() bool {
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureExtractor) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureExtractor) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureExtractor) Cooldown() int {
	return jsGet(s, "cooldown", getInt).(int)
}
