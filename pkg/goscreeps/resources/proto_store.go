package resources

import "syscall/js"

type Store struct {
	ref    js.Value
	cached map[string]bool

	contents map[ResourceConst]int
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
		cached: make(map[string]bool),
	}
}

func (s *Store) Contents() map[ResourceConst]int {
	if !s.cached["contents"] {
		entries := jsCall(jsObject, "entries", s.ref)
		length := jsGet(entries, "length").Int()
		result := make(map[ResourceConst]int, length)
		for i := 0; i < length; i++ {
			entry := entries.Index(i)
			key := entry.Index(0).String()
			value := entry.Index(1).Int()
			result[ResourceConst(key)] = value
		}
		s.contents = result
		s.cached["contents"] = true
	}
	return s.contents
}

func (s *Store) GetCapacity(resource *ResourceConst) int {
	var result js.Value
	if resource == nil {
		result = jsCall(s.ref, "getCapacity")
	} else {
		result = jsCall(s.ref, "getCapacity", string(*resource))
	}
	return result.Int()
}

func (s *Store) GetFreeCapacity(resource *ResourceConst) int {
	var result js.Value
	if resource == nil {
		result = jsCall(s.ref, "getFreeCapacity")
	} else {
		result = jsCall(s.ref, "getFreeCapacity", string(*resource))
	}
	return result.Int()
}

func (s *Store) GetUsedCapacity(resource *ResourceConst) int {
	var result js.Value
	if resource == nil {
		result = jsCall(s.ref, "getUsedCapacity")
	} else {
		result = jsCall(s.ref, "getUsedCapacity", string(*resource))
	}
	return result.Int()
}
