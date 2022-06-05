package resources

import (
	"regexp"
	"strings"
	"syscall/js"
)

var jsObject js.Value
var jsGlobal js.Value
var typeRegex = regexp.MustCompile("[a-zA-Z0-9()\\-\\_\\#\\]]+?\\s")

// WasmUpdate update global JS references. Will not update game state
func WasmUpdate() {
	jsGlobal = js.Global()
	jsObject = jsGlobal.Get("Object")
}

// IReference is the interface of any type that has an underlying JS value
type IReference interface {
	iRef() js.Value
	deRef(ref js.Value) IReference
	iCache() map[string]interface{}
}

// Ref returns the raw js.Value for an IReference. Used for logging purpose
func Ref(ref IReference) js.Value {
	return ref.iRef()
}

// retrieve and cache a value from JS using the passed function
func jsGet(ref IReference, property string, get func(ref js.Value, property string) interface{}) interface{} {
	c := ref.iCache()
	if c[property] == nil {
		c[property] = get(ref.iRef(), property)
	}
	return c[property]
}

// call a method of a JS object
func jsCall(ref js.Value, method string, args ...interface{}) js.Value {
	return ref.Call(method, args...)
}

// retrieve a simple int from a js.Value
func getInt(ref js.Value, property string) interface{} {
	return ref.Get(property).Int()
}

// retrieve a simple string from a js.Value
func getString(ref js.Value, property string) interface{} {
	return ref.Get(property).String()
}

// retrieve a simple bool from a js.Value
func getBool(ref js.Value, property string) interface{} {
	return ref.Get(property).Bool()
}

// convert a js.Value of unknown type to the corresponding go struct
func deRefUnknown(ref js.Value) IReference {
	typeStr := jsCall(ref, "toString").String()
	matches := typeRegex.FindAllString(typeStr, -1)
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
	case "(extension)":
		result = &StructureExtension{}
	case "(controller)":
		result = &Controller{}
	// TODO more
	default:
		panic("Unknown Type: \"" + typeStr + "\"")
	}
	return result.deRef(ref).(IReference)
}
