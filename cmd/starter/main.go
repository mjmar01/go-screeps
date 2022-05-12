package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
)

func main() {
	goscreeps.Start(onReset, loop)
}

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once", s.Game.Time())
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Starting loop", s.Game.Time())
	stats := s.Game.Cpu().GetHeapStatistics()
	used := float64(stats.TotalHeapSize) / float64(stats.TotalHeapSize+stats.TotalAvailableSize)
	console.Log("Heap used:", used)
}
