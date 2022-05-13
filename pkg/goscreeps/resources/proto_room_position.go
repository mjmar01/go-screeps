package resources

import "syscall/js"

type RoomPosition struct {
	ref    js.Value
	cached map[string]bool

	pX        int
	pY        int
	pRoomName string
}

var roomPositionConstructor = js.Global().Get("RoomPosition")

func (r *RoomPosition) deRef(ref js.Value) IRoomPosition {
	if ref.IsNull() {
		return nil
	}
	return &RoomPosition{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (r *RoomPosition) iRef() js.Value {
	return r.ref
}

func (r *RoomPosition) x() int {
	if !r.cached["pX"] {
		r.pX = r.ref.Get("x").Int()
		r.cached["pX"] = true
	}
	return r.pX
}

func (r *RoomPosition) y() int {
	if !r.cached["pY"] {
		r.pY = r.ref.Get("y").Int()
		r.cached["pY"] = true
	}
	return r.pY
}

func (r *RoomPosition) roomName() string {
	if !r.cached["pRoomName"] {
		r.pRoomName = r.ref.Get("roomName").String()
		r.cached["pRoomName"] = true
	}
	return r.pRoomName
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

func (r *RoomPosition) CreateConstructionSite(sType StructureConst, name string) ScreepsError {
	return createConstructionSite(r, sType, name)
}

func (r *RoomPosition) CreateFlag(name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
	return createFlag(r, name, primary, secondary)
}

func (r *RoomPosition) FindClosestTypeByPath(fType FindConst, opts *FindClosestByPathOpts) IRoomPosition {
	return findClosestTypeByPath(r, fType, opts)
}

func (r *RoomPosition) FindClosestPosByPath(targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	return findClosestPosByPath(r, targets, opts)
}

func (r *RoomPosition) FindClosestTypeByRange(fType FindConst, opts *FindFilterOpts) IRoomPosition {
	return findClosestTypeByRange(r, fType, opts)
}

func (r *RoomPosition) FindClosestPosByRange(targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	return findClosestPosByRange(r, targets, opts)
}

func (r *RoomPosition) FindTypeInRange(fType FindConst, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	return findTypeInRange(r, fType, maxRange, opts)
}

func (r *RoomPosition) FindPosInRange(targets []IRoomPosition, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	return findPosInRange(r, targets, maxRange, opts)
}

func (r *RoomPosition) FindPathTo(target IRoomPosition, opts *FindPathOpts) Path {
	return findPathTo(r, target, opts)
}

func (r *RoomPosition) GetDirectionTo(target IRoomPosition) DirectionConst {
	return getDirectionTo(r, target)
}

func (r *RoomPosition) GetRangeTo(target IRoomPosition) int {
	return getRangeTo(r, target)
}

func (r *RoomPosition) InRangeTo(target IRoomPosition, maxRange float64) bool {
	return inRangeTo(r, target, maxRange)
}

func (r *RoomPosition) IsEqualTo(target IRoomPosition) bool {
	return isEqualTo(r, target)
}

func (r *RoomPosition) IsNearTo(target IRoomPosition) bool {
	return isNearTo(r, target)
}

func (r *RoomPosition) Look() []IRoomPosition {
	return look(r)
}

func (r *RoomPosition) LookFor(lType LookConst) []IRoomPosition {
	return lookFor(r, lType)
}
