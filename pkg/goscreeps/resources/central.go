package resources

import "syscall/js"

var jsObject js.Value
var jsGlobal js.Value

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
