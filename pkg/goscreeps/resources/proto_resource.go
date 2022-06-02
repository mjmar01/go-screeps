package resources

import "syscall/js"

type Resource struct {
	ref    js.Value
	cached map[string]interface{}
}

func (r *Resource) iRef() js.Value {
	return r.ref
}

func (r *Resource) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Resource{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (r *Resource) iCache() map[string]interface{} {
	return r.cached
}

func (r *Resource) x() int {
	return r.Pos().x()
}

func (r *Resource) y() int {
	return r.Pos().x()
}

func (r *Resource) roomName() string {
	return r.Pos().roomName()
}

func (r *Resource) Pos() *RoomPosition {
	return jsGet(r, "pos", getPos).(*RoomPosition)
}

func (r *Resource) Effects() []Effect {
	return jsGet(r, "effects", getEffects).([]Effect)
}

func (r *Resource) Room() *Room {
	return jsGet(r, "room", getRoom).(*Room)
}

func (r *Resource) Amount() int {
	return jsGet(r, "amount", getInt).(int)
}

func (r *Resource) ResourceType() CResource {
	return CResource(jsGet(r, "resourceType", getString).(string))
}

func (r *Resource) Id() string {
	return jsGet(r, "id", getString).(string)
}
