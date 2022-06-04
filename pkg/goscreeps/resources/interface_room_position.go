package resources

// IRoomPosition is the interface of anything that has a position in a room. (Ex: Flags, Creeps, Structures)
type IRoomPosition interface {
	IReference

	x() int
	y() int
	roomName() string
	Pos() *RoomPosition
}

type FindPathOpts struct {
	IgnoreCreeps                 bool
	IgnoreDestructibleStructures bool
	IgnoreRoads                  bool
	// TODO CostCallback         *CostCallback
	MaxOps          uint
	HeuristicWeight float64
	MaxRooms        uint
	Range           uint
	PlainCost       uint
	SwampCost       uint
}

type FindClosestByPathOpts struct {
	FindPathOpts
	// TODO Filter
	Algorithm CAlgorithm
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
	direction CDirection
}
