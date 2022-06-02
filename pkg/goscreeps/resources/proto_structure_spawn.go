package resources

import "syscall/js"

type StructureSpawn struct {
	ref    js.Value
	cached map[string]interface{}
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
		cached: make(map[string]interface{}),
	}
}

func (s *StructureSpawn) iCache() map[string]interface{} {
	return s.cached
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
	return jsGet(s, "pos", getPos).(*RoomPosition)
}

func (s *StructureSpawn) Effects() []Effect {
	return jsGet(s, "effects", getEffects).([]Effect)
}

func (s *StructureSpawn) Room() *Room {
	return jsGet(s, "room", getRoom).(*Room)
}

func (s *StructureSpawn) Hits() int {
	return jsGet(s, "hits", getInt).(int)
}

func (s *StructureSpawn) HitsMax() int {
	return jsGet(s, "hitsMax", getInt).(int)
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
	return jsGet(s, "my", getBool).(bool)
}

func (s *StructureSpawn) Owner() string {
	return jsGet(s, "owner", getString).(string)
}

func (s *StructureSpawn) Id() string {
	return jsGet(s, "id", getString).(string)
}

func (s *StructureSpawn) Name() string {
	return jsGet(s, "name", getString).(string)
}

func (s *StructureSpawn) Spawning() *Spawning {
	return jsGet(s, "spawning", func(ref js.Value, property string) interface{} {
		return (&Spawning{}).deRef(ref.Get("spawning"))
	}).(*Spawning)
}

func (s *StructureSpawn) Store() *Store {
	return jsGet(s, "store", getStore).(*Store)
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
