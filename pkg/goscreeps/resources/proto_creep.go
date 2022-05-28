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

type CreepBody []BodyPart

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

func (c *Creep) Body() CreepBody {
	if !c.cached["body"] {
		jsBody := jsGet(c.ref, "body")
		partCount := jsBody.Length()
		result := make(CreepBody, partCount)
		for i := 0; i < partCount; i++ {
			part := jsBody.Index(i)
			result[i] = BodyPart(part.String())
		}
		c.body = result
		c.cached["body"] = true
	}
	return c.body
}

func (c *Creep) Attack(target IDamageable) ScreepsError {
	result := jsCall(c.ref, "attack", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) AttackController(target *Controller) ScreepsError {
	result := jsCall(c.ref, "attackController", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Build(target ConstructionSite) ScreepsError {
	result := jsCall(c.ref, "build", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) CancelOrder(methodName string) ScreepsError {
	result := jsCall(c.ref, "cancelOrder", methodName).Int()
	return ReturnErr(result)
}

func (c *Creep) ClaimController(target *Controller) ScreepsError {
	result := jsCall(c.ref, "claimController", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Dismantle(target IStructure) ScreepsError {
	result := jsCall(c.ref, "dismantle", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Drop(resource ResourceConst, amount int) ScreepsError {
	jsAmount := js.ValueOf(amount)
	if amount == -1 {
		jsAmount = js.Undefined()
	}
	result := jsCall(c.ref, "drop", string(resource), jsAmount).Int()
	return ReturnErr(result)
}

func (c *Creep) GenerateSafeMode(target *Controller) ScreepsError {
	result := jsCall(c.ref, "generateSafeMode", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) GetActiveBodyParts(partType BodyPart) int {
	return jsCall(c.ref, "getActiveBodyparts", string(partType)).Int()
}

func (c *Creep) Harvest(target IRoomPosition) ScreepsError {
	result := jsCall(c.ref, "harvest", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Heal(target Creep) ScreepsError {
	result := jsCall(c.ref, "heal", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Move(direction DirectionConst) ScreepsError {
	result := jsCall(c.ref, "move", int(direction)).Int()
	return ReturnErr(result)
}

func (c *Creep) MoveTo(target IRoomPosition) ScreepsError { // TODO opts
	result := jsCall(c.ref, "moveTo", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) MoveByPath(path Path) ScreepsError {
	result := jsCall(c.ref, "moveByPath", packPath(path)).Int()
	return ReturnErr(result)
}

func (c *Creep) NotifyWhenAttacked(enabled bool) ScreepsError {
	result := jsCall(c.ref, "notifyWhenAttacked", enabled).Int()
	return ReturnErr(result)
}

func (c *Creep) Pickup(target IRoomObject) ScreepsError {
	result := jsCall(c.ref, "pickup", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Pull(target Creep) ScreepsError {
	result := jsCall(c.ref, "pull", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) RangedAttack(target IDamageable) ScreepsError {
	result := jsCall(c.ref, "rangedAttack", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) RangedHeal(target Creep) ScreepsError {
	result := jsCall(c.ref, "rangedHeal", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) RangedMassAttack() ScreepsError {
	result := jsCall(c.ref, "rangedMassAttack").Int()
	return ReturnErr(result)
}

func (c *Creep) Repair(target IStructure) ScreepsError {
	result := jsCall(c.ref, "repair", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) ReserveController(target *Controller) ScreepsError {
	result := jsCall(c.ref, "reserveController", target.iRef()).Int()
	return ReturnErr(result)
}

func (c *Creep) Say(message string, public bool) ScreepsError {
	result := jsCall(c.ref, "say", message, public).Int()
	return ReturnErr(result)
}

func (c *Creep) SignController(target *Controller, text string) ScreepsError {
	result := jsCall(c.ref, "signController", target.iRef(), text).Int()
	return ReturnErr(result)
}

func (c *Creep) Suicide() ScreepsError {
	result := jsCall(c.ref, "suicide").Int()
	return ReturnErr(result)
}

func (c *Creep) Transfer(target IRoomObject, resource ResourceConst, amount int) ScreepsError {
	jsAmount := js.ValueOf(amount)
	if amount == -1 {
		jsAmount = js.Undefined()
	}
	result := jsCall(c.ref, "transfer", target.iRef(), string(resource), jsAmount).Int()
	return ReturnErr(result)
}

func (c *Creep) UpgradeController(target *Controller) ScreepsError {
	result := jsCall(c.ref, "upgradeController", target.iRef()).Int()
	return ReturnErr(result)
}

// TODO functions
