package resources

import "syscall/js"

var jsObject js.Value
var jsGlobal js.Value

func WasmUpdate() {
	jsGlobal = js.Global()
	jsObject = jsGlobal.Get("Object")
}
