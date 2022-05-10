package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
)

func main() {
	goscreeps.Start(loop)
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Starting loop", s.Game.Time())
	stats := s.Game.Cpu().GetHeapStatistics()
	used := float64(stats.TotalHeapSize)/float64(stats.TotalHeapSize + stats.TotalAvailableSize)
	console.Log("Heap used:", used)

	for _, room := range s.Game.Rooms() {
		err := room.CreateFlagAtCoords(15, 15, "", resources.COLOR_WHITE, resources.COLOR_RED)
		if err != nil {
			console.Log(err.Error())
		}
	}
}
