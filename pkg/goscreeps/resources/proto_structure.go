package resources

import "syscall/js"

type Structure struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	hits          int
	hitsMax       int
	id            string
	structureType StructureConst
}

func (s *Structure) iRef() js.Value {
	return s.ref
}

func (s *Structure) CC() {
	s.cached = make(map[string]bool)
}

func (s *Structure) deRef(ref js.Value) IRoomPosition {
	if ref.IsNull() {
		return nil
	}
	return &Structure{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *Structure) x() int {
	return s.Pos().x()
}

func (s *Structure) y() int {
	return s.Pos().y()
}

func (s *Structure) roomName() string {
	return s.Pos().roomName()
}

func (s *Structure) Pos() *RoomPosition {
	if !s.cached["pos"] {
		ref := s.ref.Get("pos")
		s.pos = (&RoomPosition{}).deRef(ref).(*RoomPosition)
		s.cached["pos"] = true
	}
	return s.pos
}

func (s *Structure) Effects() []Effect {
	if !s.cached["effects"] {
		s.effects = getEffects(s.ref)
		s.cached["effects"] = true
	}
	return s.effects
}

func (s *Structure) Room() *Room {
	if !s.cached["room"] {
		s.room = deRefRoom(s.ref.Get("room"))
		s.cached["room"] = true
	}
	return s.room
}

func (s *Structure) Hits() int {
	if !s.cached["hits"] {
		s.hits = s.ref.Get("hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *Structure) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = s.ref.Get("hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *Structure) Id() string {
	if !s.cached["id"] {
		s.id = s.ref.Get("id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *Structure) StructureType() StructureConst {
	if !s.cached["structureType"] {
		s.structureType = StructureConst(s.ref.Get("structureType").String())
		s.cached["structureType"] = true
	}
	return s.structureType
}

func (s *Structure) Destroy() ScreepsError {
	result := s.ref.Call("destroy").Int()
	return ReturnErr(ErrorCode(result))
}

func (s *Structure) IsActive() bool {
	return s.ref.Call("isActive").Bool()
}

func (s *Structure) NotifyWhenAttacked(enabled bool) ScreepsError {
	result := s.ref.Call("notifyWhenAttacked", enabled).Int()
	return ReturnErr(ErrorCode(result))
}
