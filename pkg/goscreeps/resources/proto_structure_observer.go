package resources

import "syscall/js"

type StructureObserver struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *StructureObserver) iRef() js.Value {
	return s.ref
}

func (s *StructureObserver) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureObserver{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *StructureObserver) iCache() map[string]interface{} {
	return s.cached
}

func (s *StructureObserver) x() int {
	return s.Pos().x()
}

func (s *StructureObserver) y() int {
	return s.Pos().y()
}

func (s *StructureObserver) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureObserver) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureObserver) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureObserver) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureObserver) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureObserver) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
}

func (s *StructureObserver) StructureType() CStructure {
	return STRUCTURE_OBSERVER
}

func (s *StructureObserver) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureObserver) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureObserver) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureObserver) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureObserver) ObserverRoom(roomName string) error {
	result := jsCall(s.ref, "observerRoom", roomName).Int()
	return returnErr(result)
}
