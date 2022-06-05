package resources

import "syscall/js"

type Store struct {
	ref    js.Value
	cached map[string]interface{}
}

func (s *Store) iRef() js.Value {
	return s.ref
}

func (s *Store) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Store{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (s *Store) iCache() map[string]interface{} {
	return s.cached
}

func (s *Store) Contents() map[CResource]int {
	return jsGet(s, "entries", func(ref js.Value, property string) interface{} {
		entries := jsCall(jsObject, property, ref)
		length := entries.Get("length").Int()
		result := make(map[CResource]int, length)
		for i := 0; i < length; i++ {
			entry := entries.Index(i)
			key := entry.Index(0).String()
			value := entry.Index(1).Int()
			result[CResource(key)] = value
		}
		return result
	}).(map[CResource]int)
}

func (s *Store) GetCapacity(resource CResource) int {
	var result js.Value
	if resource == RESOURCE_ANY {
		result = jsCall(s.ref, "getCapacity")
	} else {
		result = jsCall(s.ref, "getCapacity", string(resource))
	}
	return result.Int()
}

func (s *Store) GetFreeCapacity(resource CResource) int {
	var result js.Value
	if resource == RESOURCE_ANY {
		result = jsCall(s.ref, "getFreeCapacity")
	} else {
		result = jsCall(s.ref, "getFreeCapacity", string(resource))
	}
	return result.Int()
}

func (s *Store) GetUsedCapacity(resource CResource) int {
	var result js.Value
	if resource == RESOURCE_ANY {
		result = jsCall(s.ref, "getUsedCapacity")
	} else {
		result = jsCall(s.ref, "getUsedCapacity", string(resource))
	}
	return result.Int()
}
