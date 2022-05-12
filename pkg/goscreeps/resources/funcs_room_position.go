package resources

// TODO implement
func createConstructionSite(src IRoomPosition, sType StructureConst, name string) ScreepsError {
	panic("Implement me")
}

func createFlag(src IRoomPosition, name string, primary ColorConst, secondary ColorConst) (string, ScreepsError) {
	panic("Implement me")
}

func findClosestTypeByPath(src IRoomPosition, fType FindObjectConst, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByPath(src IRoomPosition, targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestTypeByRange(src IRoomPosition, fType FindObjectConst, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByRange(src IRoomPosition, targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findTypeInRange(src IRoomPosition, fType FindObjectConst, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
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

func getRangeTo(src IRoomPosition, target IRoomPosition) float64 {
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

func look(src IRoomPosition) []IRoomObject {
	panic("Implement me")
}

func lookFor(src IRoomPosition, lType LookConst) []IRoomObject {
	panic("Implement me")
}
