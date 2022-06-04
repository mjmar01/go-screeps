package resources

import "syscall/js"

// IStructure is the interface of any structure. (Ex: Spawn, Road, InvaderCore)
type IStructure interface {
	IDamageable

	StructureType() CStructure

	Destroy() error
	IsActive() bool
	NotifyWhenAttacked(enabled bool) error
}

// IDamageable is the interface of any IRoomObject that has hit points. Note that these objects aren't necessarily destructible like KeeperLairs
type IDamageable interface {
	IRoomObject

	Hits() int
	HitsMax() int
}

// IOwnedStructure is any IStructure that has an owner
type IOwnedStructure interface {
	IStructure
	IOwned
}

// IOwned is any IRoomObject that has an owner
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

// retrieve a store for a reference. Used for jsGet
func getStore(ref js.Value, property string) interface{} {
	storeRef := ref.Get(property)
	return (&Store{}).deRef(storeRef)
}
