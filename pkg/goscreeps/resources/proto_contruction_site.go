package resources

import "syscall/js"

type ConstructionSite struct {
	ref    js.Value
	cached map[string]bool

	effects []Effect
	pos     *RoomPosition
	room    *Room
	id      string

	my    bool
	owner string

	progress      int
	progressTotal int
	structureType StructureConst
}

func (c *ConstructionSite) iRef() js.Value {
	return c.ref
}

func (c *ConstructionSite) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &ConstructionSite{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (c *ConstructionSite) x() int {
	return c.Pos().x()
}

func (c *ConstructionSite) y() int {
	return c.Pos().y()
}

func (c *ConstructionSite) roomName() string {
	return c.Pos().roomName()
}

func (c *ConstructionSite) Pos() *RoomPosition {
	if !c.cached["pos"] {
		c.pos = pos(c.ref)
		c.cached["pos"] = true
	}
	return c.pos
}

func (c *ConstructionSite) Effects() []Effect {
	if !c.cached["effects"] {
		c.effects = effects(c.ref)
		c.cached["effects"] = true
	}
	return c.effects
}

func (c *ConstructionSite) Room() *Room {
	if !c.cached["room"] {
		c.room = (&Room{}).deRef(c.ref).(*Room)
		c.cached["room"] = true
	}
	return c.room
}

func (c *ConstructionSite) My() bool {
	if !c.cached["my"] {
		c.my = jsGet(c.ref, "my").Bool()
		c.cached["my"] = true
	}
	return c.my
}

func (c *ConstructionSite) Owner() string {
	if !c.cached["owner"] {
		c.owner = jsGet(c.ref, "owner").String()
		c.cached["owner"] = true
	}
	return c.owner
}

func (c *ConstructionSite) Id() string {
	if !c.cached["id"] {
		c.id = jsGet(c.ref, "id").String()
		c.cached["id"] = true
	}
	return c.id
}

func (c *ConstructionSite) Progress() int {
	if !c.cached["progress"] {
		c.progress = jsGet(c.ref, "progress").Int()
		c.cached["progress"] = true
	}
	return c.progress
}
func (c *ConstructionSite) ProgressTotal() int {
	if !c.cached["progressTotal"] {
		c.progressTotal = jsGet(c.ref, "progressTotal").Int()
		c.cached["progressTotal"] = true
	}
	return c.progressTotal
}

func (c *ConstructionSite) StructureType() StructureConst {
	if !c.cached["structureType"] {
		c.structureType = StructureConst(jsGet(c.ref, "structureType").String())
		c.cached["structureType"] = true
	}
	return c.structureType
}
