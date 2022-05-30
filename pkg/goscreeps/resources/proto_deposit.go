package resources

import "syscall/js"

type Deposit struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room
	id      string

	cooldown     int
	lastCooldown int
	ticksToDecay int
	depositType  CResource
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
		cached: make(map[string]bool),
	}
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
	if !d.cached["pos"] {
		d.pos = pos(d.ref)
		d.cached["pos"] = true
	}
	return d.pos
}

func (d *Deposit) Effects() []Effect {
	if !d.cached["effects"] {
		d.effects = effects(d.ref)
		d.cached["effects"] = true
	}
	return d.effects
}

func (d *Deposit) Room() *Room {
	if !d.cached["room"] {
		d.room = (&Room{}).deRef(d.ref).(*Room)
		d.cached["room"] = true
	}
	return d.room
}

func (d *Deposit) Cooldown() int {
	if !d.cached["cooldown"] {
		d.cooldown = jsGet(d.ref, "cooldown").Int()
		d.cached["cooldown"] = true
	}
	return d.cooldown
}

func (d *Deposit) LastCooldown() int {
	if !d.cached["lastCooldown"] {
		d.lastCooldown = jsGet(d.ref, "lastCooldown").Int()
		d.cached["lastCooldown"] = true
	}
	return d.lastCooldown
}

func (d *Deposit) TicksToDecay() int {
	if !d.cached["ticksToDecay"] {
		d.ticksToDecay = jsGet(d.ref, "ticksToDecay").Int()
		d.cached["ticksToDecay"] = true
	}
	return d.ticksToDecay
}

func (d *Deposit) Id() string {
	if !d.cached["id"] {
		d.id = jsGet(d.ref, "id").String()
		d.cached["id"] = true
	}
	return d.id
}

func (d *Deposit) DepositType() CResource {
	if !d.cached["depositType"] {
		d.depositType = CResource(jsGet(d.ref, "depositType").String())
		d.cached["depositType"] = true
	}
	return d.depositType
}
