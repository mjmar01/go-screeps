package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	rs "github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"strconv"
)

func main() {
	goscreeps.Start(onReset, loop)
}

var mainRoom *rs.Room
var source *rs.Source
var spawn *rs.StructureSpawn
var controller *rs.Controller
var fill map[string]bool
var creepCount int

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	spawn = s.Game.Spawns()["0x0000"]
	mainRoom = s.Game.Rooms()["W1N8"]
	source = mainRoom.Find(rs.FIND_SOURCES, nil)[1].(*rs.Source)
	controller = mainRoom.Controller()
	creepCount = 0
	//for _, creep := range s.Game.Creeps() {
	//	creep.Suicide()
	//}
	fill = make(map[string]bool)
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	spawn = s.Game.Spawns()["0x0000"]
	creeps := s.Game.Creeps()
	for _, creep := range creeps {
		if fill[creep.Id()] {
			move(creep, source, 1)
			creep.Harvest(source)
			if creep.Store().GetFreeCapacity(nil) == 0 {
				fill[creep.Id()] = false
			}
		} else {
			move(creep, controller, 3)
			creep.UpgradeController(controller)
			if creep.Store().GetFreeCapacity(nil) == creep.Store().GetCapacity(nil) {
				fill[creep.Id()] = true
			}
		}
	}
	e := rs.RESOURCE_ENERGY
	if spawn.Store().GetUsedCapacity(&e) == 300 {
		spawn.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, strconv.Itoa(creepCount), nil)
		creepCount++
	}
}

func move(c *rs.Creep, target rs.IRoomPosition, r int) {
	if !c.Pos().InRangeTo(target, r) {
		c.MoveTo(target)
	}
}
