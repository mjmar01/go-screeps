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

func (r *Room) iRef() js.Value {
	return r.ref
}

func (r *Room) deRef(ref js.Value) IReference {
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
		r.energyAvailable = jsGet(r.ref, "energyAvailable").Int()
		r.cached["energyAvailable"] = true
	}
	return r.energyAvailable
}

func (r *Room) EnergyCapacityAvailable() int {
	if !r.cached["energyCapacityAvailable"] {
		r.energyCapacityAvailable = jsGet(r.ref, "energyCapacityAvailable").Int()
		r.cached["energyCapacityAvailable"] = true
	}
	return r.energyCapacityAvailable
}

func (r *Room) Name() string {
	if !r.cached["name"] {
		r.name = jsGet(r.ref, "name").String()
		r.cached["name"] = true
	}
	return r.name
}

// TODO Storage()
// TODO Terminal()
// TODO Visual()

func (r *Room) SerializePath(path Path) string {
	packedPath := packPath(path)
	return jsCall(r.ref, "serializePath", packedPath).String()
}

func (r *Room) DeserializePath(path string) Path {
	deserializedPath := jsCall(r.ref, "deserializePath", path)
	return unpackPath(deserializedPath)
}

func (r *Room) CreateConstructionSite(x, y int, sType StructureConst, name string) ScreepsError {
	return createConstructionSite(r.GetPositionAt(x, y), sType, name)
}

func (r *Room) CreateFlag(x, y int, name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
	return createFlag(r.GetPositionAt(x, y), name, primary, secondary)
}

func (r *Room) Find(fType FindConst, opts *FindFilterOpts) []IRoomPosition {
	// TODO Filter
	foundPositions := jsCall(r.ref, "find", int(fType))
	foundPositionsCount := foundPositions.Length()
	result := make([]IRoomPosition, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
		ref := foundPositions.Index(i)
		result[i] = deRefUnknown(ref).(IRoomPosition)
	}
	return result
}

func (r *Room) FindExitTo(roomName string) FindConst {
	return FindConst(jsCall(r.ref, "findExitTo", roomName).Int())
}

func (r *Room) FindPath(from, to IRoomPosition, opts *FindPathOpts) Path {
	jsOpts := packFindPathOpts(opts)
	path := jsCall(r.ref, "findPath", from.iRef(), to.iRef(), jsOpts)
	return unpackPath(path)
}

func (r *Room) GetEventLog() string {
	return jsCall(r.ref, "getEventLog", true).String()
}

func (r *Room) GetPositionAt(x, y int) *RoomPosition {
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
	return (&Terrain{}).deRef(jsCall(r.ref, "getTerrain")).(*Terrain)
}

// TODO Look*

func unpackPath(path js.Value) Path {
	pathLength := path.Length()
	result := make(Path, pathLength)
	for i := 0; i < pathLength; i++ {
		step := path.Index(i)
		result[i] = PathStep{
			x:         jsGet(step, "x").Int(),
			y:         jsGet(step, "y").Int(),
			dx:        jsGet(step, "dx").Int(),
			dy:        jsGet(step, "dy").Int(),
			direction: DirectionConst(jsGet(step, "direction").Int()),
		}
	}
	return result
}

func packPath(path Path) js.Value {
	length := len(path)
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		step := map[string]interface{}{}
		step["x"] = path[i].x
		step["y"] = path[i].y
		step["dx"] = path[i].dx
		step["dy"] = path[i].dy
		step["direction"] = int(path[i].direction)
		result[i] = step
	}
	return js.ValueOf(result)
}

func packFindPathOpts(opts *FindPathOpts) js.Value {
	if opts == nil {
		return js.Undefined()
	}
	result := make(map[string]interface{}, 10)
	result["ignoreCreeps"] = opts.IgnoreCreeps
	result["ignoreDestructibleStructures"] = opts.IgnoreDestructibleStructures
	result["ignoreRoads"] = opts.IgnoreRoads
	result["costCallback"] = js.Undefined() // TODO
	if opts.MaxOps == 0 {
		result["maxOps"] = 2000
	} else {
		result["maxOps"] = opts.MaxOps
	}
	if opts.HeuristicWeight == 0 {
		result["heuristicWeight"] = 1.2
	}
	result["serialize"] = false
	if opts.MaxRooms == 0 {
		result["maxRooms"] = 16
	} else {
		result["maxRooms"] = opts.MaxRooms
	}
	result["range"] = opts.Range
	if opts.PlainCost == 0 {
		result["plainCost"] = 1
	} else {
		result["plainCost"] = opts.PlainCost
	}
	if opts.SwampCost == 0 {
		result["swampCost"] = 1
	} else {
		result["swampCost"] = opts.SwampCost
	}
	return js.ValueOf(result)
}
