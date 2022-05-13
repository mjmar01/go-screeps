package resources

import "syscall/js"

type RoomObject struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room
}

func (r *RoomObject) iRef() js.Value {
	return r.ref
}

func (r *RoomObject) deRef(ref js.Value) IRoomPosition {
	if ref.IsNull() {
		return nil
	}
	return &RoomObject{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (r *RoomObject) Pos() *RoomPosition {
	if !r.cached["pos"] {
		r.pos = pos(r.ref)
		r.cached["pos"] = true
	}
	return r.pos
}

func (r *RoomObject) Effects() []Effect {
	if !r.cached["effects"] {
		r.effects = effects(r.ref)
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
