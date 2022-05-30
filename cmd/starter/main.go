package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	rs "github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
)

func main() {
	goscreeps.Start(onReset, loop)
}

var mainRoom *rs.Room
var source *rs.Source
var controller *rs.Controller

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")
	spawns := s.Game.Spawns()
	for _, spawn := range spawns {
		err := spawn.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, "Jeff", nil)
		if err != nil {
			console.Log(err.Error())
		}
	}
	mainRoom = s.Game.Rooms()["W1N8"]
	source = mainRoom.Find(rs.FIND_SOURCES, nil)[0].(*rs.Source)
	controller = mainRoom.Controller()
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	creeps := s.Game.Creeps()

	stopGather := false
	for _, creep := range creeps {
		if creep.Store().GetFreeCapacity(nil) > 0 && !stopGather {
			err := creep.MoveTo(source)
			if err != nil {
				console.Log(err.Error())
			}
			err = creep.Harvest(source)
			if err != nil {
				console.Log(err.Error())
			}
		} else if creep.Store().GetFreeCapacity(nil) == 0 && !stopGather {
			creep.Say("I'm full", false)
			stopGather = true
		}

		if stopGather {
			err := creep.MoveTo(controller)
			if err != nil {
				console.Log(err.Error())
			}
			err = creep.UpgradeController(controller)
			if err != nil {
				console.Log(err.Error())
			}
			if creep.Store().GetFreeCapacity(nil) == creep.Store().GetCapacity(nil) {
				stopGather = false
			}
		}
	}
}
