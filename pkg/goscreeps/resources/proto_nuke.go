package resources

import "syscall/js"

type Nuke struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room
	id      string

	launchRoomName string
	timeToLand     int
}

func (n *Nuke) iRef() js.Value {
	return n.ref
}

func (n *Nuke) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Deposit{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (n *Nuke) x() int {
	return n.Pos().x()
}

func (n *Nuke) y() int {
	return n.Pos().x()
}

func (n *Nuke) roomName() string {
	return n.Pos().roomName()
}

func (n *Nuke) Pos() *RoomPosition {
	if !n.cached["pos"] {
		n.pos = pos(n.ref)
		n.cached["pos"] = true
	}
	return n.pos
}

func (n *Nuke) Effects() []Effect {
	if !n.cached["effects"] {
		n.effects = effects(n.ref)
		n.cached["effects"] = true
	}
	return n.effects
}

func (n *Nuke) Room() *Room {
	if !n.cached["room"] {
		n.room = (&Room{}).deRef(n.ref).(*Room)
		n.cached["room"] = true
	}
	return n.room
}

func (n *Nuke) Id() string {
	if !n.cached["id"] {
		n.id = jsGet(n.ref, "id").String()
		n.cached["id"] = true
	}
	return n.id
}

func (n *Nuke) LaunchRoomName() string {
	if !n.cached["launchRoomName"] {
		n.launchRoomName = jsCall(n.ref, "launchRoomName", true).String()
		n.cached["launchRoomName"] = true
	}
	return n.launchRoomName
}

func (n *Nuke) TimeToLand() int {
	if !n.cached["timeToLand"] {
		n.timeToLand = jsCall(n.ref, "timeToLand", true).Int()
		n.cached["timeToLand"] = true
	}
	return n.timeToLand
}
