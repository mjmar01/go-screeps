package resources

import "syscall/js"

type Deposit struct {
	ref    js.Value
	cached map[string]interface{}
}

func (d *Deposit) iRef() js.Value {
	return d.ref
}

func (d *Deposit) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Deposit{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (d *Deposit) iCache() map[string]interface{} {
	return d.cached
}

func (d *Deposit) x() int {
	return d.Pos().x()
}

func (d *Deposit) y() int {
	return d.Pos().x()
}

func (d *Deposit) roomName() string {
	return d.Pos().roomName()
}

func (d *Deposit) Pos() *RoomPosition {
	return jsGet(d, "pos", getPos).(*RoomPosition)
}

func (d *Deposit) Effects() []Effect {
	return jsGet(d, "effects", getEffects).([]Effect)
}

func (d *Deposit) Room() *Room {
	return jsGet(d, "room", getRoom).(*Room)
}

func (d *Deposit) Cooldown() int {
	return jsGet(d, "cooldown", getInt).(int)
}

func (d *Deposit) LastCooldown() int {
	return jsGet(d, "lastCooldown", getInt).(int)
}

func (d *Deposit) TicksToDecay() int {
	return jsGet(d, "ticksToDecay", getInt).(int)
}

func (d *Deposit) Id() string {
	return jsGet(d, "id", getString).(string)
}

func (d *Deposit) DepositType() CResource {
	return CResource(jsGet(d, "depositType", getString).(string))
}
