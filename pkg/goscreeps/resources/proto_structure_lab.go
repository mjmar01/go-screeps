package resources

import "syscall/js"

type StructureLab struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureLab) iRef() js.Value {
	return s.ref
}

func (s *StructureLab) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureLab{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureLab) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureLab) x() int {
	return s.Pos().x()
}

func (s *StructureLab) y() int {
	return s.Pos().y()
}

func (s *StructureLab) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureLab) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureLab) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureLab) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureLab) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureLab) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureLab) StructureType() CStructure {
	return STRUCTURE_LAB
}

func (s *StructureLab) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureLab) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureLab) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureLab) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureLab) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureLab) Cooldown() int {
	return jsGet(s, "cooldown", getInt).(int)
}

func (s *StructureLab) BoostCreep(target Creep, bodyPartCount int) error {
	jsCount := js.Undefined()
	if bodyPartCount > 0 {
		jsCount = js.ValueOf(bodyPartCount)
	}
	result := jsCall(s.ref, "boostCreep", target.iRef(), jsCount).Int()
	return returnErr(result)
}

func (s *StructureLab) ReverseReaction(lab1, lab2 *StructureLab) error {
	result := jsCall(s.ref, "reverseReaction", lab1.iRef(), lab2.iRef()).Int()
	return returnErr(result)
}

func (s *StructureLab) RunReaction(lab1, lab2 *StructureLab) error {
	result := jsCall(s.ref, "runReaction", lab1.iRef(), lab2.iRef()).Int()
	return returnErr(result)
}

func (s *StructureLab) UnBoostCreep(target Creep) error {
	result := jsCall(s.ref, "unBoostCreep", target.iRef()).Int()
	return returnErr(result)
}
