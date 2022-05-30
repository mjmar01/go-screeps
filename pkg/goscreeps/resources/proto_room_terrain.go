package resources

import (
	"strconv"
	"syscall/js"
)

type Terrain struct {
	ref    js.Value
	cached map[string]bool

	raw [][]CTerrain
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
		cached: make(map[string]bool),
	}
}

func NewTerrain(roomName string) *Terrain {
	result := jsGet(jsGet(jsGlobal, "Room"), "Terrain").New(roomName)
	return &Terrain{
		ref:    result,
		cached: make(map[string]bool),
	}
}

func (t *Terrain) Get(x, y int) CTerrain {
	if !(t.cached[strconv.Itoa(x)+":"+strconv.Itoa(y)] || t.cached["raw"]) {
		t.raw[x][y] = CTerrain(jsCall(t.ref, "get", x, y).Int())
		t.cached[strconv.Itoa(x)+":"+strconv.Itoa(y)] = true
	}
	return t.raw[x][y]
}

func (t *Terrain) GetRaw() [][]CTerrain {
	if !t.cached["raw"] {
		jsList := jsCall(t.ref, "getRawBuffer")
		idx := 0
		for y := 0; y < 50; y++ {
			for x := 0; x < 50; x++ {
				t.raw[x][y] = CTerrain(jsList.Index(idx).Int())
			}
		}
		t.cached["raw"] = true
	}
	return t.raw
}
