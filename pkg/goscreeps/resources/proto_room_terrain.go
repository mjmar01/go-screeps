package resources

import (
	"syscall/js"
)

type Terrain struct {
	ref    js.Value
	cached map[string]interface{}
}

func (t *Terrain) iRef() js.Value {
	return t.ref
}

func (t *Terrain) deRef(ref js.Value) IReference {
	if ref.IsNull() {
		return nil
	}
	return &Terrain{
		ref:    ref,
		cached: make(map[string]interface{}),
	}
}

func (t *Terrain) iCache() map[string]interface{} {
	return t.cached
}

func NewTerrain(roomName string) *Terrain {
	result := jsGlobal.Get("Room").Get("Terrain").New(roomName)
	return &Terrain{
		ref:    result,
		cached: make(map[string]interface{}),
	}
}

func (t *Terrain) Get(x, y int) CTerrain {
	return jsGet(t, "getRawBuffer", getRawTerrain).([][]CTerrain)[x][y]
}

func (t *Terrain) GetRaw() [][]CTerrain {
	return jsGet(t, "getRawBuffer", getRawTerrain).([][]CTerrain)
}

func getRawTerrain(ref js.Value, property string) interface{} {
	jsList := ref.Call(property)
	idx := 0
	raw := make([][]CTerrain, 50, 50)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			raw[x][y] = CTerrain(jsList.Index(idx).Int())
		}
	}
	return raw
}
