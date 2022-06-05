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

type ScreepsOptions struct {
	MemoryLog bool
}

var loopTrigger chan bool
var cpu resources.Cpu
var opts *ScreepsOptions

func Start(onReset, loop func(s Screeps, console Console), options *ScreepsOptions) {
	// Log on failure
	defer func() {
		if r := recover(); r != nil {
			console.Log(fmt.Sprint(r))
			console.Log(getStack())
		}
	}()

	// Read options
	opts = options
	if opts == nil {
		opts = new(ScreepsOptions)
	}

	// Listen for loop events
	loopTrigger = make(chan bool)
	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loopTrigger <- true
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	// Init initial state
	resources.WasmUpdate()
	s := Screeps{
		Game: new(resources.Game),
	}
	s.Game.WasmUpdate()
	if opts.MemoryLog {
		cpu = s.Game.Cpu()
	}

	// Run reset func
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

	// Start loop
	for {
		// Wait for loop event
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

		// Fast GC
		js.Global().Call("cleanupGo")

		// Print memory info
		if opts.MemoryLog {
			stats := cpu.GetHeapStatistics()
			used := float64(stats.TotalHeapSize) / float64(stats.TotalHeapSize+stats.TotalAvailableSize)
			console.Log("Memory in use (%): ", used)
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
