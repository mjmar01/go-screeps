package resources

import "syscall/js"

type Nuke struct {
	ref    js.Value
	cached map[string]interface{}
}

func (n *Nuke) iRef() js.Value {
	return n.ref
}

func (n *Nuke) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Nuke{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (n *Nuke) iCache() map[string]interface{} {
	return n.cached
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
	return jsGet(n, "pos", getPos).(*RoomPosition)
}

func (n *Nuke) Effects() []Effect {
	return jsGet(n, "effects", getEffects).([]Effect)
}

func (n *Nuke) Room() *Room {
	return jsGet(n, "room", getRoom).(*Room)
}

func (n *Nuke) Id() string {
	return jsGet(n, "id", getString).(string)
}

func (n *Nuke) LaunchRoomName() string {
	return jsGet(n, "launchRoomName", getString).(string)
}

func (n *Nuke) TimeToLand() int {
	return jsGet(n, "timeToLand", getInt).(int)
}
