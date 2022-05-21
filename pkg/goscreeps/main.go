package goscreeps

import (
	"fmt"
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"runtime"
	"syscall/js"
)

type Screeps struct {
	Game             *resources.Game
	RawMemory        *resources.RawMemory
	InterShardMemory *resources.InterShardMemory
	Memory           *resources.Memory
}

var loopTrigger chan bool
var cpu resources.Cpu

func Start(onReset, loop func(s Screeps, console Console)) {
	defer func() {
		if r := recover(); r != nil {
			console.Log(fmt.Sprint(r))
			console.Log(getStack())
		}
	}()
	loopTrigger = make(chan bool)

	resources.WasmUpdate()

	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loopTrigger <- true
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	s := Screeps{
		Game:             new(resources.Game),
		RawMemory:        new(resources.RawMemory),
		InterShardMemory: new(resources.InterShardMemory),
		Memory:           new(resources.Memory),
	}
	s.Game.WasmUpdate()
	s.RawMemory.WasmUpdate()
	s.Memory.WasmUpdate()
	s.InterShardMemory.WasmUpdate()
	cpu = s.Game.Cpu()
	func() {
		defer func() {
			if r := recover(); r != nil {
				console.Log(fmt.Sprint(r))
				console.Log(getStack())
			}
		}()
		onReset(s, console)
	}()
	for {
		<-loopTrigger
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
		stats := cpu.GetHeapStatistics()
		used := float64(stats.TotalHeapSize) / float64(stats.TotalHeapSize+stats.TotalAvailableSize)
		if used >= 0.9 {
			runtime.GC()
			console.Log("goscreeps: cleared garbage")
		}
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
