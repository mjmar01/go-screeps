package resources

type IRoomObject interface {
	IRoomPosition

	Pos() *RoomPosition
	Effects() []Effect
	Room() *Room
}

type Effect struct {
	Effect         EffectTypeConst
	Level          int
	TicksRemaining int
}
