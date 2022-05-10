package goscreeps

import (
	"fmt"
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"runtime"
	"syscall/js"
)

type Screeps struct {
	Game *resources.Game
	RawMemory *resources.RawMemory
	InterShardMemory *resources.InterShardMemory
	Memory *resources.Memory
}

var block chan bool

func Start(loop func(s Screeps, console Console))  {
	defer func() {
		if r := recover(); r != nil {
			console.Log(fmt.Sprint(r))
			console.Log(getStack())
		}
	}()
	block = make(chan bool)


	resources.WasmUpdate()

	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		block <- false
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	s := Screeps{
		Game:      new(resources.Game),
		RawMemory: new(resources.RawMemory),
		InterShardMemory: new(resources.InterShardMemory),
		Memory: new(resources.Memory),
	}
	s.InterShardMemory.WasmUpdate()
	for {
		<-block
		s.Game.WasmUpdate()
		s.RawMemory.WasmUpdate()
		s.Memory.WasmUpdate()
		func() {
			defer func() {
				if r := recover(); r != nil {
					console.Log(fmt.Sprint(r))
					console.Log(getStack())
				}
			}()
			loop(s, console)
		}()
		s.RawMemory.WasmSave()
		runtime.GC()
	}
}

func getStack() string {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}
