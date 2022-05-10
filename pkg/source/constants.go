package screeps

type FindExitConstant int
type FindRoomObjectConstant int

const (
	FIND_EXIT_TOP                   FindExitConstant       = 1
	FIND_EXIT_RIGHT                 FindExitConstant       = 3
	FIND_EXIT_BOTTOM                FindExitConstant       = 5
	FIND_EXIT_LEFT                  FindExitConstant       = 7
	FIND_EXIT                       FindExitConstant       = 10
	FIND_CREEPS                     FindRoomObjectConstant = 101
	FIND_MY_CREEPS                  FindRoomObjectConstant = 102
	FIND_HOSTILE_CREEPS             FindRoomObjectConstant = 103
	FIND_SOURCES_ACTIVE             FindRoomObjectConstant = 104
	FIND_SOURCES                    FindRoomObjectConstant = 105
	FIND_DROPPED_RESOURCES          FindRoomObjectConstant = 106
	FIND_STRUCTURES                 FindRoomObjectConstant = 107
	FIND_MY_STRUCTURES              FindRoomObjectConstant = 108
	FIND_HOSTILE_STRUCTURES         FindRoomObjectConstant = 109
	FIND_FLAGS                      FindRoomObjectConstant = 110
	FIND_CONSTRUCTION_SITES         FindRoomObjectConstant = 111
	FIND_MY_SPAWNS                  FindRoomObjectConstant = 112
	FIND_HOSTILE_SPAWNS             FindRoomObjectConstant = 113
	FIND_MY_CONSTRUCTION_SITES      FindRoomObjectConstant = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FindRoomObjectConstant = 115
	FIND_MINERALS                   FindRoomObjectConstant = 116
	FIND_NUKES                      FindRoomObjectConstant = 117
	FIND_TOMBSTONES                 FindRoomObjectConstant = 118
	FIND_POWER_CREEPS               FindRoomObjectConstant = 119
	FIND_MY_POWER_CREEPS            FindRoomObjectConstant = 120
	FIND_HOSTILE_POWER_CREEPS       FindRoomObjectConstant = 121
	FIND_DEPOSITS                   FindRoomObjectConstant = 122
	FIND_RUINS                      FindRoomObjectConstant = 123
)

type DirectionConstant int

const (
	TOP          DirectionConstant = 1
	TOP_RIGHT    DirectionConstant = 2
	RIGHT        DirectionConstant = 3
	BOTTOM_RIGHT DirectionConstant = 4
	BOTTOM       DirectionConstant = 5
	BOTTOM_LEFT  DirectionConstant = 6
	LEFT         DirectionConstant = 7
	TOP_LEFT     DirectionConstant = 8
)



type LookConstant string

const (
	LOOK_CREEPS             LookConstant = "creep"
	LOOK_ENERGY             LookConstant = "energy"
	LOOK_RESOURCES          LookConstant = "resource"
	LOOK_SOURCES            LookConstant = "source"
	LOOK_MINERALS           LookConstant = "mineral"
	LOOK_DEPOSITS           LookConstant = "deposit"
	LOOK_STRUCTURES         LookConstant = "structure"
	LOOK_FLAGS              LookConstant = "flag"
	LOOK_CONSTRUCTION_SITES LookConstant = "constructionSite"
	LOOK_NUKES              LookConstant = "nuke"
	//LOOK_TERRAIN          LookConstant = "terrain" // use Game.map
	LOOK_TOMBSTONES   LookConstant = "tombstone"
	LOOK_POWER_CREEPS LookConstant = "powerCreep"
	LOOK_RUINS        LookConstant = "ruin"
)

var OBSTACLE_OBJECT_TYPES = []string{
	"spawn",
	"creep",
	"powerCreep",
	"source",
	"mineral",
	"deposit",
	"controller",
	"constructedWall",
	"extension",
	"link",
	"storage",
	"tower",
	"observer",
	"powerSpawn",
	"powerBank",
	"lab",
	"terminal",
	"nuker",
	"factory",
	"invaderCore",
}

type BodyPartConstant string

const (
	MOVE          BodyPartConstant = "move"
	WORK          BodyPartConstant = "work"
	CARRY         BodyPartConstant = "carry"
	ATTACK        BodyPartConstant = "attack"
	RANGED_ATTACK BodyPartConstant = "ranged_attack"
	TOUGH         BodyPartConstant = "tough"
	HEAL          BodyPartConstant = "heal"
	CLAIM         BodyPartConstant = "claim"
)

var BODYPART_COST = map[string]int{
	"move":          50,
	"work":          100,
	"attack":        80,
	"carry":         50,
	"heal":          250,
	"ranged_attack": 150,
	"tough":         10,
	"claim":         600,
}

const CREEP_LIFE_TIME = 1500
const CREEP_CLAIM_LIFE_TIME = 600
const CREEP_CORPSE_RATE = 0.2
const CREEP_PART_MAX_ENERGY = 125

const CARRY_CAPACITY = 50
const HARVEST_POWER = 2
const HARVEST_MINERAL_POWER = 1
const HARVEST_DEPOSIT_POWER = 1
const REPAIR_POWER = 100
const DISMANTLE_POWER = 50
const BUILD_POWER = 5
const ATTACK_POWER = 30
const UPGRADE_CONTROLLER_POWER = 1
const RANGED_ATTACK_POWER = 10
const HEAL_POWER = 12
const RANGED_HEAL_POWER = 4
const REPAIR_COST = 0.01
const DISMANTLE_COST = 0.005

const RAMPART_DECAY_AMOUNT = 300
const RAMPART_DECAY_TIME = 100
const RAMPART_HITS = 1

var RAMPART_HITS_MAX = map[int]int{2: 300000, 3: 1000000, 4: 3000000, 5: 10000000, 6: 30000000, 7: 100000000, 8: 300000000}

const ENERGY_REGEN_TIME = 300
const ENERGY_DECAY = 1000

const SPAWN_HITS = 5000
const SPAWN_ENERGY_START = 300
const SPAWN_ENERGY_CAPACITY = 300
const CREEP_SPAWN_TIME = 3
const SPAWN_RENEW_RATIO = 1.2

const SOURCE_ENERGY_CAPACITY = 3000
const SOURCE_ENERGY_NEUTRAL_CAPACITY = 1500
const SOURCE_ENERGY_KEEPER_CAPACITY = 4000

const WALL_HITS = 1
const WALL_HITS_MAX = 300000000

const EXTENSION_HITS = 1000

var EXTENSION_ENERGY_CAPACITY = map[int]int{0: 50, 1: 50, 2: 50, 3: 50, 4: 50, 5: 50, 6: 50, 7: 100, 8: 200}

const ROAD_HITS = 5000
const ROAD_WEAROUT = 1
const ROAD_WEAROUT_POWER_CREEP = 100
const ROAD_DECAY_AMOUNT = 100
const ROAD_DECAY_TIME = 1000

const LINK_HITS = 1000
const LINK_HITS_MAX = 1000
const LINK_CAPACITY = 800
const LINK_COOLDOWN = 1
const LINK_LOSS_RATIO = 0.03

const STORAGE_CAPACITY = 1000000
const STORAGE_HITS = 10000

type StructureConstant string

const (
	STRUCTURE_SPAWN        StructureConstant = "spawn"
	STRUCTURE_EXTENSION    StructureConstant = "extension"
	STRUCTURE_ROAD         StructureConstant = "road"
	STRUCTURE_WALL         StructureConstant = "constructedWall"
	STRUCTURE_RAMPART      StructureConstant = "rampart"
	STRUCTURE_KEEPER_LAIR  StructureConstant = "keeperLair"
	STRUCTURE_PORTAL       StructureConstant = "portal"
	STRUCTURE_CONTROLLER   StructureConstant = "controller"
	STRUCTURE_LINK         StructureConstant = "link"
	STRUCTURE_STORAGE      StructureConstant = "storage"
	STRUCTURE_TOWER        StructureConstant = "tower"
	STRUCTURE_OBSERVER     StructureConstant = "observer"
	STRUCTURE_POWER_BANK   StructureConstant = "powerBank"
	STRUCTURE_POWER_SPAWN  StructureConstant = "powerSpawn"
	STRUCTURE_EXTRACTOR    StructureConstant = "extractor"
	STRUCTURE_LAB          StructureConstant = "lab"
	STRUCTURE_TERMINAL     StructureConstant = "terminal"
	STRUCTURE_CONTAINER    StructureConstant = "container"
	STRUCTURE_NUKER        StructureConstant = "nuker"
	STRUCTURE_FACTORY      StructureConstant = "factory"
	STRUCTURE_INVADER_CORE StructureConstant = "invaderCore"
)

var CONSTRUCTION_COST = map[string]int{
	"spawn":           15000,
	"extension":       3000,
	"road":            300,
	"constructedWall": 1,
	"rampart":         1,
	"link":            5000,
	"storage":         30000,
	"tower":           5000,
	"observer":        8000,
	"powerSpawn":      100000,
	"extractor":       5000,
	"lab":             50000,
	"terminal":        100000,
	"container":       5000,
	"nuker":           100000,
	"factory":         100000,
}

const CONSTRUCTION_COST_ROAD_SWAMP_RATIO = 5
const CONSTRUCTION_COST_ROAD_WALL_RATIO = 150

var CONTROLLER_LEVELS = map[int]int{1: 200, 2: 45000, 3: 135000, 4: 405000, 5: 1215000, 6: 3645000, 7: 10935000}
var CONTROLLER_STRUCTURES = map[string]map[int]int{
	"spawn":           {0: 0, 1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 2, 8: 3},
	"extension":       {0: 0, 1: 0, 2: 5, 3: 10, 4: 20, 5: 30, 6: 40, 7: 50, 8: 60},
	"link":            {1: 0, 2: 0, 3: 0, 4: 0, 5: 2, 6: 3, 7: 4, 8: 6},
	"road":            {0: 2500, 1: 2500, 2: 2500, 3: 2500, 4: 2500, 5: 2500, 6: 2500, 7: 2500, 8: 2500},
	"constructedWall": {1: 0, 2: 2500, 3: 2500, 4: 2500, 5: 2500, 6: 2500, 7: 2500, 8: 2500},
	"rampart":         {1: 0, 2: 2500, 3: 2500, 4: 2500, 5: 2500, 6: 2500, 7: 2500, 8: 2500},
	"storage":         {1: 0, 2: 0, 3: 0, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1},
	"tower":           {1: 0, 2: 0, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 6},
	"observer":        {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 1},
	"powerSpawn":      {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 1},
	"extractor":       {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 1, 7: 1, 8: 1},
	"terminal":        {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 1, 7: 1, 8: 1},
	"lab":             {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 3, 7: 6, 8: 10},
	"container":       {0: 5, 1: 5, 2: 5, 3: 5, 4: 5, 5: 5, 6: 5, 7: 5, 8: 5},
	"nuker":           {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 1},
	"factory":         {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 1, 8: 1},
}
var CONTROLLER_DOWNGRADE = map[int]int{1: 20000, 2: 10000, 3: 20000, 4: 40000, 5: 80000, 6: 120000, 7: 150000, 8: 200000}

const CONTROLLER_DOWNGRADE_RESTORE = 100
const CONTROLLER_DOWNGRADE_SAFEMODE_THRESHOLD = 5000
const CONTROLLER_CLAIM_DOWNGRADE = 300
const CONTROLLER_RESERVE = 1
const CONTROLLER_RESERVE_MAX = 5000
const CONTROLLER_MAX_UPGRADE_PER_TICK = 15
const CONTROLLER_ATTACK_BLOCKED_UPGRADE = 1000
const CONTROLLER_NUKE_BLOCKED_UPGRADE = 200

const SAFE_MODE_DURATION = 20000
const SAFE_MODE_COOLDOWN = 50000
const SAFE_MODE_COST = 1000

const TOWER_HITS = 3000
const TOWER_CAPACITY = 1000
const TOWER_ENERGY_COST = 10
const TOWER_POWER_ATTACK = 600
const TOWER_POWER_HEAL = 400
const TOWER_POWER_REPAIR = 800
const TOWER_OPTIMAL_RANGE = 5
const TOWER_FALLOFF_RANGE = 20
const TOWER_FALLOFF = 0.75

const OBSERVER_HITS = 500
const OBSERVER_RANGE = 10

const POWER_BANK_HITS = 2000000
const POWER_BANK_CAPACITY_MAX = 5000
const POWER_BANK_CAPACITY_MIN = 500
const POWER_BANK_CAPACITY_CRIT = 0.3
const POWER_BANK_DECAY = 5000
const POWER_BANK_HIT_BACK = 0.5

const POWER_SPAWN_HITS = 5000
const POWER_SPAWN_ENERGY_CAPACITY = 5000
const POWER_SPAWN_POWER_CAPACITY = 100
const POWER_SPAWN_ENERGY_RATIO = 50

const EXTRACTOR_HITS = 500
const EXTRACTOR_COOLDOWN = 5

const LAB_HITS = 500
const LAB_MINERAL_CAPACITY = 3000
const LAB_ENERGY_CAPACITY = 2000
const LAB_BOOST_ENERGY = 20
const LAB_BOOST_MINERAL = 30
const LAB_REACTION_AMOUNT = 5
const LAB_UNBOOST_ENERGY = 0
const LAB_UNBOOST_MINERAL = 15

const GCL_POW = 2.4
const GCL_MULTIPLY = 1000000
const GCL_NOVICE = 3

const TERRAIN_MASK_WALL = 1
const TERRAIN_MASK_SWAMP = 2
const TERRAIN_MASK_LAVA = 4

const MAX_CONSTRUCTION_SITES = 100
const MAX_CREEP_SIZE = 50

const MINERAL_REGEN_TIME = 50000

var MINERAL_MIN_AMOUNT = map[string]int{
	"H": 35000,
	"O": 35000,
	"L": 35000,
	"K": 35000,
	"Z": 35000,
	"U": 35000,
	"X": 35000,
}

const MINERAL_RANDOM_FACTOR = 2

var MINERAL_DENSITY = map[int]int{
	1: 15000,
	2: 35000,
	3: 70000,
	4: 100000,
}
var MINERAL_DENSITY_PROBABILITY = map[int]float32{
	1: 0.1,
	2: 0.5,
	3: 0.9,
	4: 1.0,
}

const MINERAL_DENSITY_CHANGE = 0.05

const DENSITY_LOW = 1
const DENSITY_MODERATE = 2
const DENSITY_HIGH = 3
const DENSITY_ULTRA = 4

const DEPOSIT_EXHAUST_MULTIPLY = 0.001
const DEPOSIT_EXHAUST_POW = 1.2
const DEPOSIT_DECAY_TIME = 50000

const TERMINAL_CAPACITY = 300000
const TERMINAL_HITS = 3000
const TERMINAL_SEND_COST = 0.1
const TERMINAL_MIN_SEND = 100
const TERMINAL_COOLDOWN = 10

const CONTAINER_HITS = 250000
const CONTAINER_CAPACITY = 2000
const CONTAINER_DECAY = 5000
const CONTAINER_DECAY_TIME = 100
const CONTAINER_DECAY_TIME_OWNED = 500

const NUKER_HITS = 1000
const NUKER_COOLDOWN = 100000
const NUKER_ENERGY_CAPACITY = 300000
const NUKER_GHODIUM_CAPACITY = 5000
const NUKE_LAND_TIME = 50000
const NUKE_RANGE = 10

var NUKE_DAMAGE = map[int]int{
	0: 10000000,
	2: 5000000,
}

const FACTORY_HITS = 1000
const FACTORY_CAPACITY = 50000

const TOMBSTONE_DECAY_PER_PART = 5
const TOMBSTONE_DECAY_POWER_CREEP = 500

const RUIN_DECAY = 500

var RUIN_DECAY_STRUCTURES = map[string]int{
	"powerBank": 10,
}

const PORTAL_DECAY = 30000

const ORDER_SELL = "sell"
const ORDER_BUY = "const buy"

const MARKET_FEE = 0.05

const MARKET_MAX_ORDERS = 300
const MARKET_ORDER_LIFE_TIME = 1000 * 60 * 60 * 24 * 30

const FLAGS_LIMIT = 10000



const PIXEL_CPU_COST = 10000

type ResourceConstant string

var (
	RESOURCE_ENERGY ResourceConstant = "energy"
	RESOURCE_POWER  ResourceConstant = "power"

	RESOURCE_HYDROGEN  ResourceConstant = "H"
	RESOURCE_OXYGEN    ResourceConstant = "O"
	RESOURCE_UTRIUM    ResourceConstant = "U"
	RESOURCE_LEMERGIUM ResourceConstant = "L"
	RESOURCE_KEANIUM   ResourceConstant = "K"
	RESOURCE_ZYNTHIUM  ResourceConstant = "Z"
	RESOURCE_CATALYST  ResourceConstant = "X"
	RESOURCE_GHODIUM   ResourceConstant = "G"

	RESOURCE_SILICON ResourceConstant = "silicon"
	RESOURCE_METAL   ResourceConstant = "metal"
	RESOURCE_BIOMASS ResourceConstant = "biomass"
	RESOURCE_MIST    ResourceConstant = "mist"

	RESOURCE_HYDROXIDE        ResourceConstant = "OH"
	RESOURCE_ZYNTHIUM_KEANITE ResourceConstant = "ZK"
	RESOURCE_UTRIUM_LEMERGITE ResourceConstant = "UL"

	RESOURCE_UTRIUM_HYDRIDE    ResourceConstant = "UH"
	RESOURCE_UTRIUM_OXIDE      ResourceConstant = "UO"
	RESOURCE_KEANIUM_HYDRIDE   ResourceConstant = "KH"
	RESOURCE_KEANIUM_OXIDE     ResourceConstant = "KO"
	RESOURCE_LEMERGIUM_HYDRIDE ResourceConstant = "LH"
	RESOURCE_LEMERGIUM_OXIDE   ResourceConstant = "LO"
	RESOURCE_ZYNTHIUM_HYDRIDE  ResourceConstant = "ZH"
	RESOURCE_ZYNTHIUM_OXIDE    ResourceConstant = "ZO"
	RESOURCE_GHODIUM_HYDRIDE   ResourceConstant = "GH"
	RESOURCE_GHODIUM_OXIDE     ResourceConstant = "GO"

	RESOURCE_UTRIUM_ACID        ResourceConstant = "UH2O"
	RESOURCE_UTRIUM_ALKALIDE    ResourceConstant = "UHO2"
	RESOURCE_KEANIUM_ACID       ResourceConstant = "KH2O"
	RESOURCE_KEANIUM_ALKALIDE   ResourceConstant = "KHO2"
	RESOURCE_LEMERGIUM_ACID     ResourceConstant = "LH2O"
	RESOURCE_LEMERGIUM_ALKALIDE ResourceConstant = "LHO2"
	RESOURCE_ZYNTHIUM_ACID      ResourceConstant = "ZH2O"
	RESOURCE_ZYNTHIUM_ALKALIDE  ResourceConstant = "ZHO2"
	RESOURCE_GHODIUM_ACID       ResourceConstant = "GH2O"
	RESOURCE_GHODIUM_ALKALIDE   ResourceConstant = "GHO2"

	RESOURCE_CATALYZED_UTRIUM_ACID        ResourceConstant = "XUH2O"
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE    ResourceConstant = "XUHO2"
	RESOURCE_CATALYZED_KEANIUM_ACID       ResourceConstant = "XKH2O"
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE   ResourceConstant = "XKHO2"
	RESOURCE_CATALYZED_LEMERGIUM_ACID     ResourceConstant = "XLH2O"
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE ResourceConstant = "XLHO2"
	RESOURCE_CATALYZED_ZYNTHIUM_ACID      ResourceConstant = "XZH2O"
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE  ResourceConstant = "XZHO2"
	RESOURCE_CATALYZED_GHODIUM_ACID       ResourceConstant = "XGH2O"
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE   ResourceConstant = "XGHO2"

	RESOURCE_OPS ResourceConstant = "ops"

	RESOURCE_UTRIUM_BAR    ResourceConstant = "utrium_bar"
	RESOURCE_LEMERGIUM_BAR ResourceConstant = "lemergium_bar"
	RESOURCE_ZYNTHIUM_BAR  ResourceConstant = "zynthium_bar"
	RESOURCE_KEANIUM_BAR   ResourceConstant = "keanium_bar"
	RESOURCE_GHODIUM_MELT  ResourceConstant = "ghodium_melt"
	RESOURCE_OXIDANT       ResourceConstant = "oxidant"
	RESOURCE_REDUCTANT     ResourceConstant = "reductant"
	RESOURCE_PURIFIER      ResourceConstant = "purifier"
	RESOURCE_BATTERY       ResourceConstant = "battery"

	RESOURCE_COMPOSITE ResourceConstant = "composite"
	RESOURCE_CRYSTAL   ResourceConstant = "crystal"
	RESOURCE_LIQUID    ResourceConstant = "liquid"

	RESOURCE_WIRE       ResourceConstant = "wire"
	RESOURCE_SWITCH     ResourceConstant = "switch"
	RESOURCE_TRANSISTOR ResourceConstant = "transistor"
	RESOURCE_MICROCHIP  ResourceConstant = "microchip"
	RESOURCE_CIRCUIT    ResourceConstant = "circuit"
	RESOURCE_DEVICE     ResourceConstant = "device"

	RESOURCE_CELL     ResourceConstant = "cell"
	RESOURCE_PHLEGM   ResourceConstant = "phlegm"
	RESOURCE_TISSUE   ResourceConstant = "tissue"
	RESOURCE_MUSCLE   ResourceConstant = "muscle"
	RESOURCE_ORGANOID ResourceConstant = "organoid"
	RESOURCE_ORGANISM ResourceConstant = "organism"

	RESOURCE_ALLOY      ResourceConstant = "alloy"
	RESOURCE_TUBE       ResourceConstant = "tube"
	RESOURCE_FIXTURES   ResourceConstant = "fixtures"
	RESOURCE_FRAME      ResourceConstant = "frame"
	RESOURCE_HYDRAULICS ResourceConstant = "hydraulics"
	RESOURCE_MACHINE    ResourceConstant = "machine"

	RESOURCE_CONDENSATE  ResourceConstant = "condensate"
	RESOURCE_CONCENTRATE ResourceConstant = "concentrate"
	RESOURCE_EXTRACT     ResourceConstant = "extract"
	RESOURCE_SPIRIT      ResourceConstant = "spirit"
	RESOURCE_EMANATION   ResourceConstant = "emanation"
	RESOURCE_ESSENCE     ResourceConstant = "essence"
)

/*
REACTIONS = {
	H = {
		O = "OH"
		L = "LH"
		K = "KH"
		U = "UH"
		Z = "ZH"
		G = "GH"
	}
	O = {
		H = "OH"
		L = "LO"
		K = "KO"
		U = "UO"
		Z = "ZO"
		G = "GO"
	}
	Z = {
		K = "ZK"
		H = "ZH"
		O = "ZO"
	}
	L = {
		U = "UL"
		H = "LH"
		O = "LO"
	}
	K = {
		Z = "ZK"
		H = "KH"
		O = "KO"
	}
	G = {
		H = "GH"
		O = "GO"
	}
	U = {
		L = "UL"
		H = "UH"
		O = "UO"
	}
	OH = {
		UH = "UH2O"
		UO = "UHO2"
		ZH = "ZH2O"
		ZO = "ZHO2"
		KH = "KH2O"
		KO = "KHO2"
		LH = "LH2O"
		LO = "LHO2"
		GH = "GH2O"
		GO = "GHO2"
	}
	X = {
		UH2O = "XUH2O"
		UHO2 = "XUHO2"
		LH2O = "XLH2O"
		LHO2 = "XLHO2"
		KH2O = "XKH2O"
		KHO2 = "XKHO2"
		ZH2O = "XZH2O"
		ZHO2 = "XZHO2"
		GH2O = "XGH2O"
		GHO2 = "XGHO2"
	}
	ZK = {
		UL = "G"
	}
	UL = {
		ZK = "G"
	}
	LH = {
		OH = "LH2O"
	}
	ZH = {
		OH = "ZH2O"
	}
	GH = {
		OH = "GH2O"
	}
	KH = {
		OH = "KH2O"
	}
	UH = {
		OH = "UH2O"
	}
	LO = {
		OH = "LHO2"
	}
	ZO = {
		OH = "ZHO2"
	}
	KO = {
		OH = "KHO2"
	}
	UO = {
		OH = "UHO2"
	}
	GO = {
		OH = "GHO2"
	}
	LH2O = {
		X = "XLH2O"
	}
	KH2O = {
		X = "XKH2O"
	}
	ZH2O = {
		X = "XZH2O"
	}
	UH2O = {
		X = "XUH2O"
	}
	GH2O = {
		X = "XGH2O"
	}
	LHO2 = {
		X = "XLHO2"
	}
	UHO2 = {
		X = "XUHO2"
	}
	KHO2 = {
		X = "XKHO2"
	}
	ZHO2 = {
		X = "XZHO2"
	}
	GHO2 = {
		X = "XGHO2"
	}
}

BOOSTS = {
	work = {
		UO = {
			harvest = 3
		}
		UHO2 = {
			harvest = 5
		}
		XUHO2 = {
			harvest = 7
		}
		LH = {
			build = 1.5
			repair = 1.5
		}
		LH2O = {
			build = 1.8
			repair = 1.8
		}
		XLH2O = {
			build = 2
			repair = 2
		}
		ZH = {
			dismantle = 2
		}
		ZH2O = {
			dismantle = 3
		}
		XZH2O = {
			dismantle = 4
		}
		GH = {
			upgradeController = 1.5
		}
		GH2O = {
			upgradeController = 1.8
		}
		XGH2O = {
			upgradeController = 2
		}
	}
	attack = {
		UH = {
			attack = 2
		}
		UH2O = {
			attack = 3
		}
		XUH2O = {
			attack = 4
		}
	}
	ranged_attack = {
		KO = {
			rangedAttack = 2
			rangedMassAttack = 2
		}
		KHO2 = {
			rangedAttack = 3
			rangedMassAttack = 3
		}
		XKHO2 = {
			rangedAttack = 4
			rangedMassAttack = 4
		}
	}
	heal = {
		LO = {
			heal = 2
			rangedHeal = 2
		}
		LHO2 = {
			heal = 3
			rangedHeal = 3
		}
		XLHO2 = {
			heal = 4
			rangedHeal = 4
		}
	}
	carry = {
		KH = {
			capacity = 2
		}
		KH2O = {
			capacity = 3
		}
		XKH2O = {
			capacity = 4
		}
	}
	move = {
		ZO = {
			fatigue = 2
		}
		ZHO2 = {
			fatigue = 3
		}
		XZHO2 = {
			fatigue = 4
		}
	}
	tough = {
		GO = {
			damage = .7
		}
		GHO2 = {
			damage = .5
		}
		XGHO2 = {
			damage = .3
		}
	}
}
*/

var REACTION_TIME = map[string]int{
	"OH":    20,
	"ZK":    5,
	"UL":    5,
	"G":     5,
	"UH":    10,
	"UH2O":  5,
	"XUH2O": 60,
	"UO":    10,
	"UHO2":  5,
	"XUHO2": 60,
	"KH":    10,
	"KH2O":  5,
	"XKH2O": 60,
	"KO":    10,
	"KHO2":  5,
	"XKHO2": 60,
	"LH":    15,
	"LH2O":  10,
	"XLH2O": 65,
	"LO":    10,
	"LHO2":  5,
	"XLHO2": 60,
	"ZH":    20,
	"ZH2O":  40,
	"XZH2O": 160,
	"ZO":    10,
	"ZHO2":  5,
	"XZHO2": 60,
	"GH":    10,
	"GH2O":  15,
	"XGH2O": 80,
	"GO":    10,
	"GHO2":  30,
	"XGHO2": 150,
}

const PORTAL_UNSTABLE = 10 * 24 * 3600 * 1000
const PORTAL_MIN_TIMEOUT = 12 * 24 * 3600 * 1000
const PORTAL_MAX_TIMEOUT = 22 * 24 * 3600 * 1000

const POWER_BANK_RESPAWN_TIME = 50000

const INVADERS_ENERGY_GOAL = 100000

// SIGN_NOVICE_AREA and SIGN_RESPAWN_AREA constants are deprecated please use SIGN_PLANNED_AREA instead
const SIGN_NOVICE_AREA = "A new Novice or Respawn Area is being planned somewhere in this sector. Please make sure all important rooms are reserved."
const SIGN_RESPAWN_AREA = "A new Novice or Respawn Area is being planned somewhere in this sector. Please make sure all important rooms are reserved."
const SIGN_PLANNED_AREA = "A new Novice or Respawn Area is being planned somewhere in this sector. Please make sure all important rooms are reserved."

const EVENT_ATTACK = 1
const EVENT_OBJECT_DESTROYED = 2
const EVENT_ATTACK_CONTROLLER = 3
const EVENT_BUILD = 4
const EVENT_HARVEST = 5
const EVENT_HEAL = 6
const EVENT_REPAIR = 7
const EVENT_RESERVE_CONTROLLER = 8
const EVENT_UPGRADE_CONTROLLER = 9
const EVENT_EXIT = 10
const EVENT_POWER = 11
const EVENT_TRANSFER = 12

const EVENT_ATTACK_TYPE_MELEE = 1
const EVENT_ATTACK_TYPE_RANGED = 2
const EVENT_ATTACK_TYPE_RANGED_MASS = 3
const EVENT_ATTACK_TYPE_DISMANTLE = 4
const EVENT_ATTACK_TYPE_HIT_BACK = 5
const EVENT_ATTACK_TYPE_NUKE = 6

const EVENT_HEAL_TYPE_MELEE = 1
const EVENT_HEAL_TYPE_RANGED = 2

const POWER_LEVEL_MULTIPLY = 1000
const POWER_LEVEL_POW = 2
const POWER_CREEP_SPAWN_COOLDOWN = 8 * 3600 * 1000
const POWER_CREEP_DELETE_COOLDOWN = 24 * 3600 * 1000
const POWER_CREEP_MAX_LEVEL = 25
const POWER_CREEP_LIFE_TIME = 5000

type PowerClassConstant string

var POWER_CLASS = struct {
	OPERATOR PowerClassConstant
}{
	OPERATOR: "operator",
}

const INVADER_CORE_HITS = 100000

var INVADER_CORE_CREEP_SPAWN_TIME = map[int]int{0: 0, 1: 0, 2: 6, 3: 3, 4: 2, 5: 1}
var INVADER_CORE_EXPAND_TIME = map[int]int{1: 4000, 2: 3500, 3: 3000, 4: 2500, 5: 2000}

const INVADER_CORE_CONTROLLER_POWER = 2
const INVADER_CORE_CONTROLLER_DOWNGRADE = 5000

var STRONGHOLD_RAMPART_HITS = map[int]int{0: 0, 1: 100000, 2: 200000, 3: 500000, 4: 1000000, 5: 2000000}

const STRONGHOLD_DECAY_TICKS = 75000

var BODYPARTS_ALL = []BodyPartConstant{
	MOVE,
	WORK,
	CARRY,
	ATTACK,
	RANGED_ATTACK,
	TOUGH,
	HEAL,
	CLAIM,
}
var RESOURCES_ALL = []ResourceConstant{
	RESOURCE_ENERGY,
	RESOURCE_POWER,

	RESOURCE_HYDROGEN,
	RESOURCE_OXYGEN,
	RESOURCE_UTRIUM,
	RESOURCE_KEANIUM,
	RESOURCE_LEMERGIUM,
	RESOURCE_ZYNTHIUM,
	RESOURCE_CATALYST,
	RESOURCE_GHODIUM,

	RESOURCE_HYDROXIDE,
	RESOURCE_ZYNTHIUM_KEANITE,
	RESOURCE_UTRIUM_LEMERGITE,

	RESOURCE_UTRIUM_HYDRIDE,
	RESOURCE_UTRIUM_OXIDE,
	RESOURCE_KEANIUM_HYDRIDE,
	RESOURCE_KEANIUM_OXIDE,
	RESOURCE_LEMERGIUM_HYDRIDE,
	RESOURCE_LEMERGIUM_OXIDE,
	RESOURCE_ZYNTHIUM_HYDRIDE,
	RESOURCE_ZYNTHIUM_OXIDE,
	RESOURCE_GHODIUM_HYDRIDE,
	RESOURCE_GHODIUM_OXIDE,

	RESOURCE_UTRIUM_ACID,
	RESOURCE_UTRIUM_ALKALIDE,
	RESOURCE_KEANIUM_ACID,
	RESOURCE_KEANIUM_ALKALIDE,
	RESOURCE_LEMERGIUM_ACID,
	RESOURCE_LEMERGIUM_ALKALIDE,
	RESOURCE_ZYNTHIUM_ACID,
	RESOURCE_ZYNTHIUM_ALKALIDE,
	RESOURCE_GHODIUM_ACID,
	RESOURCE_GHODIUM_ALKALIDE,

	RESOURCE_CATALYZED_UTRIUM_ACID,
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE,
	RESOURCE_CATALYZED_KEANIUM_ACID,
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE,
	RESOURCE_CATALYZED_LEMERGIUM_ACID,
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE,
	RESOURCE_CATALYZED_ZYNTHIUM_ACID,
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE,
	RESOURCE_CATALYZED_GHODIUM_ACID,
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE,

	RESOURCE_OPS,

	RESOURCE_SILICON,
	RESOURCE_METAL,
	RESOURCE_BIOMASS,
	RESOURCE_MIST,

	RESOURCE_UTRIUM_BAR,
	RESOURCE_LEMERGIUM_BAR,
	RESOURCE_ZYNTHIUM_BAR,
	RESOURCE_KEANIUM_BAR,
	RESOURCE_GHODIUM_MELT,
	RESOURCE_OXIDANT,
	RESOURCE_REDUCTANT,
	RESOURCE_PURIFIER,
	RESOURCE_BATTERY,
	RESOURCE_COMPOSITE,
	RESOURCE_CRYSTAL,
	RESOURCE_LIQUID,

	RESOURCE_WIRE,
	RESOURCE_SWITCH,
	RESOURCE_TRANSISTOR,
	RESOURCE_MICROCHIP,
	RESOURCE_CIRCUIT,
	RESOURCE_DEVICE,

	RESOURCE_CELL,
	RESOURCE_PHLEGM,
	RESOURCE_TISSUE,
	RESOURCE_MUSCLE,
	RESOURCE_ORGANOID,
	RESOURCE_ORGANISM,

	RESOURCE_ALLOY,
	RESOURCE_TUBE,
	RESOURCE_FIXTURES,
	RESOURCE_FRAME,
	RESOURCE_HYDRAULICS,
	RESOURCE_MACHINE,

	RESOURCE_CONDENSATE,
	RESOURCE_CONCENTRATE,
	RESOURCE_EXTRACT,
	RESOURCE_SPIRIT,
	RESOURCE_EMANATION,
	RESOURCE_ESSENCE,
}
var COLORS_ALL = []int{
	COLOR_RED,
	COLOR_PURPLE,
	COLOR_BLUE,
	COLOR_CYAN,
	COLOR_GREEN,
	COLOR_YELLOW,
	COLOR_ORANGE,
	COLOR_BROWN,
	COLOR_GREY,
	COLOR_WHITE,
}

