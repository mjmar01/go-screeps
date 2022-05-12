package resources

import "syscall/js"

type Room struct {
	ref    js.Value
	cached map[string]bool

	// TODO *Controller
	energyAvailable         int
	energyCapacityAvailable int
	name                    string
	// TODO *StructureStorage
	// TODO *StructureTerminal
	// TODO *RoomVisual
}

func deRefRoom(ref js.Value) *Room {
	if ref.IsNull() {
		return nil
	}
	return &Room{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

// TODO Controller()

func (r *Room) EnergyAvailable() int {
	if !r.cached["energyAvailable"] {
		r.energyAvailable = r.ref.Get("energyAvailable").Int()
		r.cached["energyAvailable"] = true
	}
	return r.energyAvailable
}

func (r *Room) EnergyCapacityAvailable() int {
	if !r.cached["energyCapacityAvailable"] {
		r.energyCapacityAvailable = r.ref.Get("energyCapacityAvailable").Int()
		r.cached["energyCapacityAvailable"] = true
	}
	return r.energyCapacityAvailable
}

func (r *Room) Name() string {
	if !r.cached["name"] {
		r.name = r.ref.Get("name").String()
		r.cached["name"] = true
	}
	return r.name
}

// TODO Storage()
// TODO Terminal()
// TODO Visual()

func (r *Room) SerializePath(path Path) string {
	packedPath := packPath(path)
	return r.ref.Call("serializePath", packedPath).String()
}

func (r *Room) DeserializePath(path string) Path {
	deserializedPath := r.ref.Call("deserializePath", path)
	return unpackPath(deserializedPath)
}

func (r *Room) CreateConstructionSite(x, y int, sType StructureConst, name string) ScreepsError {
	return createConstructionSite(r.GetPositionAt(x, y), sType, name)
}

func (r *Room) CreateFlag(x, y int, name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
	return createFlag(r.GetPositionAt(x, y), name, primary, secondary)
}

func (r *Room) FindObject(fType FindObjectConst, opts *FindFilterOpts) []IRoomObject {
	// TODO Filter
	foundPositions := r.ref.Call("find", int(fType))
	foundPositionsCount := foundPositions.Length()
	result := make([]IRoomObject, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
		result[i] = deRefRoomObject(foundPositions.Index(i))
	}
	return result
}

func (r *Room) FindPos(fType FindPosConst, opts *FindFilterOpts) []IRoomPosition {
	// TODO Filter
	foundPositions := r.ref.Call("find", int(fType))
	foundPositionsCount := foundPositions.Length()
	result := make([]IRoomPosition, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
		result[i] = deRefRoomPosition(foundPositions.Index(i))
	}
	return result
}

func (r *Room) FindExitTo(roomName string) FindPosConst {
	return FindPosConst(r.ref.Call("findExitTo", roomName).Int())
}

func (r *Room) FindPath(from, to IRoomPosition, opts *FindPathOpts) Path {
	jsOpts := packFindPathOpts(opts)
	path := r.ref.Call("findPath", from.iRef(), to.iRef(), jsOpts)
	return unpackPath(path)
}

func (r *Room) GetEventLog() string {
	return r.ref.Call("getEventLog", true).String()
}

func (r *Room) GetPositionAt(x, y int) IRoomPosition {
	ref := roomPositionConstructor.New(x, y, r.Name())
	return &RoomPosition{
		ref:       ref,
		cached:    map[string]bool{"pX": true, "pY": true, "pRoomName": true},
		pX:        x,
		pY:        y,
		pRoomName: r.Name(),
	}
}

func (r *Room) GetTerrain() *Terrain {
	return deRefTerrain(r.ref.Call("getTerrain"))
}

// TODO Look*

func unpackPath(path js.Value) Path {
	// TODO
	return nil
}

func packPath(path Path) js.Value {
	// TODO
	return js.Null()
}

func packFindPathOpts(opts *FindPathOpts) js.Value {
	// TODO
	return js.Null()
}
