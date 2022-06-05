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
var mainRoomName string
var source *rs.Source
var spawn *rs.StructureSpawn
var controller *rs.Controller
var fill map[string]bool
var creepCount int

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	spawn = s.Game.Spawns()["0x0000"]
	mainRoomName = spawn.Pos().RoomName()
	mainRoom = s.Game.Rooms()[mainRoomName]
	source = s.Game.GetObjectById("ab9e0774d1c107c").(*rs.Source)
	controller = mainRoom.Controller()
	creepCount = len(s.Game.Creeps())
	fill = make(map[string]bool)
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	mainRoom = s.Game.Rooms()[mainRoomName]
	spawn = s.Game.Spawns()["0x0000"]
	creeps := s.Game.Creeps()
	flags := s.Game.Flags()
	for _, flag := range flags {
		err := flag.Pos().CreateConstructionSite(rs.STRUCTURE_EXTENSION, "")
		if err != nil {
			flag.Remove()
		}
	}
c:
	for _, creep := range creeps {
		if fill[creep.Id()] {
			move(creep, source, 1)
			creep.Harvest(source)
			if creep.Store().GetFreeCapacity(rs.RESOURCE_ENERGY) == 0 {
				fill[creep.Id()] = false
			}
		} else {
			sites := s.Game.ConstructionSites()
			if len(sites) > 0 {
				for _, site := range s.Game.ConstructionSites() {
					move(creep, site, 3)
					creep.Build(site)
					if creep.Store().GetFreeCapacity(rs.RESOURCE_ANY) == creep.Store().GetCapacity(rs.RESOURCE_ANY) {
						fill[creep.Id()] = true
					}
					break
				}
				continue
			}
			structures := mainRoom.Find(rs.FIND_MY_STRUCTURES, nil)
			var ext []*rs.StructureExtension
			for _, structure := range structures {
				switch structure.(type) {
				case *rs.StructureExtension:
					if structure.(*rs.StructureExtension).Store().GetFreeCapacity(rs.RESOURCE_ENERGY) > 0 {
						ext = append(ext, structure.(*rs.StructureExtension))
					}
				}
			}
			for _, extension := range ext {
				move(creep, extension, 1)
				creep.Transfer(extension, rs.RESOURCE_ENERGY, -1)
				if creep.Store().GetFreeCapacity(rs.RESOURCE_ANY) == creep.Store().GetCapacity(rs.RESOURCE_ANY) {
					fill[creep.Id()] = true
				}
				continue c
			}
			if spawn.Store().GetFreeCapacity(rs.RESOURCE_ENERGY) > 0 {
				move(creep, spawn, 1)
				creep.Transfer(spawn, rs.RESOURCE_ENERGY, -1)
				if creep.Store().GetFreeCapacity(rs.RESOURCE_ANY) == creep.Store().GetCapacity(rs.RESOURCE_ANY) {
					fill[creep.Id()] = true
				}
				continue
			}
			move(creep, controller, 3)
			creep.UpgradeController(controller)
			if creep.Store().GetFreeCapacity(rs.RESOURCE_ANY) == creep.Store().GetCapacity(rs.RESOURCE_ANY) {
				fill[creep.Id()] = true
			}
		}
	}
	if mainRoom.EnergyAvailable() == mainRoom.EnergyCapacityAvailable() && len(creeps) < 3 {
		e := mainRoom.EnergyAvailable()
		e = e / 300
		body := rs.CreepBody{}
		for i := 0; i < e*2; i++ {
			body = append(body, rs.WORK)
		}
		for i := 0; i < e; i++ {
			body = append(body, rs.CARRY)
		}
		for i := 0; i < e; i++ {
			body = append(body, rs.MOVE)
		}
		spawn.SpawnCreep(body, strconv.Itoa(creepCount), nil)
		creepCount++
	}
}

func move(c *rs.Creep, target rs.IRoomPosition, r int) {
	if !c.Pos().InRangeTo(target, r) {
		c.MoveTo(target)
	}
}
