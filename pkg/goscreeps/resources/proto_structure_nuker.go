package resources

import "syscall/js"

type StructureNuker struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureNuker) iRef() js.Value {
	return s.ref
}

func (s *StructureNuker) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureNuker{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureNuker) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureNuker) x() int {
	return s.Pos().x()
}

func (s *StructureNuker) y() int {
	return s.Pos().y()
}

func (s *StructureNuker) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureNuker) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureNuker) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureNuker) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureNuker) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureNuker) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureNuker) StructureType() CStructure {
	return STRUCTURE_NUKER
}

func (s *StructureNuker) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureNuker) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureNuker) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureNuker) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureNuker) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
}

func (s *StructureNuker) Cooldown() int {
	return jsGet(s, "cooldown", getInt).(int)
}

func (s *StructureNuker) LaunchNuke(target IRoomPosition) error {
	result := jsCall(s.ref, "launchNuke", target.Pos().iRef()).Int()
	return returnErr(result)
}
