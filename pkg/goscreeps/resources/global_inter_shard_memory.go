package resources

import "syscall/js"

type InterShardMemory struct {
	ref    js.Value
	Exists bool
}

func (ism *InterShardMemory) WasmUpdate() {
	sharded := !jsGlobal.Get("Game").Get("cpu").Get("generatePixel").IsUndefined()
	if sharded {
		ism.ref = jsGlobal.Get("InterShardMemory")
		ism.Exists = true
	} else {
		ism.ref = js.Null()
		ism.Exists = false
	}
}

func (ism *InterShardMemory) GetLocal() string {
	return ism.ref.Call("getLocal").String()
}

func (ism *InterShardMemory) SetLocal(value string) {
	ism.ref.Call("setLocal", value)
}

func (ism *InterShardMemory) GetRemote(shard string) string {
	jsResult := ism.ref.Call("getRemote", shard)
	if jsResult.IsNull() {
		return ""
	} else {
		return jsResult.String()
	}
}
