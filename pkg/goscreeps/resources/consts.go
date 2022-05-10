package resources

type AccountResource string

const (
	SUBSCRIPTION_TOKEN AccountResource = "token"
	CPU_UNLOCK         AccountResource = "cpuUnlock"
	PIXEL              AccountResource = "pixel"
	ACCESS_KEY         AccountResource = "accessKey"
)

type EffectType int

const (
	PWR_GENERATE_OPS       EffectType = 1
	PWR_OPERATE_SPAWN      EffectType = 2
	PWR_OPERATE_TOWER      EffectType = 3
	PWR_OPERATE_STORAGE    EffectType = 4
	PWR_OPERATE_LAB        EffectType = 5
	PWR_OPERATE_EXTENSION  EffectType = 6
	PWR_OPERATE_OBSERVER   EffectType = 7
	PWR_OPERATE_TERMINAL   EffectType = 8
	PWR_DISRUPT_SPAWN      EffectType = 9
	PWR_DISRUPT_TOWER      EffectType = 10
	PWR_DISRUPT_SOURCE     EffectType = 11
	PWR_SHIELD             EffectType = 12
	PWR_REGEN_SOURCE       EffectType = 13
	PWR_REGEN_MINERAL      EffectType = 14
	PWR_DISRUPT_TERMINAL   EffectType = 15
	PWR_OPERATE_POWER      EffectType = 16
	PWR_FORTIFY            EffectType = 17
	PWR_OPERATE_CONTROLLER EffectType = 18
	PWR_OPERATE_FACTORY    EffectType = 19
	EFFECT_INVULNERABILITY EffectType = 1001
	EFFECT_COLLAPSE_TIMER  EffectType = 1002
)

type Color int

const (
	COLOR_RED    Color = 1
	COLOR_PURPLE Color = 2
	COLOR_BLUE   Color = 3
	COLOR_CYAN   Color = 4
	COLOR_GREEN  Color = 5
	COLOR_YELLOW Color = 6
	COLOR_ORANGE Color = 7
	COLOR_BROWN  Color = 8
	COLOR_GREY   Color = 9
	COLOR_WHITE  Color = 10
)

type FindType int

const (
	FIND_EXIT_TOP                   FindType = 1
	FIND_EXIT_RIGHT                 FindType = 3
	FIND_EXIT_BOTTOM                FindType = 5
	FIND_EXIT_LEFT                  FindType = 7
	FIND_EXIT                       FindType = 10
	FIND_CREEPS                     FindType = 101
	FIND_MY_CREEPS                  FindType = 102
	FIND_HOSTILE_CREEPS             FindType = 103
	FIND_SOURCES_ACTIVE             FindType = 104
	FIND_SOURCES                    FindType = 105
	FIND_DROPPED_RESOURCES          FindType = 106
	FIND_STRUCTURES                 FindType = 107
	FIND_MY_STRUCTURES              FindType = 108
	FIND_HOSTILE_STRUCTURES         FindType = 109
	FIND_FLAGS                      FindType = 110
	FIND_CONSTRUCTION_SITES         FindType = 111
	FIND_MY_SPAWNS                  FindType = 112
	FIND_HOSTILE_SPAWNS             FindType = 113
	FIND_MY_CONSTRUCTION_SITES      FindType = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FindType = 115
	FIND_MINERALS                   FindType = 116
	FIND_NUKES                      FindType = 117
	FIND_TOMBSTONES                 FindType = 118
	FIND_POWER_CREEPS               FindType = 119
	FIND_MY_POWER_CREEPS            FindType = 120
	FIND_HOSTILE_POWER_CREEPS       FindType = 121
	FIND_DEPOSITS                   FindType = 122
	FIND_RUINS                      FindType = 123
)

type AlgorithmType string

const (
	ALGORITHM_ASTAR    AlgorithmType = "astar"
	ALGORITHM_DIJKSTRA AlgorithmType = "dijkstra"
)

type DirectionType int

const (
	TOP          DirectionType = 1
	TOP_RIGHT    DirectionType = 2
	RIGHT        DirectionType = 3
	BOTTOM_RIGHT DirectionType = 4
	BOTTOM       DirectionType = 5
	BOTTOM_LEFT  DirectionType = 6
	LEFT         DirectionType = 7
	TOP_LEFT     DirectionType = 8
)

type StructureType string

const (
	STRUCTURE_SPAWN        StructureType = "spawn"
	STRUCTURE_EXTENSION    StructureType = "extension"
	STRUCTURE_ROAD         StructureType = "road"
	STRUCTURE_WALL         StructureType = "constructedWall"
	STRUCTURE_RAMPART      StructureType = "rampart"
	STRUCTURE_KEEPER_LAIR  StructureType = "keeperLair"
	STRUCTURE_PORTAL       StructureType = "portal"
	STRUCTURE_CONTROLLER   StructureType = "controller"
	STRUCTURE_LINK         StructureType = "link"
	STRUCTURE_STORAGE      StructureType = "storage"
	STRUCTURE_TOWER        StructureType = "tower"
	STRUCTURE_OBSERVER     StructureType = "observer"
	STRUCTURE_POWER_BANK   StructureType = "powerBank"
	STRUCTURE_POWER_SPAWN  StructureType = "powerSpawn"
	STRUCTURE_EXTRACTOR    StructureType = "extractor"
	STRUCTURE_LAB          StructureType = "lab"
	STRUCTURE_TERMINAL     StructureType = "terminal"
	STRUCTURE_CONTAINER    StructureType = "container"
	STRUCTURE_NUKER        StructureType = "nuker"
	STRUCTURE_FACTORY      StructureType = "factory"
	STRUCTURE_INVADER_CORE StructureType = "invaderCore"
)

type LookType string

const (
	LOOK_CREEPS             LookType = "creep"
	LOOK_ENERGY             LookType = "energy"
	LOOK_RESOURCES          LookType = "resource"
	LOOK_SOURCES            LookType = "source"
	LOOK_MINERALS           LookType = "mineral"
	LOOK_DEPOSITS           LookType = "deposit"
	LOOK_STRUCTURES         LookType = "structure"
	LOOK_FLAGS              LookType = "flag"
	LOOK_CONSTRUCTION_SITES LookType = "constructionSite"
	LOOK_NUKES              LookType = "nuke"
	LOOK_TERRAIN            LookType = "terrain"
	LOOK_TOMBSTONES         LookType = "tombstone"
	LOOK_POWER_CREEPS       LookType = "powerCreep"
	LOOK_RUINS              LookType = "ruin"
)
