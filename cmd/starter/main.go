package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
	rs "github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
)

func main() {
	goscreeps.Start(onReset, loop)
}

var room *rs.Room

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")
	room = s.Game.Rooms()["W8N3"]
	spawns := room.FindObject(rs.FIND_MY_SPAWNS, nil)
	for _, spawn := range spawns {
		_, err := spawn.Pos().CreateFlag("SpawnFlag", rs.COLOR_CYAN, rs.COLOR_RED)
		if err != nil {
			console.Log(err.Error())
		}
	}
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Starting loop", s.Game.Time())
}
