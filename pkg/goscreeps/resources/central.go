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
}

func jsGet(ref js.Value, property string) js.Value {
	return ref.Get(property)
}

func jsCall(ref js.Value, method string, args ...interface{}) js.Value {
	return ref.Call(method, args...)
}

func Ref(v IReference) js.Value {
	return v.iRef()
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
