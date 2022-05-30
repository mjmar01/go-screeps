package resources

import "syscall/js"

type Source struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	id                  string
	energy              int
	energyCapacity      int
	ticksToRegeneration int
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
		cached: make(map[string]bool),
	}
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
	if !s.cached["pos"] {
		s.pos = pos(s.ref)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *Source) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *Source) Room() *Room {
	if !s.cached["room"] {
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *Source) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *Source) Energy() int {
	if !s.cached["energy"] {
		s.energy = jsGet(s.ref, "energy").Int()
		s.cached["energy"] = true
	}
	return s.energy
}

func (s *Source) EnergyCapacity() int {
	if !s.cached["energyCapacity"] {
		s.energyCapacity = jsGet(s.ref, "energyCapacity").Int()
		s.cached["energyCapacity"] = true
	}
	return s.energyCapacity
}

func (s *Source) TicksToRegeneration() int {
	if !s.cached["ticksToRegeneration"] {
		s.ticksToRegeneration = jsGet(s.ref, "ticksToRegeneration").Int()
		s.cached["ticksToRegeneration"] = true
	}
	return s.ticksToRegeneration
}
