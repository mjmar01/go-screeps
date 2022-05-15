package resources

import (
	"regexp"
	"strings"
	"syscall/js"
)

type RoomPosition struct {
	ref    js.Value
	cached map[string]bool

	pX        int
	pY        int
	pRoomName string
}

var roomPositionConstructor = js.Global().Get("RoomPosition")
var re = regexp.MustCompile("[a-zA-Z0-9\\-\\_\\#\\]]+?\\s")

func (r *RoomPosition) iRef() js.Value {
	return r.ref
}

func (r *RoomPosition) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &RoomPosition{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (r *RoomPosition) x() int {
	if !r.cached["pX"] {
		r.pX = jsGet(r.ref, "x").Int()
		r.cached["pX"] = true
	}
	return r.pX
}

func (r *RoomPosition) y() int {
	if !r.cached["pY"] {
		r.pY = jsGet(r.ref, "y").Int()
		r.cached["pY"] = true
	}
	return r.pY
}

func (r *RoomPosition) roomName() string {
	if !r.cached["pRoomName"] {
		r.pRoomName = jsGet(r.ref, "roomName").String()
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
	panic("TODO")
}

func (r *RoomPosition) FindClosestPosByPath(targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindClosestTypeByRange(fType FindConst, opts *FindFilterOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindClosestPosByRange(targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindTypeInRange(fType FindConst, maxRange int, opts *FindFilterOpts) []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindPosInRange(targets []IRoomPosition, maxRange int, opts *FindFilterOpts) []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) FindPathTo(target IRoomPosition, opts *FindPathOpts) Path {
	panic("TODO")
}

func (r *RoomPosition) GetDirectionTo(target IRoomPosition) DirectionConst {
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
	panic("TODO")
}

func (r *RoomPosition) Look() []IRoomPosition {
	panic("TODO")
}

func (r *RoomPosition) LookFor(lType LookConst) []IRoomPosition {
	panic("TODO")
}

func createConstructionSite(src IRoomPosition, sType StructureConst, name string) ScreepsError {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := jsCall(src.iRef(), "createConstructionSite", string(sType), jsName).Int()
	return ReturnErr(result)
}

func createFlag(src IRoomPosition, name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
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
	return name, ReturnErr(result.Int())
}

func getRoomPosRefType(ref js.Value) IRoomPosition {
	typeStr := jsCall(ref, "toString").String()
	matches := re.FindAllString(typeStr, -1)
	if matches == nil {
		return &RoomPosition{}
	}
	typeStr = matches[len(matches)-1]
	typeStr = strings.TrimSpace(typeStr)
	var result IRoomPosition
	switch typeStr {
	case "pos":
		result = &RoomPosition{}
	case "spawn":
		result = &StructureSpawn{}
	default:
		panic("Unknown Type: \"" + typeStr + "\"")
	}
	return result.deRef(ref).(IRoomPosition)
}
