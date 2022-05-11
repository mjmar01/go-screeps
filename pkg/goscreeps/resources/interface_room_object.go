package resources

import "syscall/js"

type IRoomObject interface {
	iRef() js.Value

	Pos() *RoomPosition
	Effects() []Effect
	Room() *Room
}

type Effect struct {
	Effect         EffectType
	Level          int
	TicksRemaining int
}
