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
	spawn := room.Find(rs.FIND_MY_SPAWNS, nil)[0]
	switch spawn.(type) {
	case *rs.RoomPosition:
	case *rs.RoomObject:
	}

	pos := room.GetPositionAt(15, 15)
	err := pos.CreateConstructionSite(rs.STRUCTURE_EXTENSION, "")
	if err != nil {
		console.Log(err.Error())
	}
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	//console.Log("Starting loop", s.Game.Time())
}
