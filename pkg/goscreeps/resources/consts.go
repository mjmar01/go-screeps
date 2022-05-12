package resources

type (
	AccountResourceConst string
	AlgorithmConst       string
	StructureConst       string
	LookConst            string
	DirectionConst       int
	EffectTypeConst      int
	ColorConst           int
	FindObjectConst      int
	FindPosConst         int
	TerrainConst         int
)

const (
	SUBSCRIPTION_TOKEN AccountResourceConst = "token"
	CPU_UNLOCK         AccountResourceConst = "cpuUnlock"
	PIXEL              AccountResourceConst = "pixel"
	ACCESS_KEY         AccountResourceConst = "accessKey"
)

const (
	PWR_GENERATE_OPS       EffectTypeConst = 1
	PWR_OPERATE_SPAWN      EffectTypeConst = 2
	PWR_OPERATE_TOWER      EffectTypeConst = 3
	PWR_OPERATE_STORAGE    EffectTypeConst = 4
	PWR_OPERATE_LAB        EffectTypeConst = 5
	PWR_OPERATE_EXTENSION  EffectTypeConst = 6
	PWR_OPERATE_OBSERVER   EffectTypeConst = 7
	PWR_OPERATE_TERMINAL   EffectTypeConst = 8
	PWR_DISRUPT_SPAWN      EffectTypeConst = 9
	PWR_DISRUPT_TOWER      EffectTypeConst = 10
	PWR_DISRUPT_SOURCE     EffectTypeConst = 11
	PWR_SHIELD             EffectTypeConst = 12
	PWR_REGEN_SOURCE       EffectTypeConst = 13
	PWR_REGEN_MINERAL      EffectTypeConst = 14
	PWR_DISRUPT_TERMINAL   EffectTypeConst = 15
	PWR_OPERATE_POWER      EffectTypeConst = 16
	PWR_FORTIFY            EffectTypeConst = 17
	PWR_OPERATE_CONTROLLER EffectTypeConst = 18
	PWR_OPERATE_FACTORY    EffectTypeConst = 19
	EFFECT_INVULNERABILITY EffectTypeConst = 1001
	EFFECT_COLLAPSE_TIMER  EffectTypeConst = 1002
)

const (
	COLOR_RED    ColorConst = 1
	COLOR_PURPLE ColorConst = 2
	COLOR_BLUE   ColorConst = 3
	COLOR_CYAN   ColorConst = 4
	COLOR_GREEN  ColorConst = 5
	COLOR_YELLOW ColorConst = 6
	COLOR_ORANGE ColorConst = 7
	COLOR_BROWN  ColorConst = 8
	COLOR_GREY   ColorConst = 9
	COLOR_WHITE  ColorConst = 10
)

const (
	FIND_EXIT_TOP    FindPosConst = 1
	FIND_EXIT_RIGHT  FindPosConst = 3
	FIND_EXIT_BOTTOM FindPosConst = 5
	FIND_EXIT_LEFT   FindPosConst = 7
	FIND_EXIT        FindPosConst = 10
)

const (
	FIND_CREEPS                     FindObjectConst = 101
	FIND_MY_CREEPS                  FindObjectConst = 102
	FIND_HOSTILE_CREEPS             FindObjectConst = 103
	FIND_SOURCES_ACTIVE             FindObjectConst = 104
	FIND_SOURCES                    FindObjectConst = 105
	FIND_DROPPED_RESOURCES          FindObjectConst = 106
	FIND_STRUCTURES                 FindObjectConst = 107
	FIND_MY_STRUCTURES              FindObjectConst = 108
	FIND_HOSTILE_STRUCTURES         FindObjectConst = 109
	FIND_FLAGS                      FindObjectConst = 110
	FIND_CONSTRUCTION_SITES         FindObjectConst = 111
	FIND_MY_SPAWNS                  FindObjectConst = 112
	FIND_HOSTILE_SPAWNS             FindObjectConst = 113
	FIND_MY_CONSTRUCTION_SITES      FindObjectConst = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FindObjectConst = 115
	FIND_MINERALS                   FindObjectConst = 116
	FIND_NUKES                      FindObjectConst = 117
	FIND_TOMBSTONES                 FindObjectConst = 118
	FIND_POWER_CREEPS               FindObjectConst = 119
	FIND_MY_POWER_CREEPS            FindObjectConst = 120
	FIND_HOSTILE_POWER_CREEPS       FindObjectConst = 121
	FIND_DEPOSITS                   FindObjectConst = 122
	FIND_RUINS                      FindObjectConst = 123
)

const (
	ALGORITHM_ASTAR    AlgorithmConst = "astar"
	ALGORITHM_DIJKSTRA AlgorithmConst = "dijkstra"
)

const (
	TOP          DirectionConst = 1
	TOP_RIGHT    DirectionConst = 2
	RIGHT        DirectionConst = 3
	BOTTOM_RIGHT DirectionConst = 4
	BOTTOM       DirectionConst = 5
	BOTTOM_LEFT  DirectionConst = 6
	LEFT         DirectionConst = 7
	TOP_LEFT     DirectionConst = 8
)

const (
	STRUCTURE_SPAWN        StructureConst = "spawn"
	STRUCTURE_EXTENSION    StructureConst = "extension"
	STRUCTURE_ROAD         StructureConst = "road"
	STRUCTURE_WALL         StructureConst = "constructedWall"
	STRUCTURE_RAMPART      StructureConst = "rampart"
	STRUCTURE_KEEPER_LAIR  StructureConst = "keeperLair"
	STRUCTURE_PORTAL       StructureConst = "portal"
	STRUCTURE_CONTROLLER   StructureConst = "controller"
	STRUCTURE_LINK         StructureConst = "link"
	STRUCTURE_STORAGE      StructureConst = "storage"
	STRUCTURE_TOWER        StructureConst = "tower"
	STRUCTURE_OBSERVER     StructureConst = "observer"
	STRUCTURE_POWER_BANK   StructureConst = "powerBank"
	STRUCTURE_POWER_SPAWN  StructureConst = "powerSpawn"
	STRUCTURE_EXTRACTOR    StructureConst = "extractor"
	STRUCTURE_LAB          StructureConst = "lab"
	STRUCTURE_TERMINAL     StructureConst = "terminal"
	STRUCTURE_CONTAINER    StructureConst = "container"
	STRUCTURE_NUKER        StructureConst = "nuker"
	STRUCTURE_FACTORY      StructureConst = "factory"
	STRUCTURE_INVADER_CORE StructureConst = "invaderCore"
)

const (
	LOOK_CREEPS             LookConst = "creep"
	LOOK_ENERGY             LookConst = "energy"
	LOOK_RESOURCES          LookConst = "resource"
	LOOK_SOURCES            LookConst = "source"
	LOOK_MINERALS           LookConst = "mineral"
	LOOK_DEPOSITS           LookConst = "deposit"
	LOOK_STRUCTURES         LookConst = "structure"
	LOOK_FLAGS              LookConst = "flag"
	LOOK_CONSTRUCTION_SITES LookConst = "constructionSite"
	LOOK_NUKES              LookConst = "nuke"
	LOOK_TERRAIN            LookConst = "terrain"
	LOOK_TOMBSTONES         LookConst = "tombstone"
	LOOK_POWER_CREEPS       LookConst = "powerCreep"
	LOOK_RUINS              LookConst = "ruin"
)

const (
	TERRAIN_MASK_WALL  TerrainConst = 1
	TERRAIN_MASK_SWAMP TerrainConst = 2
	TERRAIN_MASK_LAVA  TerrainConst = 4
)
