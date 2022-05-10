package screeps

import (
	"syscall/js"
)

func (g game) ConstructionSites() map[string]ConstructionSite {
	jsConstructionSites := g.ref.Get("constructionSites")
	result := map[string]ConstructionSite{}

	entries := object.Call("entries", jsConstructionSites)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		result[entry.Index(0).String()] = ConstructionSite{RoomObject{
			ref: entry.Index(1),
		}}
	}

	return result
}

func (g game) Creeps() map[string]Creep {
	jsCreeps := g.ref.Get("creeps")
	result := map[string]Creep{}

	entries := object.Call("entries", jsCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Creep{RoomObject{
			ref: value,
		}}
	}

	return result
}

func (g game) Flags() map[string]Flag {
	jsFlags := g.ref.Get("flags")
	result := map[string]Flag{}

	entries := object.Call("entries", jsFlags)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Flag{RoomObject{
			ref: value,
		}}
	}

	return result
}



func (g game) Map() Map {
	return Map{ref: g.ref.Get("map")}
}

func (g game) Market() Market {
	return Market{ref: g.ref.Get("market")}
}

func (g game) PowerCreeps() map[string]PowerCreep {
	jsPowerCreeps := g.ref.Get("powerCreeps")
	result := map[string]PowerCreep{}

	entries := object.Call("entries", jsPowerCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = PowerCreep{RoomObject{
			ref: value,
		}}
	}

	return result
}

func (g game) Rooms() map[string]Room {
	jsPowerCreeps := g.ref.Get("rooms")
	result := map[string]Room{}

	entries := object.Call("entries", jsPowerCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Room{
			ref: value,
		}
	}

	return result
}


func (g game) Spawns() map[string]StructureSpawn {
	jsSpawns := g.ref.Get("spawns")
	result := map[string]StructureSpawn{}

	entries := object.Call("entries", jsSpawns)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = StructureSpawn{OwnedStructure{Structure{RoomObject{
			ref: value,
		}}}}
	}

	return result
}

func (g game) Structures() map[string]Structure {
	jsStructures := g.ref.Get("structures")
	result := map[string]Structure{}

	entries := object.Call("entries", jsStructures)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Structure{
			RoomObject{
				ref: value,
			},
		}
	}

	return result
}

func (g game) GetObjectById(id string) RoomObject {
	jsRoomObject := g.ref.Call("getObjectById", id)
	return RoomObject{
		ref: jsRoomObject,
	}
}

