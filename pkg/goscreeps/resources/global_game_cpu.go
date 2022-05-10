package resources

import "syscall/js"

type Cpu struct {
	ref js.Value

	Limit        int
	TickLimit    int
	Bucket       int
	ShardLimits  map[string]int
	Unlocked     bool // doesn't exist on private servers
	UnlockedTime int
}

type HeapStatistics struct {
	TotalHeapSize           int
	TotalHeapSizeExecutable int
	TotalPhysicalSize       int
	TotalAvailableSize      int
	UsedHeapSize            int
	HeapSizeLimit           int
	MallocedMemory          int
	PeakMallocedMemory      int
	DoesZapGarbage          int
	ExternallyAllocatedSize int
}

func (g *Game) Cpu() Cpu {
	if !g.cached["cpu"] {
		jsCpu := g.ref.Get("cpu")
		result := Cpu{
			ref:       jsCpu,
			Limit:     jsCpu.Get("limit").Int(),
			TickLimit: jsCpu.Get("tickLimit").Int(),
			Bucket:    jsCpu.Get("bucket").Int(),
		}

		// shard limits
		jsShardLimits := jsCpu.Get("shardLimits")
		if !jsShardLimits.IsUndefined() {
			shardLimits := make(map[string]int)
			shardLimitsEntries := jsObject.Call("entries", jsCpu.Get("shardLimits"))
			shardLimitsLength := shardLimitsEntries.Get("length").Int()
			for i := 0; i < shardLimitsLength; i++ {
				entry := shardLimitsEntries.Index(i)
				key := entry.Index(0).String()
				value := entry.Index(1).Int()
				shardLimits[key] = value
			}
			result.ShardLimits = shardLimits
		}

		// unlocked
		jsUnlocked := jsCpu.Get("unlocked")
		if !jsUnlocked.IsUndefined() {
			unlocked := jsUnlocked.Bool()
			result.Unlocked = unlocked
		}

		// unlocked time
		jsUnlockedTime := jsCpu.Get("unlockedTime")
		if !jsUnlockedTime.IsUndefined() {
			unlockedTime := jsUnlockedTime.Int()
			result.UnlockedTime = unlockedTime
		}
		g.cpu = result
		g.cached["cpu"] = true
	}
	return g.cpu
}

func (c Cpu) GetHeapStatistics() HeapStatistics {
	jsHeapStatistics := c.ref.Call("getHeapStatistics")
	return HeapStatistics{
		TotalHeapSize:           jsHeapStatistics.Get("total_heap_size").Int(),
		TotalHeapSizeExecutable: jsHeapStatistics.Get("total_heap_size_executable").Int(),
		TotalPhysicalSize:       jsHeapStatistics.Get("total_physical_size").Int(),
		TotalAvailableSize:      jsHeapStatistics.Get("total_available_size").Int(),
		UsedHeapSize:            jsHeapStatistics.Get("used_heap_size").Int(),
		HeapSizeLimit:           jsHeapStatistics.Get("heap_size_limit").Int(),
		MallocedMemory:          jsHeapStatistics.Get("malloced_memory").Int(),
		PeakMallocedMemory:      jsHeapStatistics.Get("peak_malloced_memory").Int(),
		ExternallyAllocatedSize: jsHeapStatistics.Get("externally_allocated_size").Int(),
	}

}

func (c Cpu) GetUsed() float64 {
	return c.ref.Call("getUsed").Float()
}

func (c Cpu) Halt() {
	c.ref.Call("halt")
}

func (c Cpu) SetShardLimits(limits map[string]int) ScreepsError {
	jsLimits := map[string]interface{}{}
	for k, v := range limits {
		jsLimits[k] = v
	}

	result := c.ref.Call("setShardLimits", jsLimits).Int()
	return ReturnErr(ErrorCode(result))
}

func (c Cpu) Unlock() ScreepsError {
	result := c.ref.Call("unlock").Int()
	return ReturnErr(ErrorCode(result))
}

func (c Cpu) GeneratePixel() ScreepsError {
	result := c.ref.Call("generatePixel").Int()
	return ReturnErr(ErrorCode(result))
}
