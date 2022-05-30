package resources

import "syscall/js"

type StructureFactory struct {
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
	level    int
	store    *Store
}

func (s *StructureFactory) iRef() js.Value {
	return s.ref
}

func (s *StructureFactory) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &StructureFactory{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *StructureFactory) x() int {
	return s.Pos().x()
}

func (s *StructureFactory) y() int {
	return s.Pos().y()
}

func (s *StructureFactory) roomName() string {
	return s.Pos().roomName()
}

func (s *StructureFactory) Pos() *RoomPosition {
	if !s.cached["pos"] {
		s.pos = pos(s.ref)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *StructureFactory) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = effects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *StructureFactory) Room() *Room {
	if !s.cached["room"] {
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureFactory) Hits() int {
	if !s.cached["hits"] {
		s.hits = jsGet(s.ref, "hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureFactory) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = jsGet(s.ref, "hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureFactory) StructureType() CStructure {
	return STRUCTURE_FACTORY
}

func (s *StructureFactory) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureFactory) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureFactory) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureFactory) My() bool {
	if !s.cached["my"] {
		s.my = jsGet(s.ref, "my").Bool()
		s.cached["my"] = true
	}
	return s.my
}

func (s *StructureFactory) Owner() string {
	if !s.cached["owner"] {
		s.owner = jsGet(s.ref, "owner").String()
		s.cached["owner"] = true
	}
	return s.owner
}

func (s *StructureFactory) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureFactory) Store() *Store {
	if !s.cached["store"] {
		s.store = (&Store{}).deRef(s.ref).(*Store)
		s.cached["store"] = true
	}
	return s.store
}

func (s *StructureFactory) Cooldown() int {
	if !s.cached["cooldown"] {
		s.cooldown = jsGet(s.ref, "cooldown").Int()
		s.cached["cooldown"] = true
	}
	return s.cooldown
}

func (s *StructureFactory) Level() int {
	if !s.cached["level"] {
		s.level = jsGet(s.ref, "level").Int()
		s.cached["level"] = true
	}
	return s.level
}

func (s *StructureFactory) Produce(product CResource) error {
	result := jsCall(s.ref, "produce", string(product)).Int()
	return returnErr(result)
}
