package resources

import "syscall/js"

type RoomObject interface {
	ref() js.Value

	Pos() Positionable
	Effects() []Effect
	// TODO Room() *Room
}

type Effect struct {
	Effect EffectType
	Level int
	TicksRemaining int
}