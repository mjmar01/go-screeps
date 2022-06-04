package resources

import "syscall/js"

type StructureLink struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureLink) iRef() js.Value {
	return s.ref
}

func (s *StructureLink) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureLink{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureLink) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureLink) x() int {
	return s.Pos().x()
}

func (s *StructureLink) y() int {
	return s.Pos().y()
}

func (s *StructureLink) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureLink) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureLink) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureLink) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureLink) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureLink) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureLink) StructureType() CStructure {
	return STRUCTURE_LINK
}

func (s *StructureLink) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureLink) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureLink) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureLink) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureLink) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureLink) Cooldown() int {
	return jsGet(s, "cooldown", getInt).(int)
}

func (s *StructureLink) TransferEnergy(target *StructureLink, amount int) error {
	jsAmount := js.Undefined()
	if amount == -1 {
		jsAmount = js.ValueOf(amount)
	}
	result := jsCall(s.ref, "transferEnergy", target.iRef(), jsAmount).Int()
	return returnErr(result)
}
