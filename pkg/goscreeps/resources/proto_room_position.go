package resources

import "syscall/js"

type RoomPosition struct {
	ref    js.Value
	cached map[string]bool

	x        int
	y        int
	roomName string
}

var roomPositionConstructor = js.Global().Get("RoomPosition")

func deRefRoomPosition(ref js.Value) *RoomPosition {
	return &RoomPosition{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func NewRoomPosition(x, y int, roomName string) *RoomPosition {
	ref := roomPositionConstructor.New(x, y, roomName)
	return &RoomPosition{
		ref:      ref,
		cached:   map[string]bool{"x": true, "y": true, "roomName": true},
		x:        x,
		y:        y,
		roomName: roomName,
	}
}

func (r *RoomPosition) iRef() js.Value {
	return r.ref
}

func (r *RoomPosition) X() int {
	if !r.cached["x"] {
		r.x = r.ref.Get("x").Int()
		r.cached["x"] = true
	}
	return r.x
}

func (r *RoomPosition) Y() int {
	if !r.cached["y"] {
		r.y = r.ref.Get("y").Int()
		r.cached["y"] = true
	}
	return r.y
}

func (r *RoomPosition) RoomName() string {
	if !r.cached["roomName"] {
		r.roomName = r.ref.Get("roomName").String()
		r.cached["roomName"] = true
	}
	return r.roomName
}

func (r *RoomPosition) CreateConstructionSite(sType StructureType, name string) ScreepsError {
	return createConstructionSite(r, sType, name)
}

func (r *RoomPosition) CreateFlag(name string, primary Color, secondary Color) (string, ScreepsError) {
	return createFlag(r, name, primary, secondary)
}

func (r *RoomPosition) FindClosestTypeByPath(fType FindType, opts *FindClosestByPathOpts) IRoomPosition {
	return findClosestTypeByPath(r, fType, opts)
}

func (r *RoomPosition) FindClosestPosByPath(targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	return findClosestPosByPath(r, targets, opts)
}

func (r *RoomPosition) FindClosestTypeByRange(fType FindType, opts *FindFilterOpts) IRoomPosition {
	return findClosestTypeByRange(r, fType, opts)
}

func (r *RoomPosition) FindClosestPosByRange(targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	return findClosestPosByRange(r, targets, opts)
}

func (r *RoomPosition) FindTypeInRange(fType FindType, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	return findTypeInRange(r, fType, maxRange, opts)
}

func (r *RoomPosition) FindPosInRange(targets []IRoomPosition, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	return findPosInRange(r, targets, maxRange, opts)
}

func (r *RoomPosition) FindPathToCoords(x, y int, opts *FindPathOpts) Path {
	return findPathTo(r, NewRoomPosition(x, y, r.RoomName()), opts)
}

func (r *RoomPosition) FindPathToTarget(target IRoomPosition, opts *FindPathOpts) Path {
	return findPathTo(r, target, opts)
}

func (r *RoomPosition) GetDirectionToCoords(x, y int) DirectionType {
	return getDirection(r, NewRoomPosition(x, y, r.RoomName()))
}

func (r *RoomPosition) GetDirectionToTarget(target IRoomPosition) DirectionType {
	return getDirection(r, target)
}

func (r *RoomPosition) GetRangeToCoords(x, y int) float64 {
	return getRangeTo(r, NewRoomPosition(x, y, r.RoomName()))
}

func (r *RoomPosition) GetRangeToTarget(target IRoomPosition) float64 {
	return getRangeTo(r, target)
}

func (r *RoomPosition) InRangeToCoords(x, y int, maxRange float64) bool {
	return inRangeTo(r, NewRoomPosition(x, y, r.RoomName()), maxRange)
}

func (r *RoomPosition) InRangeToTarget(target IRoomPosition, maxRange float64) bool {
	return inRangeTo(r, target, maxRange)
}

func (r *RoomPosition) IsEqualToCoords(x, y int) bool {
	return isEqualTo(r, NewRoomPosition(x, y, r.RoomName()))
}

func (r *RoomPosition) IsEqualToTarget(target IRoomPosition) bool {
	return isEqualTo(r, target)
}

func (r *RoomPosition) IsNearToCoords(x, y int) bool {
	return isNearTo(r, NewRoomPosition(x, y, r.RoomName()))
}

func (r *RoomPosition) IsNearToTarget(target IRoomPosition) bool {
	return isNearTo(r, target)
}

func (r *RoomPosition) Look() []IRoomObject {
	return look(r)
}

func (r *RoomPosition) LookFor(lType LookType) []IRoomObject {
	return lookFor(r, lType)
}
