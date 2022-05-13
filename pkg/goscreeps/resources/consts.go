package resources

type (
	AccountResourceConst string
	AlgorithmConst       string
	StructureConst       string
	LookConst            string
	ResourceConst        string
	DirectionConst       int
	EffectTypeConst      int
	ColorConst           int
	FindConst            int
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
	FIND_EXIT_TOP                   FindConst = 1
	FIND_EXIT_RIGHT                 FindConst = 3
	FIND_EXIT_BOTTOM                FindConst = 5
	FIND_EXIT_LEFT                  FindConst = 7
	FIND_EXIT                       FindConst = 10
	FIND_CREEPS                     FindConst = 101
	FIND_MY_CREEPS                  FindConst = 102
	FIND_HOSTILE_CREEPS             FindConst = 103
	FIND_SOURCES_ACTIVE             FindConst = 104
	FIND_SOURCES                    FindConst = 105
	FIND_DROPPED_RESOURCES          FindConst = 106
	FIND_STRUCTURES                 FindConst = 107
	FIND_MY_STRUCTURES              FindConst = 108
	FIND_HOSTILE_STRUCTURES         FindConst = 109
	FIND_FLAGS                      FindConst = 110
	FIND_CONSTRUCTION_SITES         FindConst = 111
	FIND_MY_SPAWNS                  FindConst = 112
	FIND_HOSTILE_SPAWNS             FindConst = 113
	FIND_MY_CONSTRUCTION_SITES      FindConst = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FindConst = 115
	FIND_MINERALS                   FindConst = 116
	FIND_NUKES                      FindConst = 117
	FIND_TOMBSTONES                 FindConst = 118
	FIND_POWER_CREEPS               FindConst = 119
	FIND_MY_POWER_CREEPS            FindConst = 120
	FIND_HOSTILE_POWER_CREEPS       FindConst = 121
	FIND_DEPOSITS                   FindConst = 122
	FIND_RUINS                      FindConst = 123
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

const (
	RESOURCE_ENERGY                       ResourceConst = "energy"
	RESOURCE_POWER                        ResourceConst = "power"
	RESOURCE_HYDROGEN                     ResourceConst = "H"
	RESOURCE_OXYGEN                       ResourceConst = "O"
	RESOURCE_UTRIUM                       ResourceConst = "U"
	RESOURCE_LEMERGIUM                    ResourceConst = "L"
	RESOURCE_KEANIUM                      ResourceConst = "K"
	RESOURCE_ZYNTHIUM                     ResourceConst = "Z"
	RESOURCE_CATALYST                     ResourceConst = "X"
	RESOURCE_GHODIUM                      ResourceConst = "G"
	RESOURCE_SILICON                      ResourceConst = "silicon"
	RESOURCE_METAL                        ResourceConst = "metal"
	RESOURCE_BIOMASS                      ResourceConst = "biomass"
	RESOURCE_MIST                         ResourceConst = "mist"
	RESOURCE_HYDROXIDE                    ResourceConst = "OH"
	RESOURCE_ZYNTHIUM_KEANITE             ResourceConst = "ZK"
	RESOURCE_UTRIUM_LEMERGITE             ResourceConst = "UL"
	RESOURCE_UTRIUM_HYDRIDE               ResourceConst = "UH"
	RESOURCE_UTRIUM_OXIDE                 ResourceConst = "UO"
	RESOURCE_KEANIUM_HYDRIDE              ResourceConst = "KH"
	RESOURCE_KEANIUM_OXIDE                ResourceConst = "KO"
	RESOURCE_LEMERGIUM_HYDRIDE            ResourceConst = "LH"
	RESOURCE_LEMERGIUM_OXIDE              ResourceConst = "LO"
	RESOURCE_ZYNTHIUM_HYDRIDE             ResourceConst = "ZH"
	RESOURCE_ZYNTHIUM_OXIDE               ResourceConst = "ZO"
	RESOURCE_GHODIUM_HYDRIDE              ResourceConst = "GH"
	RESOURCE_GHODIUM_OXIDE                ResourceConst = "GO"
	RESOURCE_UTRIUM_ACID                  ResourceConst = "UH2O"
	RESOURCE_UTRIUM_ALKALIDE              ResourceConst = "UHO2"
	RESOURCE_KEANIUM_ACID                 ResourceConst = "KH2O"
	RESOURCE_KEANIUM_ALKALIDE             ResourceConst = "KHO2"
	RESOURCE_LEMERGIUM_ACID               ResourceConst = "LH2O"
	RESOURCE_LEMERGIUM_ALKALIDE           ResourceConst = "LHO2"
	RESOURCE_ZYNTHIUM_ACID                ResourceConst = "ZH2O"
	RESOURCE_ZYNTHIUM_ALKALIDE            ResourceConst = "ZHO2"
	RESOURCE_GHODIUM_ACID                 ResourceConst = "GH2O"
	RESOURCE_GHODIUM_ALKALIDE             ResourceConst = "GHO2"
	RESOURCE_CATALYZED_UTRIUM_ACID        ResourceConst = "XUH2O"
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE    ResourceConst = "XUHO2"
	RESOURCE_CATALYZED_KEANIUM_ACID       ResourceConst = "XKH2O"
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE   ResourceConst = "XKHO2"
	RESOURCE_CATALYZED_LEMERGIUM_ACID     ResourceConst = "XLH2O"
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE ResourceConst = "XLHO2"
	RESOURCE_CATALYZED_ZYNTHIUM_ACID      ResourceConst = "XZH2O"
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE  ResourceConst = "XZHO2"
	RESOURCE_CATALYZED_GHODIUM_ACID       ResourceConst = "XGH2O"
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE   ResourceConst = "XGHO2"
	RESOURCE_OPS                          ResourceConst = "ops"
	RESOURCE_UTRIUM_BAR                   ResourceConst = "utrium_bar"
	RESOURCE_LEMERGIUM_BAR                ResourceConst = "lemergium_bar"
	RESOURCE_ZYNTHIUM_BAR                 ResourceConst = "zynthium_bar"
	RESOURCE_KEANIUM_BAR                  ResourceConst = "keanium_bar"
	RESOURCE_GHODIUM_MELT                 ResourceConst = "ghodium_melt"
	RESOURCE_OXIDANT                      ResourceConst = "oxidant"
	RESOURCE_REDUCTANT                    ResourceConst = "reductant"
	RESOURCE_PURIFIER                     ResourceConst = "purifier"
	RESOURCE_BATTERY                      ResourceConst = "battery"
	RESOURCE_COMPOSITE                    ResourceConst = "composite"
	RESOURCE_CRYSTAL                      ResourceConst = "crystal"
	RESOURCE_LIQUID                       ResourceConst = "liquid"
	RESOURCE_WIRE                         ResourceConst = "wire"
	RESOURCE_SWITCH                       ResourceConst = "switch"
	RESOURCE_TRANSISTOR                   ResourceConst = "transistor"
	RESOURCE_MICROCHIP                    ResourceConst = "microchip"
	RESOURCE_CIRCUIT                      ResourceConst = "circuit"
	RESOURCE_DEVICE                       ResourceConst = "device"
	RESOURCE_CELL                         ResourceConst = "cell"
	RESOURCE_PHLEGM                       ResourceConst = "phlegm"
	RESOURCE_TISSUE                       ResourceConst = "tissue"
	RESOURCE_MUSCLE                       ResourceConst = "muscle"
	RESOURCE_ORGANOID                     ResourceConst = "organoid"
	RESOURCE_ORGANISM                     ResourceConst = "organism"
	RESOURCE_ALLOY                        ResourceConst = "alloy"
	RESOURCE_TUBE                         ResourceConst = "tube"
	RESOURCE_FIXTURES                     ResourceConst = "fixtures"
	RESOURCE_FRAME                        ResourceConst = "frame"
	RESOURCE_HYDRAULICS                   ResourceConst = "hydraulics"
	RESOURCE_MACHINE                      ResourceConst = "machine"
	RESOURCE_CONDENSATE                   ResourceConst = "condensate"
	RESOURCE_CONCENTRATE                  ResourceConst = "concentrate"
	RESOURCE_EXTRACT                      ResourceConst = "extract"
	RESOURCE_SPIRIT                       ResourceConst = "spirit"
	RESOURCE_EMANATION                    ResourceConst = "emanation"
	RESOURCE_ESSENCE                      ResourceConst = "essence"
)
