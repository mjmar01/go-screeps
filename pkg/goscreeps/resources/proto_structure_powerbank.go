package resources

import "syscall/js"

type StructurePowerBank struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructurePowerBank) iRef() js.Value {
	return s.ref
}

func (s *StructurePowerBank) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructurePowerBank{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructurePowerBank) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructurePowerBank) x() int {
	return s.Pos().x()
}

func (s *StructurePowerBank) y() int {
	return s.Pos().y()
}

func (s *StructurePowerBank) roomName() string {
	return s.Pos().roomName()
}

func (s *StructurePowerBank) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructurePowerBank) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructurePowerBank) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructurePowerBank) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructurePowerBank) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructurePowerBank) StructureType() CStructure {
	return STRUCTURE_POWER_BANK
}

func (s *StructurePowerBank) Destroy() error {
	return destroy(s.ref)
}

func (s *StructurePowerBank) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructurePowerBank) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructurePowerBank) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructurePowerBank) Power() int {
	return jsGet(s, "power", getInt).(int)
}

func (s *StructurePowerBank) TicksToDecay() int {
	return jsGet(s, "ticksToDecay", getInt).(int)
}
