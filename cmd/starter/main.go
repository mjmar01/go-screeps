package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	rs "github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
)

func main() {
	goscreeps.Start(onReset, loop)
}

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")
	spawns := s.Game.Spawns()
	for _, spawn := range spawns {
		err := spawn.SpawnCreep(rs.CreepBody{rs.WORK, rs.WORK, rs.CARRY, rs.MOVE}, "Jeff", nil)
		if err != nil {
			console.Log(err.Error())
		}
	}
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	flags := s.Game.Flags()
	for _, flag := range flags {
		console.Log(flag.Name())
		flag.Remove()
	}
}
