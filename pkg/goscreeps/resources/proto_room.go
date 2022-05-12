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

var anyRoom = js.Global().Get("Room")

func deRefRoom(ref js.Value) *Room {
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

func SerializePath(path Path) string {
	packedPath := packFindPathResult(path)
	return anyRoom.Call("serializePath", packedPath).String()
}

func DeserializePath(path string) Path {
	deserializedPath := anyRoom.Call("deserializePath", path)
	return unpackFindPathResult(deserializedPath)
}

func (r *Room) CreateConstructionSiteAtCoords(x, y int, sType StructureType, name string) ScreepsError {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := r.ref.Call("createConstructionSite", x, y, string(sType), jsName).Int()
	return ReturnErr(ErrorCode(result))
}

func (r *Room) CreateConstructionSiteAtTarget(pos IRoomPosition, sType StructureType, name string) ScreepsError {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := r.ref.Call("createConstructionSite", pos.iRef(), string(sType), jsName).Int()
	return ReturnErr(ErrorCode(result))
}

func (r *Room) CreateFlagAtTarget(pos IRoomPosition, name string, color Color, secondaryColor Color) (string, ScreepsError) {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}

	jsColor := js.ValueOf(int(color))
	jsSecondaryColor := js.ValueOf(int(secondaryColor))

	result := r.ref.Call("createFlag", pos.iRef(), jsName, jsColor, jsSecondaryColor)
	switch result.Type() {
	case js.TypeNumber:
		return "", ReturnErr(ErrorCode(result.Int()))
	}
	return result.String(), nil
}

func (r *Room) CreateFlagAtCoords(x, y int, name string, color Color, secondaryColor Color) (string, ScreepsError) {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	jsColor := js.ValueOf(int(color))
	jsSecondaryColor := js.ValueOf(int(secondaryColor))

	result := r.ref.Call("createFlag", x, y, jsName, jsColor, jsSecondaryColor)
	switch result.Type() {
	case js.TypeNumber:
		return "", ReturnErr(ErrorCode(result.Int()))
	}

	return result.String(), nil
}

func (r *Room) Find(fType FindType, opts *FindFilterOpts) []IRoomObject {
	// TODO Filter
	foundPositions := r.ref.Call("find", int(fType))
	foundPositionsCount := foundPositions.Length()
	result := make([]IRoomObject, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
		result[i] = deRefRoomObject(foundPositions.Index(i))
	}
	return result
}
func unpackFindPathResult(path js.Value) Path {
	// TODO
	return nil
}

func packFindPathResult(path Path) js.Value {
	// TODO
	return js.Null()
}
