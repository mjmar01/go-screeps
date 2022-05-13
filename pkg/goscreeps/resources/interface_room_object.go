package resources

import "syscall/js"

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

func pos(ref js.Value) *RoomPosition {
	posRef := ref.Get("pos")
	return (&RoomPosition{}).deRef(posRef).(*RoomPosition)
}

func effects(src js.Value) []Effect {
	jsEffects := src.Get("effects")
	effectCount := jsEffects.Length()
	result := make([]Effect, effectCount)
	for i := 0; i < effectCount; i++ {
		effect := jsEffects.Index(i)
		result[i] = Effect{
			Effect:         EffectTypeConst(effect.Get("effect").Int()),
			TicksRemaining: effect.Get("ticksRemaining").Int(),
		}
		level := effect.Get("level")
		if !level.IsUndefined() {
			result[i].Level = level.Int()
		}
	}
	return result
}
