package resources

import "syscall/js"

type StructureExtension struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	hits          int
	hitsMax       int
	id            string
	structureType CStructure
	my            bool
	owner         string

	store *Store
}

func (s *StructureExtension) iRef() js.Value {
	return s.ref
}

func (s *StructureExtension) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureExtension{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *StructureExtension) x() int {
	return s.Pos().x()
}

func (s *StructureExtension) y() int {
	return s.Pos().y()
}

func (s *StructureExtension) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureExtension) Pos() *RoomPosition {
	if !s.cached["pos"] {
		s.pos = pos(s.ref)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *StructureExtension) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *StructureExtension) Room() *Room {
	if !s.cached["room"] {
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureExtension) My() bool {
	if !s.cached["my"] {
		s.my = jsGet(s.ref, "my").Bool()
		s.cached["my"] = true
	}
	return s.my
}

func (s *StructureExtension) Owner() string {
	if !s.cached["owner"] {
		s.owner = jsGet(s.ref, "owner").String()
		s.cached["owner"] = true
	}
	return s.owner
}

func (s *StructureExtension) Hits() int {
	if !s.cached["hits"] {
		s.hits = jsGet(s.ref, "hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureExtension) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = jsGet(s.ref, "hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureExtension) StructureType() CStructure {
	return STRUCTURE_EXTENSION
}

func (s *StructureExtension) Store() *Store {
	if !s.cached["store"] {
		s.store = (&Store{}).deRef(s.ref).(*Store)
		s.cached["store"] = true
	}
	return s.store
}

func (s *StructureExtension) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureExtension) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureExtension) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureExtension) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}
