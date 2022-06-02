package resources

import (
	"regexp"
	"strings"
	"syscall/js"
)

var jsObject js.Value
var jsGlobal js.Value
var re = regexp.MustCompile("[a-zA-Z0-9\\-\\_\\#\\]]+?\\s")

func WasmUpdate() {
	jsGlobal = js.Global()
	jsObject = jsGlobal.Get("Object")
}

type IReference interface {
	iRef() js.Value
	deRef(ref js.Value) IReference
	iCache() map[string]interface{}
}

func Ref(ref IReference) js.Value {
	return ref.iRef()
}

func jsGet(ref IReference, property string, get func(ref js.Value, property string) interface{}) interface{} {
	c := ref.iCache()
	if c[property] == nil {
		c[property] = get(ref.iRef(), property)
	}
	return c[property]
}

func jsCall(ref js.Value, method string, args ...interface{}) js.Value {
	return ref.Call(method, args...)
}

func getInt(ref js.Value, property string) interface{} {
	return ref.Get(property).Int()
}

func getString(ref js.Value, property string) interface{} {
	return ref.Get(property).String()
}

func getBool(ref js.Value, property string) interface{} {
	return ref.Get(property).Bool()
}

func deRefUnknown(ref js.Value) IReference {
	typeStr := jsCall(ref, "toString").String()
	matches := re.FindAllString(typeStr, -1)
	if matches == nil {
		panic("Unable to match: \"" + typeStr + "\"")
	}
	typeStr = matches[len(matches)-1]
	typeStr = strings.TrimSpace(typeStr)
	var result IReference
	switch typeStr {
	case "pos":
		result = &RoomPosition{}
	case "spawn":
		result = &StructureSpawn{}
	case "source":
		result = &Source{}
	// TODO more
	default:
		panic("Unknown Type: \"" + typeStr + "\"")
	}
	return result.deRef(ref).(IReference)
}
