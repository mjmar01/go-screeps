package resources

type IRoomPosition interface {
	IReference

	x() int
	y() int
	roomName() string
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
	Algorithm AlgorithmConst
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
	direction DirectionConst
}
