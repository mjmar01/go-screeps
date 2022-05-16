package main

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps"
)

func main() {
	goscreeps.Start(onReset, loop)
}

func onReset(s goscreeps.Screeps, console goscreeps.Console) {
	console.Log("Started once")
	sites := s.Game.Structures()
	console.Log(len(sites))
	for k, site := range sites {
		console.Log(k, string(site.StructureType()))
	}
}

func loop(s goscreeps.Screeps, console goscreeps.Console) {
	//console.Log("Starting loop", s.Game.Time())
}
