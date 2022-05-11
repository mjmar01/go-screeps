package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
)

func main() {
	goscreeps.Start(onInit, loop)
}

func onInit(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")

	room := s.Game.Rooms()["W4N8"]
	name, err := room.CreateFlagAtCoords(16, 16, "", resources.COLOR_WHITE, resources.COLOR_RED)
	if err != nil {
		console.Log(err.Error())
	} else {
		console.Log("Created flag:", name)
	}

	err = room.CreateConstructionSiteAtTarget(resources.NewRoomPosition(18, 19, room.Name()), resources.STRUCTURE_CONTAINER, "")
	if err != nil {
		console.Log(err.Error())
	}
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Starting loop", s.Game.Time())
	stats := s.Game.Cpu().GetHeapStatistics()
	used := float64(stats.TotalHeapSize) / float64(stats.TotalHeapSize+stats.TotalAvailableSize)
	console.Log("Heap used:", used)
}
