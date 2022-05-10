package screeps

import "syscall/js"

type FindClosestByPathAlgorithm string

const (
	ALGORITHM_ASTAR    FindClosestByPathAlgorithm = "astar"
	ALGORITHM_DIJKSTRA FindClosestByPathAlgorithm = "dijkstra"
)





func (pos RoomPosition) CreateConstructionSite(structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	result := pos.ref.Call("createConstructionSite", string(structureType), jsName).Int()
	return ErrorCode(result)
}

func (pos RoomPosition) CreateFlag(name *string, color *ColorConstant, secondaryColor *ColorConstant) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}

	var jsColor js.Value
	if color == nil {
		jsColor = js.Undefined()
	} else {
		jsColor = js.ValueOf(int(*color))
	}

	var jsSecondaryColor js.Value
	if secondaryColor == nil {
		jsSecondaryColor = js.Undefined()
	} else {
		jsSecondaryColor = js.ValueOf(int(*secondaryColor))
	}

	result := pos.ref.Call("createFlag", jsName, jsColor, jsSecondaryColor).Int()
	return ErrorCode(result)
}

