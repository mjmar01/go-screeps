package resources

import "syscall/js"

type IStructure interface {
	IRoomObject

	Hits() int
	HitsMax() int
	Id() string
	StructureType() StructureConst

	Destroy() ScreepsError
	IsActive() bool
	NotifyWhenAttacked(enabled bool) ScreepsError
}

type IOwnedStructure interface {
	IStructure

	My() bool
	Owner() string
}

func destroy(ref js.Value) ScreepsError {
	return ReturnErr(ErrorCode(ref.Call("destroy").Int()))
}

func isActive(ref js.Value) bool {
	return ref.Call("isActive").Bool()
}

func notifyWhenAttacked(ref js.Value, enabled bool) ScreepsError {
	result := ref.Call("notifyWhenAttacked", enabled).Int()
	return ReturnErr(ErrorCode(result))
}

func my(ref js.Value) bool {
	return ref.Get("my").Bool()
}

func owner(ref js.Value) string {
	return ref.Get("owner").String()
}
