package resources

import "syscall/js"

type Store struct {
	ref    js.Value
	cached map[string]bool
}

func (s *Store) iRef() js.Value {
	return s.ref
}

func (s *Store) CC() {
	s.cached = make(map[string]bool)
}

// TODO Get Set

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
