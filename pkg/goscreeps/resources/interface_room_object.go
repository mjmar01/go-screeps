package resources

import "syscall/js"

// IRoomObject is the interface of any type that has a position inside a room and can have effects applied. (Ex: Creeps, Structures)
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

// retrieve a position for a reference. Used for jsGet
func getPos(ref js.Value, property string) interface{} {
	posRef := ref.Get(property)
	return (&RoomPosition{}).deRef(posRef)
}

// retrieve effects for a reference. Used for jsGet
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

// retrieve a room for a reference. Used for jsGet
func getRoom(ref js.Value, property string) interface{} {
	roomRef := ref.Get(property)
	return (&Room{}).deRef(roomRef)
}
