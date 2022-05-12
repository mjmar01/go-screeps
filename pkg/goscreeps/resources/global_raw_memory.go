package resources

import (
	"strconv"
	"syscall/js"
)

type ForeignSegment struct {
	Username string
	Id       int
	Data     string
}

type RawMemory struct {
	ref      js.Value
	segments map[int]string
}

func (rm *RawMemory) WasmUpdate() {
	rm.ref = jsGlobal.Get("RawMemory")
	rm.loadSegments()
}

func (rm *RawMemory) WasmSave() {
	rm.saveSegments()
}

func (rm *RawMemory) Segments() map[int]string {
	return rm.segments
}

func (rm *RawMemory) ForeignSegment() *ForeignSegment {
	jsForeignSegment := rm.ref.Get("ForeignSegment")
	if jsForeignSegment.IsNull() {
		return nil
	} else {
		result := new(ForeignSegment)
		result.Username = jsForeignSegment.Get("username").String()
		result.Id = jsForeignSegment.Get("id").Int()
		result.Data = jsForeignSegment.Get("data").String()
		return result
	}
}

func (rm *RawMemory) Get() string {
	return rm.ref.Call("get").String()
}

func (rm *RawMemory) Set(value string) {
	rm.ref.Call("set", value)
}

func (rm *RawMemory) SetActiveSegments(ids []int) {
	rm.ref.Call("setActiveSegments", ids)
}

func (rm *RawMemory) SetActiveForeignSegment(username string, id *int) {
	var jsId js.Value
	if id == nil {
		jsId = js.Null()
	} else {
		jsId = js.ValueOf(*id)
	}
	rm.ref.Call("setActiveForeignSegment", username, jsId)
}

func (rm *RawMemory) SetDefaultPublicSegment(id int) {
	rm.ref.Call("setDefaultPublicSegment", id)
}

func (rm *RawMemory) SetPublicSegments(ids []int) {
	rm.ref.Call("setPublicSegments", ids)
}

func (rm *RawMemory) loadSegments() {
	for k := range rm.segments {
		delete(rm.segments, k)
	}

	jsSegments := rm.ref.Get("segments")

	entries := jsObject.Call("entries", jsSegments)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key, _ := strconv.Atoi(entry.Index(0).String())
		value := entry.Index(1).String()
		rm.segments[key] = value
		jsSegments.Set(entry.Index(0).String(), js.Undefined())
	}
}

func (rm *RawMemory) saveSegments() {
	jsSegments := rm.ref.Get("segments")
	for k := range rm.segments {
		jsSegments.Set(strconv.Itoa(k), rm.segments[k])
	}
}
