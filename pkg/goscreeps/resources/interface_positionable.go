package resources

import "syscall/js"

type Positionable interface {
	ref() js.Value

	X() int
	Y() int
	RoomName() string

	createConstructionSite(StructureType, string) ScreepsError
	createFlag(string, Color, Color) ScreepsError
	findClosestTypeByPath(FindType, *FindClosestByPathOpts) Positionable
	findClosestPosByPath([]Positionable, *FindClosestByPathOpts) Positionable
	findTypeInRange(FindType, *FindFilterOpts) []Positionable
	findPosInRange([]Positionable, *FindFilterOpts) []Positionable
	findPathToCoords(int, int, *FindPathOpts) Path
	findPathToTarget(Positionable, *FindPathOpts) Path
	getDirectionToCoords(int, int) DirectionType
	getDirectionToTarget(Positionable) DirectionType
	getRangeToCoords(int, int) float64
	getRangeToTarget(Positionable) float64
	inRangeToCoords(int, int, int) bool
	inRangeToTarget(Positionable, int) bool
	isEqualToCoords(int, int) bool
	isEqualToTarget(Positionable) bool
	isNearToCoords(int, int) bool
	isNearToTarget(Positionable) bool
	look() RoomObject
	lookFor(LookType) RoomObject
}

type FindPathOpts struct {
	IgnoreCreeps                 bool
	IgnoreDestructibleStructures bool
	IgnoreRoads                  bool
	// TODO CostCallback         *CostCallback
	MaxOps                       uint
	HeuristicWeight              float64
	Serialize                    bool
	MaxRooms                     uint
	Range                        uint
	PlainCost                    uint
	SwampCost                    uint
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
	x int
	y int
	dx int
	dy int
	direction DirectionType
}