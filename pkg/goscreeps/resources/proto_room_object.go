package resources

import "syscall/js"

type RoomObject struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room
}

func deRefRoomObject(ref js.Value) *RoomObject {
	if ref.IsNull() {
		return nil
	}
	return &RoomObject{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (r *RoomObject) iRef() js.Value {
	return r.ref
}

func (r *RoomObject) Pos() *RoomPosition {
	if !r.cached["pos"] {
		ref := r.ref.Get("pos")
		r.pos = deRefRoomPosition(ref)
		r.cached["pos"] = true
	}
	return r.pos
}

func (r *RoomObject) Effects() []Effect {
	if !r.cached["effects"] {
		jsEffects := r.ref.Get("effects")
		effectCount := jsEffects.Get("length").Int()
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
		r.effects = result
		r.cached["effects"] = true
	}
	return r.effects
}

func (r *RoomObject) Room() *Room {
	if !r.cached["room"] {
		r.room = deRefRoom(r.ref.Get("room"))
		r.cached["room"] = true
	}
	return r.room
}

func (r *RoomObject) x() int {
	return r.Pos().x()
}

func (r *RoomObject) y() int {
	return r.Pos().y()
}

func (r *RoomObject) roomName() string {
	return r.Pos().roomName()
}
