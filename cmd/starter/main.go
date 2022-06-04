package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	rs "github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"strconv"
)

func main() {
	goscreeps.Start(onReset, loop, nil)
}

var mainRoom *rs.Room
var source *rs.Source
var spawn *rs.StructureSpawn
var controller *rs.Controller
var fill map[string]bool
var creepCount int

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	spawn = s.Game.Spawns()["0x0000"]
	mainRoomName := spawn.Pos().RoomName()
	mainRoom = s.Game.Rooms()[mainRoomName]
	source = mainRoom.Find(rs.FIND_SOURCES, nil)[1].(*rs.Source)
	controller = mainRoom.Controller()
	creepCount = len(s.Game.Creeps())
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
	if spawn.Store().GetUsedCapacity(rs.RESOURCE_ENERGY) == 300 {
		spawn.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, strconv.Itoa(creepCount), nil)
		creepCount++
	}
}

func move(c *rs.Creep, target rs.IRoomPosition, r int) {
	if !c.Pos().InRangeTo(target, r) {
		c.MoveTo(target)
	}
}
