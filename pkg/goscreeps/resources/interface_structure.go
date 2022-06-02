package resources

import "syscall/js"

type IStructure interface {
	IDamageable

	StructureType() CStructure

	Destroy() error
	IsActive() bool
	NotifyWhenAttacked(enabled bool) error
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

func destroy(ref js.Value) error {
	return returnErr(jsCall(ref, "destroy").Int())
}

func isActive(ref js.Value) bool {
	return jsCall(ref, "isActive").Bool()
}

func notifyWhenAttacked(ref js.Value, enabled bool) error {
	result := jsCall(ref, "notifyWhenAttacked", enabled).Int()
	return returnErr(result)
}

func getStore(ref js.Value, property string) interface{} {
	storeRef := ref.Get(property)
	return (&Store{}).deRef(storeRef)
}
