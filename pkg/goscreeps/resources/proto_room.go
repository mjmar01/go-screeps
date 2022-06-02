package resources

import "syscall/js"

type Room struct {
	ref    js.Value
	cached map[string]interface{}
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
		cached: make(map[string]interface{}),
	}
}

func (r *Room) iCache() map[string]interface{} {
	return r.cached
}

func (r *Room) Controller() *Controller {
	return jsGet(r, "controller", getController).(*Controller)
}

func (r *Room) EnergyAvailable() int {
	return jsGet(r, "energyAvailable", getInt).(int)
}

func (r *Room) EnergyCapacityAvailable() int {
	return jsGet(r, "energyCapacityAvailable", getInt).(int)
}

func (r *Room) Name() string {
	return jsGet(r, "name", getString).(string)
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

func (r *Room) CreateConstructionSite(x, y int, sType CStructure, name string) error {
	return createConstructionSite(r.GetPositionAt(x, y), sType, name)
}

func (r *Room) CreateFlag(x, y int, name string, primary CColor, secondary CColor) (string, error) {
	return createFlag(r.GetPositionAt(x, y), name, primary, secondary)
}

func (r *Room) Find(fType CFind, opts *FindFilterOpts) []IRoomPosition {
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

func (r *Room) FindExitTo(roomName string) CFind {
	return CFind(jsCall(r.ref, "findExitTo", roomName).Int())
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
		ref:    ref,
		cached: map[string]interface{}{"x": x, "y": y, "roomName": r.Name()},
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
			x:         step.Get("x").Int(),
			y:         step.Get("y").Int(),
			dx:        step.Get("dx").Int(),
			dy:        step.Get("dy").Int(),
			direction: CDirection(step.Get("direction").Int()),
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
