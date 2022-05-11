package resources

import "syscall/js"

type IRoomPosition interface {
	iRef() js.Value

	X() int
	Y() int
	RoomName() string
}

type FindPathOpts struct {
	IgnoreCreeps                 bool
	IgnoreDestructibleStructures bool
	IgnoreRoads                  bool
	// TODO CostCallback         *CostCallback
	MaxOps          uint
	HeuristicWeight float64
	Serialize       bool
	MaxRooms        uint
	Range           uint
	PlainCost       uint
	SwampCost       uint
}

type FindClosestByPathOpts struct {
	FindPathOpts
	// TODO Filter
	Algorithm AlgorithmType
}

type FindFilterOpts struct {
	// TODO Filter
}

type Path []PathStep
type PathStep struct {
	x         int
	y         int
	dx        int
	dy        int
	direction DirectionType
}

// TODO implement
func createConstructionSite(src IRoomPosition, sType StructureType, name string) ScreepsError {
	panic("Implement me")
}

func createFlag(src IRoomPosition, name string, primary Color, secondary Color) (string, ScreepsError) {
	panic("Implement me")
}

func findClosestTypeByPath(src IRoomPosition, fType FindType, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByPath(src IRoomPosition, targets []IRoomPosition, opts *FindClosestByPathOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestTypeByRange(src IRoomPosition, fType FindType, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findClosestPosByRange(src IRoomPosition, targets []IRoomPosition, opts *FindFilterOpts) IRoomPosition {
	panic("Implement me")
}

func findTypeInRange(src IRoomPosition, fType FindType, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	panic("Implement me")
}

func findPosInRange(src IRoomPosition, targets []IRoomPosition, maxRange float64, opts *FindFilterOpts) []IRoomPosition {
	panic("Implement me")
}

func findPathTo(src IRoomPosition, target IRoomPosition, opts *FindPathOpts) Path {
	panic("Implement me")
}

func getDirection(src IRoomPosition, target IRoomPosition) DirectionType {
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

func lookFor(src IRoomPosition, lType LookType) []IRoomObject {
	panic("Implement me")
}
