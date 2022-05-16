package resources

import "syscall/js"

type Flag struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room

	color          ColorConst
	secondaryColor ColorConst
	name           string
}

func (f *Flag) iRef() js.Value {
	return f.ref
}

func (f *Flag) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Flag{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (f *Flag) x() int {
	return f.Pos().x()
}

func (f *Flag) y() int {
	return f.Pos().y()
}

func (f *Flag) roomName() string {
	return f.Pos().roomName()
}

func (f *Flag) Pos() *RoomPosition {
	if !f.cached["pos"] {
		f.pos = pos(f.ref)
		f.cached["pos"] = true
	}
	return f.pos
}

func (f *Flag) Effects() []Effect {
	if !f.cached["effects"] {
		f.effects = effects(f.ref)
		f.cached["effects"] = true
	}
	return f.effects
}

func (f *Flag) Room() *Room {
	if !f.cached["room"] {
		f.room = (&Room{}).deRef(f.ref).(*Room)
		f.cached["room"] = true
	}
	return f.room
}

func (f *Flag) Name() string {
	if !f.cached["name"] {
		f.name = jsGet(f.ref, "name").String()
		f.cached["name"] = true
	}
	return f.name
}

func (f *Flag) Color() ColorConst {
	if !f.cached["color"] {
		f.color = ColorConst(jsGet(f.ref, "color").Int())
		f.cached["color"] = true
	}
	return f.color
}

func (f *Flag) SecondaryColor() ColorConst {
	if !f.cached["secondaryColor"] {
		f.secondaryColor = ColorConst(jsGet(f.ref, "secondaryColor").Int())
		f.cached["secondaryColor"] = true
	}
	return f.secondaryColor
}

func (f *Flag) Remove() {
	jsCall(f.ref, "remove")
}

func (f *Flag) SetColor(primary, secondary ColorConst) ScreepsError {
	result := jsCall(f.ref, "setColor", int(primary), int(secondary)).Int()
	return ReturnErr(result)
}

func (f *Flag) SetPos(pos IRoomPosition) ScreepsError {
	result := jsCall(f.ref, "setPosition", pos.iRef()).Int()
	return ReturnErr(result)
}
