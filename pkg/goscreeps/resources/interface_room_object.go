package resources

import "syscall/js"

type IRoomObject interface {
	IRoomPosition

	Effects() []Effect
	Room() *Room
}

type Effect struct {
	Effect         CEffect
	Level          int
	TicksRemaining int
}

func getPos(ref js.Value, property string) interface{} {
	posRef := ref.Get(property)
	return (&RoomPosition{}).deRef(posRef)
}

func getEffects(ref js.Value, property string) interface{} {
	jsEffects := ref.Get(property)
	effectCount := jsEffects.Length()
	result := make([]Effect, effectCount)
	for i := 0; i < effectCount; i++ {
		effect := jsEffects.Index(i)
		result[i] = Effect{
			Effect:         CEffect(effect.Get("effect").Int()),
			TicksRemaining: effect.Get("ticksRemaining").Int(),
		}
		level := effect.Get("level")
		if !level.IsUndefined() {
			result[i].Level = level.Int()
		}
	}
	return result
}

func getRoom(ref js.Value, property string) interface{} {
	roomRef := ref.Get(property)
	return (&Room{}).deRef(roomRef)
}
