package resources

import (
	"regexp"
	"syscall/js"
)

var re = regexp.MustCompile("[a-zA-Z0-9\\-\\_\\#\\]]+?\\s")

func getRoomPosRefType(ref js.Value) IRoomPosition {
	typeStr := ref.Call("toString").String()
	matches := re.FindAllString(typeStr, -1)
	if matches == nil {
		return &RoomPosition{}
	}
	typeStr = matches[len(matches)-1]
	switch typeStr {
	case "pos":
		return &RoomPosition{}
	default:
		return &RoomObject{}
	}
}

// TODO implement
func createConstructionSite(src IRoomPosition, sType StructureConst, name string) ScreepsError {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := src.iRef().Call("createConstructionSite", string(sType), jsName).Int()
	return ReturnErr(ErrorCode(result))
}

func createFlag(src IRoomPosition, name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
	var jsName js.Value
	if name == "" {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(name)
	}
	result := src.iRef().Call("createFlag", jsName, int(primary), int(secondary))
	if result.Type() == js.TypeString {
		return result.String(), nil
	}
	return name, ReturnErr(ErrorCode(result.Int()))
}

func findClosestTypeByPath(src IRoomPosition, fType FindConst, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByPath(src IRoomPosition, targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestTypeByRange(src IRoomPosition, fType FindConst, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByRange(src IRoomPosition, targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findTypeInRange(src IRoomPosition, fType FindConst, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	panic("Implement me")
}

func findPosInRange(src IRoomPosition, targets []IRoomPosition, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	panic("Implement me")
}

func findPathTo(src IRoomPosition, target IRoomPosition, opts *FindPathOpts) Path {
	panic("Implement me")
}

func getDirectionTo(src IRoomPosition, target IRoomPosition) DirectionConst {
	panic("Implement me")
}

func getRangeTo(src IRoomPosition, target IRoomPosition) int {
	panic("Implement me")
}

func inRangeTo(src IRoomPosition, target IRoomPosition, maxRange float64) bool {
	panic("Implement me")
}

func isEqualTo(src IRoomPosition, target IRoomPosition) bool {
	panic("Implement me")
}

func isNearTo(src IRoomPosition, target IRoomPosition) bool {
	panic("Implement me")
}

func look(src IRoomPosition) []IRoomPosition {
	panic("Implement me")
}

func lookFor(src IRoomPosition, lType LookConst) []IRoomPosition {
	panic("Implement me")
}
