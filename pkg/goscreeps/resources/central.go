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
	CC()
}
