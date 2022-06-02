package resources

import "syscall/js"

type Flag struct {
	ref    js.Value
	cached map[string]interface{}
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
		cached: make(map[string]interface{}),
	}
}

func (f *Flag) iCache() map[string]interface{} {
	return f.cached
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
	return jsGet(f, "pos", getPos).(*RoomPosition)
}

func (f *Flag) Effects() []Effect {
	return jsGet(f, "effects", getEffects).([]Effect)
}

func (f *Flag) Room() *Room {
	return jsGet(f, "room", getRoom).(*Room)
}

func (f *Flag) Name() string {
	return jsGet(f, "name", getString).(string)
}

func (f *Flag) Color() CColor {
	return CColor(jsGet(f, "color", getInt).(int))
}

func (f *Flag) SecondaryColor() CColor {
	return CColor(jsGet(f, "secondaryColor", getInt).(int))
}

func (f *Flag) Remove() {
	jsCall(f.ref, "remove")
}

func (f *Flag) SetColor(primary, secondary CColor) error {
	result := jsCall(f.ref, "setColor", int(primary), int(secondary)).Int()
	return returnErr(result)
}

func (f *Flag) SetPos(pos IRoomPosition) error {
	result := jsCall(f.ref, "setPosition", pos.iRef()).Int()
	return returnErr(result)
}
