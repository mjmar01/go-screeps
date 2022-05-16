package resources

import "syscall/js"

type Creep struct {
	ref    js.Value
	cached map[string]bool

	pos         *RoomPosition
	effects     []Effect
	room        *Room
	body        CreepBody
	fatigue     int
	hits        int
	hitsMax     int
	id          string
	my          bool
	owner       string
	name        string
	saying      string
	spawning    bool
	store       *Store
	ticksToLive int
}

func (c *Creep) iRef() js.Value {
	return c.ref
}

func (c *Creep) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Creep{
		ref:    ref,
		cached: make(map[string]bool),
	}
}

func (c *Creep) x() int {
	return c.Pos().x()
}

func (c *Creep) y() int {
	return c.Pos().y()
}

func (c *Creep) roomName() string {
	return c.Pos().roomName()
}

func (c *Creep) Pos() *RoomPosition {
	if !c.cached["pos"] {
		c.pos = pos(c.ref)
		c.cached["pos"] = true
	}
	return c.pos
}

func (c *Creep) Effects() []Effect {
	if !c.cached["effects"] {
		c.effects = effects(c.ref)
		c.cached["effects"] = true
	}
	return c.effects
}

func (c *Creep) Room() *Room {
	if !c.cached["room"] {
		c.room = (&Room{}).deRef(c.ref).(*Room)
		c.cached["room"] = true
	}
	return c.room
}

func (c *Creep) My() bool {
	if !c.cached["my"] {
		c.my = jsGet(c.ref, "my").Bool()
		c.cached["my"] = true
	}
	return c.my
}

func (c *Creep) Owner() string {
	if !c.cached["owner"] {
		c.owner = jsGet(c.ref, "owner").String()
		c.cached["owner"] = true
	}
	return c.owner
}

func (c *Creep) Hits() int {
	if !c.cached["hits"] {
		c.hits = jsGet(c.ref, "hits").Int()
		c.cached["hits"] = true
	}
	return c.hits
}

func (c *Creep) HitsMax() int {
	if !c.cached["hitsMax"] {
		c.hitsMax = jsGet(c.ref, "hitsMax").Int()
		c.cached["hitsMax"] = true
	}
	return c.hitsMax
}

func (c *Creep) Id() string {
	if !c.cached["id"] {
		c.id = jsGet(c.ref, "id").String()
		c.cached["id"] = true
	}
	return c.id
}

func (c *Creep) Name() string {
	if !c.cached["name"] {
		c.name = jsGet(c.ref, "name").String()
		c.cached["name"] = true
	}
	return c.name
}

func (c *Creep) Saying() string {
	if !c.cached["saying"] {
		c.saying = jsGet(c.ref, "saying").String()
		c.cached["saying"] = true
	}
	return c.saying
}

func (c *Creep) Spawning() bool {
	if !c.cached["spawning"] {
		c.spawning = jsGet(c.ref, "spawning").Bool()
		c.cached["spawning"] = true
	}
	return c.spawning
}

func (c *Creep) Fatigue() int {
	if !c.cached["fatigue"] {
		c.fatigue = jsGet(c.ref, "fatigue").Int()
		c.cached["fatigue"] = true
	}
	return c.fatigue
}

func (c *Creep) TicksToLive() int {
	if !c.cached["ticksToLive"] {
		c.ticksToLive = jsGet(c.ref, "ticksToLive").Int()
		c.cached["ticksToLive"] = true
	}
	return c.ticksToLive
}

func (c *Creep) Store() *Store {
	if !c.cached["store"] {
		c.store = (&Store{}).deRef(c.ref).(*Store)
		c.cached["store"] = true
	}
	return c.store
}

// TODO Body, functions
