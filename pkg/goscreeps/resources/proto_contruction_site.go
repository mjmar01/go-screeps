package resources

import "syscall/js"

type ConstructionSite struct {
	ref    js.Value
	cached map[string]interface{}
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
		cached: make(map[string]interface{}),
	}
}

func (c *ConstructionSite) iCache() map[string]interface{} {
	return c.cached
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
	return jsGet(c, "pos", getPos).(*RoomPosition)
}

func (c *ConstructionSite) Effects() []Effect {
	return jsGet(c, "effects", getEffects).([]Effect)
}

func (c *ConstructionSite) Room() *Room {
	return jsGet(c, "room", getRoom).(*Room)
}

func (c *ConstructionSite) My() bool {
	return jsGet(c, "my", getBool).(bool)
}

func (c *ConstructionSite) Owner() string {
	return jsGet(c, "owner", getString).(string)
}

func (c *ConstructionSite) Id() string {
	return jsGet(c, "id", getString).(string)
}

func (c *ConstructionSite) Progress() int {
	return jsGet(c, "progress", getInt).(int)
}
func (c *ConstructionSite) ProgressTotal() int {
	return jsGet(c, "progressTotal", getInt).(int)
}

func (c *ConstructionSite) StructureType() CStructure {
	return CStructure(jsGet(c, "structureType", getString).(string))
}
