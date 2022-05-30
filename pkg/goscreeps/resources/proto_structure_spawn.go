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
	structureType CStructure

	my    bool
	owner string

	name     string
	spawning *Spawning
	store    *Store
}

type SpawnCreepOpts struct {
	Memory js.Value
	// TODO EnergyStructures []SpawnOrExtension
	DryRun     bool
	Directions []CDirection
}

func (s *StructureSpawn) iRef() js.Value {
	return s.ref
}

func (s *StructureSpawn) deRef(ref js.Value) IReference {
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
		s.pos = pos(s.ref)
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
		s.room = (&Room{}).deRef(s.ref).(*Room)
		s.cached["room"] = true
	}
	return s.room
}

func (s *StructureSpawn) Hits() int {
	if !s.cached["hits"] {
		s.hits = jsGet(s.ref, "hits").Int()
		s.cached["hits"] = true
	}
	return s.hits
}

func (s *StructureSpawn) HitsMax() int {
	if !s.cached["hitsMax"] {
		s.hitsMax = jsGet(s.ref, "hitsMax").Int()
		s.cached["hitsMax"] = true
	}
	return s.hitsMax
}

func (s *StructureSpawn) StructureType() CStructure {
	return STRUCTURE_SPAWN
}

func (s *StructureSpawn) Destroy() error {
	return destroy(s.ref)
}

func (s *StructureSpawn) IsActive() bool {
	return isActive(s.ref)
}

func (s *StructureSpawn) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(s.ref, enabled)
}

func (s *StructureSpawn) My() bool {
	if !s.cached["my"] {
		s.my = jsGet(s.ref, "my").Bool()
		s.cached["my"] = true
	}
	return s.my
}

func (s *StructureSpawn) Owner() string {
	if !s.cached["owner"] {
		s.owner = jsGet(s.ref, "owner").String()
		s.cached["owner"] = true
	}
	return s.owner
}

func (s *StructureSpawn) Id() string {
	if !s.cached["id"] {
		s.id = jsGet(s.ref, "id").String()
		s.cached["id"] = true
	}
	return s.id
}

func (s *StructureSpawn) Name() string {
	if !s.cached["name"] {
		s.name = jsGet(s.ref, "name").String()
		s.cached["name"] = true
	}
	return s.name
}

func (s *StructureSpawn) Spawning() *Spawning {
	if s.cached["spawning"] {
		s.spawning = (&Spawning{}).deRef(jsGet(s.ref, "spawning")).(*Spawning)
		s.cached["spawning"] = true
	}
	return s.spawning
}

func (s *StructureSpawn) Store() *Store {
	if !s.cached["store"] {
		s.store = (&Store{}).deRef(s.ref).(*Store)
		s.cached["store"] = true
	}
	return s.store
}

func (s *StructureSpawn) SpawnCreep(body CreepBody, name string, opts *SpawnCreepOpts) error {
	jsBody := make([]interface{}, len(body))
	for i, part := range body {
		jsBody[i] = string(part)
	}
	jsOpts := packSpawnCreepOpts(opts)
	result := jsCall(s.ref, "spawnCreep", jsBody, name, jsOpts).Int()
	return returnErr(result)
}

// TODO recycle, renewCreep

func packSpawnCreepOpts(opts *SpawnCreepOpts) js.Value {
	if opts == nil {
		return js.Undefined()
	}
	result := make(map[string]interface{}, 4)
	result["memory"] = opts.Memory
	// TODO result["energyStructures"] = opts.EnergyStructures
	result["dryRun"] = opts.DryRun
	if len(opts.Directions) == 0 {
		result["directions"] = js.Undefined()
	} else {
		result["directions"] = opts.Directions
	}
	return js.ValueOf(result)
}
