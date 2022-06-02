package resources

import "syscall/js"

type StructureFactory struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureFactory) iRef() js.Value {
	return s.ref
}

func (s *StructureFactory) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureFactory{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureFactory) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureFactory) x() int {
	return s.Pos().x()
}

func (s *StructureFactory) y() int {
	return s.Pos().y()
}

func (s *StructureFactory) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureFactory) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)

}

func (s *StructureFactory) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureFactory) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureFactory) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureFactory) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureFactory) StructureType() CStructure {
	return STRUCTURE_FACTORY
}

func (s *StructureFactory) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureFactory) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureFactory) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureFactory) My() bool {
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureFactory) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureFactory) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureFactory) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureFactory) Cooldown() int {
	return jsGet(s, "cooldown", getInt).(int)
}

func (s *StructureFactory) Level() int {
	return jsGet(s, "level", getInt).(int)
}

func (s *StructureFactory) Produce(product CResource) error {
	result := jsCall(s.ref, "produce", string(product)).Int()
	return returnErr(result)
}
