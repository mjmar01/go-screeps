package resources

import "syscall/js"

type Spawning struct {
	ref    js.Value
	cached map[string]bool

	directions    []DirectionConst
	name          string
	needTime      int
	remainingTime int
	spawn         *StructureSpawn
}

func (s *Spawning) iRef() js.Value {
	return s.ref
}

func (s *Spawning) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Spawning{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (s *Spawning) Directions() []DirectionConst {
	// TODO arrays
	if s.cached["directions"] {
		jsDirections := jsGet(s.ref, "directions")
		var result []DirectionConst
		if !jsDirections.IsUndefined() {
			directionsCount := jsDirections.Length()
			result = make([]DirectionConst, directionsCount)
			for i := 0; i < directionsCount; i++ {
				result[i] = DirectionConst(jsDirections.Index(i).Int())
			}
		}
		s.directions = result
		s.cached["directions"] = true
	}
	return s.directions
}

func (s *Spawning) Name() string {
	if !s.cached["name"] {
		s.name = jsGet(s.ref, "name").String()
		s.cached["name"] = true
	}
	return s.name
}

func (s *Spawning) NeedTime() int {
	if !s.cached["needTime"] {
		s.needTime = jsGet(s.ref, "needTime").Int()
		s.cached["needTime"] = true
	}
	return s.needTime
}

func (s *Spawning) RemainingTime() int {
	if !s.cached["remainingTime"] {
		s.remainingTime = jsGet(s.ref, "remainingTime").Int()
		s.cached["remainingTime"] = true
	}
	return s.remainingTime
}

func (s *Spawning) Spawn() *StructureSpawn {
	if !s.cached["spawn"] {
		s.spawn = (&StructureSpawn{}).deRef(jsGet(s.ref, "spawn")).(*StructureSpawn)
		s.cached["spawn"] = true
	}
	return s.spawn
}

func (s *Spawning) Cancel() ScreepsError {
	result := jsCall(s.ref, "cancel").Int()
	return ReturnErr(result)
}

func (s *Spawning) SetDirections(directions []DirectionConst) ScreepsError {
	jsDirections := make([]interface{}, len(directions))
	for i, v := range directions {
		jsDirections[i] = int(v)
	}
	result := jsCall(s.ref, "setDirections", directions).Int()
	return ReturnErr(result)
}
