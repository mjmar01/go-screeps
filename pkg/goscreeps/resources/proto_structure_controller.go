package resources

import (
	"syscall/js"
	"time"
)

type Controller struct {
	ref    js.Value
	cached map[string]bool

	pos     *RoomPosition
	effects []Effect
	room    *Room

	id                string
	structureType     StructureConst
	my                bool
	owner             string
	isPowerEnabled    bool
	level             int
	progress          int
	progressTotal     int
	reservation       Reservation
	safeMode          int
	safeModeAvailable int
	safeModeCooldown  int
	sign              Sign
	ticksToDowngrade  int
	upgradeBlocked    int
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
		cached: make(map[string]bool),
	}
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
	if !c.cached["pos"] {
		c.pos = pos(c.ref)
		c.cached["pos"] = true
	}
	return c.pos
}

func (c *Controller) Effects() []Effect {
	if !c.cached["effects"] {
		c.effects = effects(c.ref)
		c.cached["effects"] = true
	}
	return c.effects
}

func (c *Controller) Room() *Room {
	if !c.cached["room"] {
		c.room = (&Room{}).deRef(c.ref).(*Room)
		c.cached["room"] = true
	}
	return c.room
}

func (c *Controller) Hits() int {
	return 0
}

func (c *Controller) HitsMax() int {
	return 0
}

func (c *Controller) StructureType() StructureConst {
	return STRUCTURE_CONTROLLER
}

func (c *Controller) Destroy() ScreepsError {
	return nil
}

func (c *Controller) IsActive() bool {
	return true
}

func (c *Controller) NotifyWhenAttacked(enabled bool) ScreepsError {
	return notifyWhenAttacked(c.ref, enabled)
}

func (c *Controller) My() bool {
	if !c.cached["my"] {
		c.my = jsGet(c.ref, "my").Bool()
		c.cached["my"] = true
	}
	return c.my
}

func (c *Controller) Owner() string {
	if !c.cached["owner"] {
		c.owner = jsGet(c.ref, "owner").String()
		c.cached["owner"] = true
	}
	return c.owner
}

func (c *Controller) Id() string {
	if !c.cached["id"] {
		c.id = jsGet(c.ref, "id").String()
		c.cached["id"] = true
	}
	return c.id
}

func (c *Controller) IsPowerEnabled() bool {
	if !c.cached["isPowerEnabled"] {
		c.isPowerEnabled = jsGet(c.ref, "isPowerEnabled").Bool()
		c.cached["isPowerEnabled"] = true
	}
	return c.isPowerEnabled
}

func (c *Controller) Level() int {
	if !c.cached["level"] {
		c.level = jsGet(c.ref, "level").Int()
		c.cached["level"] = true
	}
	return c.level
}

func (c *Controller) Progress() int {
	if !c.cached["progress"] {
		c.progress = jsGet(c.ref, "progress").Int()
		c.cached["progress"] = true
	}
	return c.progress
}

func (c *Controller) ProgressTotal() int {
	if !c.cached["progressTotal"] {
		c.progressTotal = jsGet(c.ref, "progressTotal").Int()
		c.cached["progressTotal"] = true
	}
	return c.progressTotal
}

func (c *Controller) Reservation() Reservation {
	if !c.cached["reservation"] {
		ref := jsGet(c.ref, "reservation")
		c.reservation = Reservation{
			username:   jsGet(ref, "username").String(),
			ticksToEnd: jsGet(ref, "ticksToEnd").Int(),
		}
		c.cached["reservation"] = true
	}
	return c.reservation
}

func (c *Controller) SafeMode() int {
	if !c.cached["safeMode"] {
		c.safeMode = jsGet(c.ref, "safeMode").Int()
		c.cached["safeMode"] = true
	}
	return c.safeMode
}

func (c *Controller) SafeModeAvailable() int {
	if !c.cached["safeModeAvailable"] {
		c.safeModeAvailable = jsGet(c.ref, "safeModeAvailable").Int()
		c.cached["safeModeAvailable"] = true
	}
	return c.safeModeAvailable
}

func (c *Controller) SafeModeCooldown() int {
	if !c.cached["safeModeCooldown"] {
		c.safeModeCooldown = jsGet(c.ref, "safeModeCooldown").Int()
		c.cached["safeModeCooldown"] = true
	}
	return c.safeModeCooldown
}

func (c *Controller) Sign() Sign {
	if !c.cached["sign"] {
		ref := jsGet(c.ref, "sign")
		c.sign = Sign{
			username: jsGet(ref, "username").String(),
			text:     jsGet(ref, "text").String(),
			time:     jsGet(ref, "time").Int(),
			datetime: time.Time{}, // TODO
		}
		c.cached["sign"] = true
	}
	return c.sign
}

func (c *Controller) TicksToDowngrade() int {
	if !c.cached["ticksToDowngrade"] {
		c.ticksToDowngrade = jsGet(c.ref, "ticksToDowngrade").Int()
		c.cached["ticksToDowngrade"] = true
	}
	return c.ticksToDowngrade
}

func (c *Controller) UpgradeBlocked() int {
	if !c.cached["upgradeBlocked"] {
		c.upgradeBlocked = jsGet(c.ref, "upgradeBlocked").Int()
		c.cached["upgradeBlocked"] = true
	}
	return c.upgradeBlocked
}

func (c *Controller) ActivateSafeMode() ScreepsError {
	result := jsCall(c.ref, "activateSafeMode").Int()
	return ReturnErr(result)
}
func (c *Controller) Unclaim() ScreepsError {
	result := jsCall(c.ref, "unclaim").Int()
	return ReturnErr(result)
}
