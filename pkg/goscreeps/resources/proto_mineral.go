package resources

import "syscall/js"

type Mineral struct {
	ref    js.Value
	cached map[string]interface{}
}

func (m *Mineral) iRef() js.Value {
	return m.ref
}

func (m *Mineral) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Mineral{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (m *Mineral) iCache() map[string]interface{} {
	return m.cached
}

func (m *Mineral) x() int {
	return m.Pos().x()
}

func (m *Mineral) y() int {
	return m.Pos().x()
}

func (m *Mineral) roomName() string {
	return m.Pos().roomName()
}

func (m *Mineral) Pos() *RoomPosition {
	return jsGet(m, "pos", getPos).(*RoomPosition)
}

func (m *Mineral) Effects() []Effect {
	return jsGet(m, "effects", getEffects).([]Effect)
}

func (m *Mineral) Room() *Room {
	return jsGet(m, "room", getRoom).(*Room)
}

func (m *Mineral) MineralAmount() int {
	return jsGet(m, "mineralAmount", getInt).(int)
}

func (m *Mineral) MineralType() CResource {
	return CResource(jsGet(m, "mineralType", getString).(string))
}

func (m *Mineral) Density() CDensity {
	return CDensity(jsGet(m, "density", getInt).(int))
}

func (m *Mineral) Id() string {
	return jsGet(m, "id", getString).(string)
}

func (m *Mineral) TicksToRegeneration() int {
	return jsGet(m, "ticksToRegeneration", getInt).(int)
}
