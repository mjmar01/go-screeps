package resources

import "syscall/js"

type StructureExtractor struct {
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

	cooldown int
}

func (s *StructureExtractor) iRef() js.Value {
	return s.ref
}

func (s *StructureExtractor) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureExtractor{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *StructureExtractor) x() int {
	return s.Pos().x()
}

func (s *StructureExtractor) y() int {
	return s.Pos().y()
}

func (s *StructureExtractor) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureExtractor) Pos() *RoomPosition {
	if !s.cached["pos"] {
		s.pos = pos(s.ref)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *StructureExtractor) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *StructureExtractor) Room() *Room {
	if !s.cached["room"] {
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureExtractor) Hits() int {
	if !s.cached["hits"] {
		s.hits = jsGet(s.ref, "hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureExtractor) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = jsGet(s.ref, "hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureExtractor) StructureType() CStructure {
	return STRUCTURE_EXTRACTOR
}

func (s *StructureExtractor) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureExtractor) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureExtractor) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureExtractor) My() bool {
	if !s.cached["my"] {
		s.my = jsGet(s.ref, "my").Bool()
		s.cached["my"] = true
	}
	return s.my
}

func (s *StructureExtractor) Owner() string {
	if !s.cached["owner"] {
		s.owner = jsGet(s.ref, "owner").String()
		s.cached["owner"] = true
	}
	return s.owner
}

func (s *StructureExtractor) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureExtractor) Cooldown() int {
	if !s.cached["cooldown"] {
		s.cooldown = jsGet(s.ref, "cooldown").Int()
		s.cached["cooldown"] = true
	}
	return s.cooldown
}
