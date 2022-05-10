package resources

import "syscall/js"

type Memory struct {
	ref js.Value
}

func (m *Memory) WasmUpdate() {
	m.ref = jsGlobal.Get("Memory")
}

func (m *Memory) Get(property string) js.Value {
	return m.ref.Get(property)
}

func (m *Memory) Set(property string, value interface{}) {
	m.ref.Set(property, value)
}

func (m *Memory) Delete(property string) {
	m.ref.Delete(property)
}
