package resources

import (
	"syscall/js"
	"time"
)

type Controller struct {
	ref    js.Value
	cached map[string]interface{}
}

type Reservation struct {
	username   string
	ticksToEnd int
}

type Sign struct {
	username string
	text     string
	time     int
	datetime time.Time
}

func (c *Controller) iRef() js.Value {
	return c.ref
}

func (c *Controller) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Controller{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (c *Controller) iCache() map[string]interface{} {
	return c.cached
}

func (c *Controller) x() int {
	return c.Pos().x()
}

func (c *Controller) y() int {
	return c.Pos().y()
}

func (c *Controller) roomName() string {
	return c.Pos().roomName()
}

func (c *Controller) Pos() *RoomPosition {
	return jsGet(c, "pos", getPos).(*RoomPosition)
}

func (c *Controller) Effects() []Effect {
	return jsGet(c, "effects", getEffects).([]Effect)
}

func (c *Controller) Room() *Room {
	return jsGet(c, "room", getRoom).(*Room)
}

func (c *Controller) Hits() int {
	return 0
}

func (c *Controller) HitsMax() int {
	return 0
}

func (c *Controller) StructureType() CStructure {
	return STRUCTURE_CONTROLLER
}

func (c *Controller) Destroy() error {
	return nil
}

func (c *Controller) IsActive() bool {
	return true
}

func (c *Controller) NotifyWhenAttacked(enabled bool) error {
	return notifyWhenAttacked(c.ref, enabled)
}

func (c *Controller) My() bool {
	return jsGet(c, "my", getBool).(bool)
}

func (c *Controller) Owner() string {
	return jsGet(c, "owner", getString).(string)
}

func (c *Controller) Id() string {
	return jsGet(c, "id", getString).(string)
}

func (c *Controller) IsPowerEnabled() bool {
	return jsGet(c, "isPowerEnabled", getBool).(bool)
}

func (c *Controller) Level() int {
	return jsGet(c, "level", getInt).(int)
}

func (c *Controller) Progress() int {
	return jsGet(c, "progress", getInt).(int)
}

func (c *Controller) ProgressTotal() int {
	return jsGet(c, "progressTotal", getInt).(int)
}

func (c *Controller) Reservation() Reservation {
	return jsGet(c, "reservation", func(ref js.Value, property string) interface{} {
		reservation := ref.Get(property)
		return Reservation{
			username:   reservation.Get("username").String(),
			ticksToEnd: reservation.Get("ticksToEnd").Int(),
		}
	}).(Reservation)
}

func (c *Controller) SafeMode() int {
	return jsGet(c, "safeMode", getInt).(int)
}

func (c *Controller) SafeModeAvailable() int {
	return jsGet(c, "safeModeAvailable", getInt).(int)
}

func (c *Controller) SafeModeCooldown() int {
	return jsGet(c, "safeModeCooldown", getInt).(int)
}

func (c *Controller) Sign() Sign {
	return jsGet(c, "sign", func(ref js.Value, property string) interface{} {
		sign := ref.Get(property)
		return Sign{
			username: sign.Get("username").String(),
			text:     sign.Get("text").String(),
			time:     sign.Get("time").Int(),
			datetime: time.Time{}, // TODO
		}
	}).(Sign)
}

func (c *Controller) TicksToDowngrade() int {
	return jsGet(c, "ticksToDowngrade", getInt).(int)
}

func (c *Controller) UpgradeBlocked() int {
	return jsGet(c, "upgradeBlocked", getInt).(int)
}

func (c *Controller) ActivateSafeMode() error {
	result := jsCall(c.ref, "activateSafeMode").Int()
	return returnErr(result)
}
func (c *Controller) Unclaim() error {
	result := jsCall(c.ref, "unclaim").Int()
	return returnErr(result)
}

func getController(ref js.Value, property string) interface{} {
	return (&Controller{}).deRef(ref.Get(property))
}
