package resources

import "syscall/js"

type Mineral struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room
	id      string

	density             CDensity
	mineralAmount       int
	mineralType         CResource
	ticksToRegeneration int
}

func (m *Mineral) iRef() js.Value {
	return m.ref
}

func (m *Mineral) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Deposit{
		ref:    ref,
		cached: make(map[string]bool),
	}
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
	if !m.cached["pos"] {
		m.pos = pos(m.ref)
		m.cached["pos"] = true
	}
	return m.pos
}

func (m *Mineral) Effects() []Effect {
	if !m.cached["effects"] {
		m.effects = effects(m.ref)
		m.cached["effects"] = true
	}
	return m.effects
}

func (m *Mineral) Room() *Room {
	if !m.cached["room"] {
		m.room = (&Room{}).deRef(m.ref).(*Room)
		m.cached["room"] = true
	}
	return m.room
}

func (m *Mineral) MineralAmount() int {
	if !m.cached["mineralAmount"] {
		m.mineralAmount = jsGet(m.ref, "mineralAmount").Int()
		m.cached["mineralAmount"] = true
	}
	return m.mineralAmount
}

func (m *Mineral) MineralType() CResource {
	if !m.cached["mineralType"] {
		m.mineralType = CResource(jsGet(m.ref, "mineralType").String())
		m.cached["mineralType"] = true
	}
	return m.mineralType
}

func (m *Mineral) Density() CDensity {
	if !m.cached["density"] {
		m.density = CDensity(jsGet(m.ref, "density").Int())
		m.cached["density"] = true
	}
	return m.density
}

func (m *Mineral) Id() string {
	if !m.cached["id"] {
		m.id = jsGet(m.ref, "id").String()
		m.cached["id"] = true
	}
	return m.id
}

func (m *Mineral) TicksToRegeneration() int {
	if !m.cached["ticksToRegeneration"] {
		m.ticksToRegeneration = jsGet(m.ref, "depositType").Int()
		m.cached["ticksToRegeneration"] = true
	}
	return m.ticksToRegeneration
}
