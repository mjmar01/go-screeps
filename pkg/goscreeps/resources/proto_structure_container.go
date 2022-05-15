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

func (s *StructureContainer) deRef(ref js.Value) IReference {
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
		s.pos = pos(s.ref)
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
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureContainer) Hits() int {
	if !s.cached["hits"] {
		s.hits = jsGet(s.ref, "hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureContainer) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = jsGet(s.ref, "hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureContainer) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
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
		s.store = (&Store{}).deRef(s.ref).(*Store)
		s.cached["store"] = true
	}
	return s.store
}

func (s *StructureContainer) TicksToDecay() int {
	if !s.cached["ticksToDecay"] {
		s.ticksToDecay = jsGet(s.ref, "ticksToDecay").Int()
		s.cached["ticksToDecay"] = true
	}
	return s.ticksToDecay
}
