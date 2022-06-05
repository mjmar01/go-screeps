package resources

import "syscall/js"

type Creep struct {
	ref    js.Value
	cached map[string]interface{}
}

type CreepBody []CBodyPart

func (c *Creep) iRef() js.Value {
	return c.ref
}

func (c *Creep) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Creep{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (c *Creep) iCache() map[string]interface{} {
	return c.cached
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
	return jsGet(c, "pos", getPos).(*RoomPosition)
}

func (c *Creep) Effects() []Effect {
	return jsGet(c, "effects", getEffects).([]Effect)
}

func (c *Creep) Room() *Room {
	return jsGet(c, "room", getRoom).(*Room)
}

func (c *Creep) My() bool {
	return jsGet(c, "my", getBool).(bool)
}

func (c *Creep) Owner() string {
	return jsGet(c, "owner", getString).(string)
}

func (c *Creep) Hits() int {
	return jsGet(c, "hits", getInt).(int)
}

func (c *Creep) HitsMax() int {
	return jsGet(c, "hitsMax", getInt).(int)
}

func (c *Creep) Id() string {
	return jsGet(c, "id", getString).(string)
}

func (c *Creep) Name() string {
	return jsGet(c, "name", getString).(string)
}

func (c *Creep) Saying() string {
	return jsGet(c, "saying", getString).(string)
}

func (c *Creep) Spawning() bool {
	return jsGet(c, "spawning", getBool).(bool)
}

func (c *Creep) Fatigue() int {
	return jsGet(c, "fatigue", getInt).(int)
}

func (c *Creep) TicksToLive() int {
	return jsGet(c, "ticksToLive", getInt).(int)
}

func (c *Creep) Store() *Store {
	return jsGet(c, "store", getStore).(*Store)
}

func (c *Creep) Body() CreepBody {
	return jsGet(c, "body", func(ref js.Value, property string) interface{} {
		jsBody := ref.Get(property)
		partCount := jsBody.Length()
		result := make(CreepBody, partCount)
		for i := 0; i < partCount; i++ {
			part := jsBody.Index(i)
			result[i] = CBodyPart(part.String())
		}
		return result
	}).(CreepBody)
}

func (c *Creep) Attack(target IDamageable) error {
	result := jsCall(c.ref, "attack", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) AttackController(target *Controller) error {
	result := jsCall(c.ref, "attackController", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Build(target *ConstructionSite) error {
	result := jsCall(c.ref, "build", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) CancelOrder(methodName string) error {
	result := jsCall(c.ref, "cancelOrder", methodName).Int()
	return returnErr(result)
}

func (c *Creep) ClaimController(target *Controller) error {
	result := jsCall(c.ref, "claimController", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Dismantle(target IStructure) error {
	result := jsCall(c.ref, "dismantle", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Drop(resource CResource, amount int) error {
	jsAmount := js.ValueOf(amount)
	if amount == -1 {
		jsAmount = js.Undefined()
	}
	result := jsCall(c.ref, "drop", string(resource), jsAmount).Int()
	return returnErr(result)
}

func (c *Creep) GenerateSafeMode(target *Controller) error {
	result := jsCall(c.ref, "generateSafeMode", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) GetActiveBodyParts(partType CBodyPart) int {
	return jsCall(c.ref, "getActiveBodyparts", string(partType)).Int()
}

func (c *Creep) Harvest(target IRoomPosition) error {
	result := jsCall(c.ref, "harvest", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Heal(target Creep) error {
	result := jsCall(c.ref, "heal", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Move(direction CDirection) error {
	result := jsCall(c.ref, "move", int(direction)).Int()
	return returnErr(result)
}

func (c *Creep) MoveTo(target IRoomPosition) error { // TODO opts
	result := jsCall(c.ref, "moveTo", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) MoveByPath(path Path) error {
	result := jsCall(c.ref, "moveByPath", packPath(path)).Int()
	return returnErr(result)
}

func (c *Creep) NotifyWhenAttacked(enabled bool) error {
	result := jsCall(c.ref, "notifyWhenAttacked", enabled).Int()
	return returnErr(result)
}

func (c *Creep) Pickup(target IRoomObject) error {
	result := jsCall(c.ref, "pickup", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Pull(target Creep) error {
	result := jsCall(c.ref, "pull", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) RangedAttack(target IDamageable) error {
	result := jsCall(c.ref, "rangedAttack", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) RangedHeal(target Creep) error {
	result := jsCall(c.ref, "rangedHeal", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) RangedMassAttack() error {
	result := jsCall(c.ref, "rangedMassAttack").Int()
	return returnErr(result)
}

func (c *Creep) Repair(target IStructure) error {
	result := jsCall(c.ref, "repair", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) ReserveController(target *Controller) error {
	result := jsCall(c.ref, "reserveController", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Say(message string, public bool) error {
	result := jsCall(c.ref, "say", message, public).Int()
	return returnErr(result)
}

func (c *Creep) SignController(target *Controller, text string) error {
	result := jsCall(c.ref, "signController", target.iRef(), text).Int()
	return returnErr(result)
}

func (c *Creep) Suicide() error {
	result := jsCall(c.ref, "suicide").Int()
	return returnErr(result)
}

func (c *Creep) Transfer(target IRoomObject, resource CResource, amount int) error {
	jsAmount := js.ValueOf(amount)
	if amount == -1 {
		jsAmount = js.Undefined()
	}
	result := jsCall(c.ref, "transfer", target.iRef(), string(resource), jsAmount).Int()
	return returnErr(result)
}

func (c *Creep) UpgradeController(target *Controller) error {
	result := jsCall(c.ref, "upgradeController", target.iRef()).Int()
	return returnErr(result)
}

func (c *Creep) Withdraw(target IRoomObject, resource CResource, amount int) error {
	jsAmount := js.ValueOf(amount)
	if amount == -1 {
		jsAmount = js.Undefined()
	}
	result := jsCall(c.ref, "withdraw", target.iRef(), string(resource), jsAmount).Int()
	return returnErr(result)
}
