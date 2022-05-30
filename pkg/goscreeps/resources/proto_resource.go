package resources

import "syscall/js"

type Resource struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room
	id      string

	amount       int
	resourceType CResource
}

func (r *Resource) iRef() js.Value {
	return r.ref
}

func (r *Resource) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Deposit{
		ref:    ref,
		cached: make(map[string]bool),
	}
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
	if !r.cached["pos"] {
		r.pos = pos(r.ref)
		r.cached["pos"] = true
	}
	return r.pos
}

func (r *Resource) Effects() []Effect {
	if !r.cached["effects"] {
		r.effects = effects(r.ref)
		r.cached["effects"] = true
	}
	return r.effects
}

func (r *Resource) Room() *Room {
	if !r.cached["room"] {
		r.room = (&Room{}).deRef(r.ref).(*Room)
		r.cached["room"] = true
	}
	return r.room
}

func (r *Resource) Amount() int {
	if !r.cached["amount"] {
		r.amount = jsGet(r.ref, "amount").Int()
		r.cached["amount"] = true
	}
	return r.amount
}

func (r *Resource) ResourceType() CResource {
	if !r.cached["resourceType"] {
		r.resourceType = CResource(jsGet(r.ref, "resourceType").String())
		r.cached["resourceType"] = true
	}
	return r.resourceType
}

func (r *Resource) Id() string {
	if !r.cached["id"] {
		r.id = jsGet(r.ref, "id").String()
		r.cached["id"] = true
	}
	return r.id
}
