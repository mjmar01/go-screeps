package resources

import (
	"syscall/js"
)

type RoomPosition struct {
	ref    js.Value
	cached map[string]interface{}
}

var roomPositionConstructor = js.Global().Get("RoomPosition")

func (r *RoomPosition) iRef() js.Value {
	return r.ref
}

func (r *RoomPosition) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &RoomPosition{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (r *RoomPosition) iCache() map[string]interface{} {
	return r.cached
}

func (r *RoomPosition) x() int {
	return jsGet(r, "x", getInt).(int)
}

func (r *RoomPosition) y() int {
	return jsGet(r, "y", getInt).(int)
}

func (r *RoomPosition) roomName() string {
	return jsGet(r, "roomName", getString).(string)
}

func (r *RoomPosition) X() int {
	return r.x()
}

func (r *RoomPosition) Y() int {
	return r.y()
}

func (r *RoomPosition) RoomName() string {
	return r.roomName()
}

func (r *RoomPosition) CreateConstructionSite(sType CStructure, name string) error {
	return createConstructionSite(r, sType, name)
}

func (r *RoomPosition) CreateFlag(name string, primary CColor, secondary CColor) (string, error) {
	return createFlag(r, name, primary, secondary)
}

func (r *RoomPosition) FindClosestTypeByPath(fType CFind, opts *FindClosestByPathOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindClosestPosByPath(targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindClosestTypeByRange(fType CFind, opts *FindFilterOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindClosestPosByRange(targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindTypeInRange(fType CFind, maxRange int, opts *FindFilterOpts) []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindPosInRange(targets []IRoomPosition, maxRange int, opts *FindFilterOpts) []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindPathTo(target IRoomPosition, opts *FindPathOpts) Path {
	panic("TODO")
}

func (r *RoomPosition) GetDirectionTo(target IRoomPosition) CDirection {
	panic("TODO")
}

func (r *RoomPosition) GetRangeTo(target IRoomPosition) int {
	panic("TODO")
}

func (r *RoomPosition) InRangeTo(target IRoomPosition, maxRange int) bool {
	panic("TODO")
}

func (r *RoomPosition) IsEqualTo(target IRoomPosition) bool {
	panic("TODO")
}

func (r *RoomPosition) IsNearTo(target IRoomPosition) bool {
	return jsCall(r.ref, "isNearTo", target.iRef()).Bool()
}

func (r *RoomPosition) Look() []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) LookFor(lType CLook) []IRoomPosition {
	panic("TODO")
}

func createConstructionSite(src IRoomPosition, sType CStructure, name string) error {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := jsCall(src.iRef(), "createConstructionSite", string(sType), jsName).Int()
	return returnErr(result)
}

func createFlag(src IRoomPosition, name string, primary CColor, secondary CColor) (string, error) {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := jsCall(src.iRef(), "createFlag", jsName, int(primary), int(secondary))
	if result.Type() == js.TypeString {
		return result.String(), nil
	}
	return name, returnErr(result.Int())
}
