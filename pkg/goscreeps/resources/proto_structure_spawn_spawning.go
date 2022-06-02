package resources

import "syscall/js"

type Spawning struct {
	ref    js.Value
	cached map[string]interface{}
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
		cached: make(map[string]interface{}),
	}
}

func (s *Spawning) iCache() map[string]interface{} {
	return s.cached
}

func (s *Spawning) Directions() []CDirection {
	return jsGet(s, "directions", func(ref js.Value, property string) interface{} {
		jsDirections := ref.Get(property)
		var result []CDirection
		if !jsDirections.IsUndefined() {
			directionsCount := jsDirections.Length()
			result = make([]CDirection, directionsCount)
			for i := 0; i < directionsCount; i++ {
				result[i] = CDirection(jsDirections.Index(i).Int())
			}
		}
		return result
	}).([]CDirection)
}

func (s *Spawning) Name() string {
	return jsGet(s, "name", getString).(string)
}

func (s *Spawning) NeedTime() int {
	return jsGet(s, "needTime", getInt).(int)
}

func (s *Spawning) RemainingTime() int {
	return jsGet(s, "remainingTime", getInt).(int)
}

func (s *Spawning) Spawn() *StructureSpawn {
	return jsGet(s, "spawn", func(ref js.Value, property string) interface{} {
		return (&StructureSpawn{}).deRef(ref.Get(property))
	}).(*StructureSpawn)
}

func (s *Spawning) Cancel() error {
	result := jsCall(s.ref, "cancel").Int()
	return returnErr(result)
}

func (s *Spawning) SetDirections(directions []CDirection) error {
	jsDirections := make([]interface{}, len(directions))
	for i, v := range directions {
		jsDirections[i] = int(v)
	}
	result := jsCall(s.ref, "setDirections", directions).Int()
	return returnErr(result)
}
