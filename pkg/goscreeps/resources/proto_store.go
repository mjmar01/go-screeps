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

func deRefStore(ref js.Value) *Store {
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
		entries := jsObject.Call("entries", s.ref)
		length := entries.Get("length").Int()
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
		result = s.ref.Call("getCapacity")
	} else {
		result = s.ref.Call("getCapacity", string(*resource))
	}
	return result.Int()
}

func (s *Store) GetFreeCapacity(resource *ResourceConst) int {
	var result js.Value
	if resource == nil {
		result = s.ref.Call("getFreeCapacity")
	} else {
		result = s.ref.Call("getFreeCapacity", string(*resource))
	}
	return result.Int()
}

func (s *Store) GetUsedCapacity(resource *ResourceConst) int {
	var result js.Value
	if resource == nil {
		result = s.ref.Call("getUsedCapacity")
	} else {
		result = s.ref.Call("getUsedCapacity", string(*resource))
	}
	return result.Int()
}

func getStore(ref js.Value) *Store {
	return deRefStore(ref.Get("store"))
}
