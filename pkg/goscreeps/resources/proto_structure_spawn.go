package resources

import "syscall/js"

type StructureSpawn struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	hits          int
	hitsMax       int
	id            string
	structureType StructureConst

	my    bool
	owner string

	name string
	// spawning *StructureSpawning
	store *Store
}

func (s *StructureSpawn) iRef() js.Value {
	return s.ref
}

func (s *StructureSpawn) deRef(ref js.Value) IRoomPosition {
	if ref.IsNull() {
		return nil
	}
	return &StructureSpawn{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *StructureSpawn) x() int {
	return s.Pos().x()
}

func (s *StructureSpawn) y() int {
	return s.Pos().y()
}

func (s *StructureSpawn) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureSpawn) Pos() *RoomPosition {
	if !s.cached["pos"] {
		ref := s.ref.Get("pos")
		s.pos = (&RoomPosition{}).deRef(ref).(*RoomPosition)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *StructureSpawn) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *StructureSpawn) Room() *Room {
	if !s.cached["room"] {
		s.room = deRefRoom(s.ref.Get("room"))
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureSpawn) Hits() int {
	if !s.cached["hits"] {
		s.hits = s.ref.Get("hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureSpawn) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = s.ref.Get("hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureSpawn) Id() string {
	if !s.cached["id"] {
		s.id = s.ref.Get("id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureSpawn) StructureType() StructureConst {
	return STRUCTURE_SPAWN
}

func (s *StructureSpawn) Destroy() ScreepsError {
	return destroy(s.ref)
}

func (s *StructureSpawn) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureSpawn) NotifyWhenAttacked(enabled bool) ScreepsError {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureSpawn) My() bool {
	return my(s.ref)
}

func (s *StructureSpawn) Owner() string {
	return owner(s.ref)
}
