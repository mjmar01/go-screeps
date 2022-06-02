package resources

import "syscall/js"

type Source struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *Source) iRef() js.Value {
	return s.ref
}

func (s *Source) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Source{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *Source) iCache() map[string]interface{} {
	return s.cached
}

func (s *Source) x() int {
	return s.Pos().x()
}

func (s *Source) y() int {
	return s.Pos().y()
}

func (s *Source) roomName() string {
	return s.Pos().roomName()
}

func (s *Source) Pos() *RoomPosition {
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *Source) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *Source) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *Source) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *Source) Energy() int {
	return jsGet(s, "energy", getInt).(int)
}

func (s *Source) EnergyCapacity() int {
	return jsGet(s, "energyCapacity", getInt).(int)
}

func (s *Source) TicksToRegeneration() int {
	return jsGet(s, "ticksToRegeneration", getInt).(int)
}
