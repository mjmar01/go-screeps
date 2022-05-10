package resources

import "syscall/js"

type Game struct {
	ref js.Value
	cached map[string]bool

	cpu Cpu

	gcl GlobalControlLevel
	gpl GlobalPowerLevel

	resources map[AccountResource]int
	rooms map[string]Room

	shard Shard

	time int
}

type Shard struct {
	Name string
	Type string
	Ptr  bool
}

type GlobalControlLevel struct {
	Level         int
	Progress      int
	ProgressTotal int
}
type GlobalPowerLevel struct {
	Level         int
	Progress      int
	ProgressTotal int
}

func (g *Game) WasmUpdate() {
	g.ref = jsGlobal.Get("Game")
	g.cached = make(map[string]bool)
}

// TODO constructionSites, creeps, flags,

func (g *Game) Gcl() GlobalControlLevel {
	if !g.cached["gcl"] {
		jsGcl := g.ref.Get("gcl")
		g.gcl = GlobalControlLevel{
			Level:         jsGcl.Get("level").Int(),
			Progress:      jsGcl.Get("progress").Int(),
			ProgressTotal: jsGcl.Get("progressTotal").Int(),
		}
		g.cached["gcl"] = true
	}
	return g.gcl
}

func (g *Game) Gpl() GlobalPowerLevel {
	if !g.cached["gpl"] {
		jsGpl := g.ref.Get("gpl")
		g.gpl = GlobalPowerLevel{
			Level:         jsGpl.Get("level").Int(),
			Progress:      jsGpl.Get("progress").Int(),
			ProgressTotal: jsGpl.Get("progressTotal").Int(),
		}
		g.cached["gpl"] = true
	}
	return g.gpl
}

// TODO map, market, powerCreeps

func (g *Game) Resources() map[AccountResource]int {
	if !g.cached["resources"] {
		jsResources := g.ref.Get("resources")

		entries := jsObject.Call("entries", jsResources)
		length := entries.Get("length").Int()
		result := make(map[AccountResource]int, length)
		for i := 0; i < length; i++ {
			entry := entries.Index(i)
			key := AccountResource(entry.Index(0).String())
			value := entry.Index(1).Int()
			result[key] = value
		}
		g.resources = result
		g.cached["resources"] = true
	}
	return g.resources
}

func (g *Game) Rooms() map[string]Room {
	if !g.cached["rooms"] {
		jsRooms := g.ref.Get("rooms")
		entries := jsObject.Call("entries", jsRooms)
		length := entries.Get("length").Int()
		result := make(map[string]Room, length)
		for i := 0; i < length; i++ {
			entry := entries.Index(i)
			key := entry.Index(0).String()
			value := entry.Index(1)
			result[key] = Room{
				ref: value,
			}
		}
		g.rooms = result
		g.cached["rooms"] = true
	}
	return g.rooms
}

func (g *Game) Shard() Shard {
	if !g.cached["shard"] {
		jsShard := g.ref.Get("shard")
		g.shard = Shard{
			Name: jsShard.Get("name").String(),
			Type: jsShard.Get("type").String(),
			Ptr:  jsShard.Get("ptr").Bool(),
		}
		g.cached["shard"] = true
	}
	return g.shard
}

// TODO spawns, structures

func (g *Game) Time() int {
	if !g.cached["time"] {
		g.time = g.ref.Get("time").Int()
		g.cached["time"] = true
	}
	return g.time
}

func (g *Game) Notify(message string, groupInterval *int) {
	var jsGroupInterval js.Value
	if groupInterval == nil {
		jsGroupInterval = js.Undefined()
	} else {
		jsGroupInterval = js.ValueOf(*groupInterval)
	}
	g.ref.Call("notify", message, jsGroupInterval)
}