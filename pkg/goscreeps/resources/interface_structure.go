package resources

type IStructure interface {
	IRoomObject

	Hits() int
	HitsMax() int
	Id() string
	StructureType() StructureConst

	Destroy() ScreepsError
	IsActive() bool
	NotifyWhenAttacked(enabled bool) ScreepsError
}
