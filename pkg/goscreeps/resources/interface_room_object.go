package resources

import "syscall/js"

type IRoomObject interface {
	iRef() js.Value

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
