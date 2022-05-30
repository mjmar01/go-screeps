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
var spawn *rs.StructureSpawn
var controller *rs.Controller
var creep *rs.Creep
var fill bool

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")
	spawns := s.Game.Spawns()
	for _, s := range spawns {
		err := s.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, "Jeff", nil)
		if err != nil {
			console.Log(err.Error())
		}
		spawn = s
	}
	mainRoom = s.Game.Rooms()["W1N8"]
	source = mainRoom.Find(rs.FIND_SOURCES, nil)[1].(*rs.Source)
	controller = mainRoom.Controller()
	fill = true
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	creeps := s.Game.Creeps()
	for _, c := range creeps {
		creep = c
	}
	if len(creeps) == 0 {
		spawn.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, "Jeff", nil)
	}

	if fill {
		creep.MoveTo(source)
		creep.Harvest(source)
		if creep.Store().GetFreeCapacity(nil) == 0 {
			fill = false
		}
	} else {
		creep.MoveTo(controller)
		creep.UpgradeController(controller)
		if creep.Store().GetFreeCapacity(nil) == creep.Store().GetCapacity(nil) {
			fill = true
		}
	}
}
