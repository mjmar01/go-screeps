package resources

import "syscall/js"

type IStructure interface {
	IDamageable

	StructureType() StructureConst

	Destroy() ScreepsError
	IsActive() bool
	NotifyWhenAttacked(enabled bool) ScreepsError
}

type IDamageable interface {
	IRoomObject

	Hits() int
	HitsMax() int
}

type IOwnedStructure interface {
	IStructure

	My() bool
	Owner() string
}

type IOwned interface {
	IRoomObject

	My() bool
	Owner() string
}

func destroy(ref js.Value) ScreepsError {
	return ReturnErr(jsCall(ref, "destroy").Int())
}

func isActive(ref js.Value) bool {
	return jsCall(ref, "isActive").Bool()
}

func notifyWhenAttacked(ref js.Value, enabled bool) ScreepsError {
	result := jsCall(ref, "notifyWhenAttacked", enabled).Int()
	return ReturnErr(result)
}
