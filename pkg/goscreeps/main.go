package goscreeps

import (
	"fmt"
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"runtime"
	"syscall/js"
)

type Screeps struct {
	Game *resources.Game
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
		Game: new(resources.Game),
	}
	s.Game.WasmUpdate()
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
	runtime.GC()
	for {
		<-loopTrigger
		s.Game.WasmUpdate()
		func() {
			defer func() {
				if r := recover(); r != nil {
					console.Log(fmt.Sprint(r))
					console.Log(getStack())
				}
			}()
			loop(s, console)
		}()
		js.Global().Call("cleanupGo")
		stats := cpu.GetHeapStatistics()
		used := float64(stats.TotalHeapSize) / float64(stats.TotalHeapSize+stats.TotalAvailableSize)
		console.Log("used: ", used)
		// if used >= 0.9 {
		// 	runtime.GC()
		// 	console.Log("goscreeps: cleared garbage")
		// }
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
