package resources

import "syscall/js"

type StructureContainer struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	hits          int
	hitsMax       int
	id            string
	structureType StructureConst

	store        *Store
	ticksToDecay int
}

func (s *StructureContainer) iRef() js.Value {
	return s.ref
}

func (s *StructureContainer) deRef(ref js.Value) IRoomPosition {
	if ref.IsNull() {
		return nil
	}
	return &StructureContainer{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *StructureContainer) x() int {
	return s.Pos().x()
}

func (s *StructureContainer) y() int {
	return s.Pos().y()
}

func (s *StructureContainer) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureContainer) Pos() *RoomPosition {
	if !s.cached["pos"] {
		ref := s.ref.Get("pos")
		s.pos = (&RoomPosition{}).deRef(ref).(*RoomPosition)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *StructureContainer) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *StructureContainer) Room() *Room {
	if !s.cached["room"] {
		s.room = deRefRoom(s.ref.Get("room"))
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureContainer) Hits() int {
	if !s.cached["hits"] {
		s.hits = s.ref.Get("hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureContainer) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = s.ref.Get("hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureContainer) Id() string {
	if !s.cached["id"] {
		s.id = s.ref.Get("id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureContainer) StructureType() StructureConst {
	return STRUCTURE_CONTAINER
}

func (s *StructureContainer) Destroy() ScreepsError {
	return destroy(s.ref)
}

func (s *StructureContainer) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureContainer) NotifyWhenAttacked(enabled bool) ScreepsError {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureContainer) Store() *Store {
	if !s.cached["store"] {
		s.store = getStore(s.ref)
		s.cached["store"] = true
	}
	return s.store
}

func (s *StructureContainer) TicksToDecay() int {
	if !s.cached["ticksToDecay"] {
		s.ticksToDecay = s.ref.Get("ticksToDecay").Int()
		s.cached["ticksToDecay"] = true
	}
	return s.ticksToDecay
}
