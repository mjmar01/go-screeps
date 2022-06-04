package resources

type (
	CAccountResource string
	CAlgorithm       string
	CStructure       string
	CLook            string
	CResource        string
	CBodyPart        string
	CDirection       int
	CDensity         int
	CEffect          int
	CColor           int
	CFind            int
	CTerrain         int
)

const (
	SUBSCRIPTION_TOKEN CAccountResource = "token"
	CPU_UNLOCK         CAccountResource = "cpuUnlock"
	PIXEL              CAccountResource = "pixel"
	ACCESS_KEY         CAccountResource = "accessKey"
)

const (
	PWR_GENERATE_OPS       CEffect = 1
	PWR_OPERATE_SPAWN      CEffect = 2
	PWR_OPERATE_TOWER      CEffect = 3
	PWR_OPERATE_STORAGE    CEffect = 4
	PWR_OPERATE_LAB        CEffect = 5
	PWR_OPERATE_EXTENSION  CEffect = 6
	PWR_OPERATE_OBSERVER   CEffect = 7
	PWR_OPERATE_TERMINAL   CEffect = 8
	PWR_DISRUPT_SPAWN      CEffect = 9
	PWR_DISRUPT_TOWER      CEffect = 10
	PWR_DISRUPT_SOURCE     CEffect = 11
	PWR_SHIELD             CEffect = 12
	PWR_REGEN_SOURCE       CEffect = 13
	PWR_REGEN_MINERAL      CEffect = 14
	PWR_DISRUPT_TERMINAL   CEffect = 15
	PWR_OPERATE_POWER      CEffect = 16
	PWR_FORTIFY            CEffect = 17
	PWR_OPERATE_CONTROLLER CEffect = 18
	PWR_OPERATE_FACTORY    CEffect = 19
	EFFECT_INVULNERABILITY CEffect = 1001
	EFFECT_COLLAPSE_TIMER  CEffect = 1002
)

const (
	COLOR_RED    CColor = 1
	COLOR_PURPLE CColor = 2
	COLOR_BLUE   CColor = 3
	COLOR_CYAN   CColor = 4
	COLOR_GREEN  CColor = 5
	COLOR_YELLOW CColor = 6
	COLOR_ORANGE CColor = 7
	COLOR_BROWN  CColor = 8
	COLOR_GREY   CColor = 9
	COLOR_WHITE  CColor = 10
)

const (
	FIND_EXIT_TOP                   CFind = 1
	FIND_EXIT_RIGHT                 CFind = 3
	FIND_EXIT_BOTTOM                CFind = 5
	FIND_EXIT_LEFT                  CFind = 7
	FIND_EXIT                       CFind = 10
	FIND_CREEPS                     CFind = 101
	FIND_MY_CREEPS                  CFind = 102
	FIND_HOSTILE_CREEPS             CFind = 103
	FIND_SOURCES_ACTIVE             CFind = 104
	FIND_SOURCES                    CFind = 105
	FIND_DROPPED_RESOURCES          CFind = 106
	FIND_STRUCTURES                 CFind = 107
	FIND_MY_STRUCTURES              CFind = 108
	FIND_HOSTILE_STRUCTURES         CFind = 109
	FIND_FLAGS                      CFind = 110
	FIND_CONSTRUCTION_SITES         CFind = 111
	FIND_MY_SPAWNS                  CFind = 112
	FIND_HOSTILE_SPAWNS             CFind = 113
	FIND_MY_CONSTRUCTION_SITES      CFind = 114
	FIND_HOSTILE_CONSTRUCTION_SITES CFind = 115
	FIND_MINERALS                   CFind = 116
	FIND_NUKES                      CFind = 117
	FIND_TOMBSTONES                 CFind = 118
	FIND_POWER_CREEPS               CFind = 119
	FIND_MY_POWER_CREEPS            CFind = 120
	FIND_HOSTILE_POWER_CREEPS       CFind = 121
	FIND_DEPOSITS                   CFind = 122
	FIND_RUINS                      CFind = 123
)

const (
	ALGORITHM_ASTAR    CAlgorithm = "astar"
	ALGORITHM_DIJKSTRA CAlgorithm = "dijkstra"
)

const (
	TOP          CDirection = 1
	TOP_RIGHT    CDirection = 2
	RIGHT        CDirection = 3
	BOTTOM_RIGHT CDirection = 4
	BOTTOM       CDirection = 5
	BOTTOM_LEFT  CDirection = 6
	LEFT         CDirection = 7
	TOP_LEFT     CDirection = 8
)

const (
	STRUCTURE_SPAWN        CStructure = "spawn"
	STRUCTURE_EXTENSION    CStructure = "extension"
	STRUCTURE_ROAD         CStructure = "road"
	STRUCTURE_WALL         CStructure = "constructedWall"
	STRUCTURE_RAMPART      CStructure = "rampart"
	STRUCTURE_KEEPER_LAIR  CStructure = "keeperLair"
	STRUCTURE_PORTAL       CStructure = "portal"
	STRUCTURE_CONTROLLER   CStructure = "controller"
	STRUCTURE_LINK         CStructure = "link"
	STRUCTURE_STORAGE      CStructure = "storage"
	STRUCTURE_TOWER        CStructure = "tower"
	STRUCTURE_OBSERVER     CStructure = "observer"
	STRUCTURE_POWER_BANK   CStructure = "powerBank"
	STRUCTURE_POWER_SPAWN  CStructure = "powerSpawn"
	STRUCTURE_EXTRACTOR    CStructure = "extractor"
	STRUCTURE_LAB          CStructure = "lab"
	STRUCTURE_TERMINAL     CStructure = "terminal"
	STRUCTURE_CONTAINER    CStructure = "container"
	STRUCTURE_NUKER        CStructure = "nuker"
	STRUCTURE_FACTORY      CStructure = "factory"
	STRUCTURE_INVADER_CORE CStructure = "invaderCore"
)

const (
	LOOK_CREEPS             CLook = "creep"
	LOOK_ENERGY             CLook = "energy"
	LOOK_RESOURCES          CLook = "resource"
	LOOK_SOURCES            CLook = "source"
	LOOK_MINERALS           CLook = "mineral"
	LOOK_DEPOSITS           CLook = "deposit"
	LOOK_STRUCTURES         CLook = "structure"
	LOOK_FLAGS              CLook = "flag"
	LOOK_CONSTRUCTION_SITES CLook = "constructionSite"
	LOOK_NUKES              CLook = "nuke"
	LOOK_TERRAIN            CLook = "terrain"
	LOOK_TOMBSTONES         CLook = "tombstone"
	LOOK_POWER_CREEPS       CLook = "powerCreep"
	LOOK_RUINS              CLook = "ruin"
)

const (
	TERRAIN_MASK_WALL  CTerrain = 1
	TERRAIN_MASK_SWAMP CTerrain = 2
	TERRAIN_MASK_LAVA  CTerrain = 4
)

const (
	RESOURCE_ANY                          CResource = ""
	RESOURCE_ENERGY                       CResource = "energy"
	RESOURCE_POWER                        CResource = "power"
	RESOURCE_HYDROGEN                     CResource = "H"
	RESOURCE_OXYGEN                       CResource = "O"
	RESOURCE_UTRIUM                       CResource = "U"
	RESOURCE_LEMERGIUM                    CResource = "L"
	RESOURCE_KEANIUM                      CResource = "K"
	RESOURCE_ZYNTHIUM                     CResource = "Z"
	RESOURCE_CATALYST                     CResource = "X"
	RESOURCE_GHODIUM                      CResource = "G"
	RESOURCE_SILICON                      CResource = "silicon"
	RESOURCE_METAL                        CResource = "metal"
	RESOURCE_BIOMASS                      CResource = "biomass"
	RESOURCE_MIST                         CResource = "mist"
	RESOURCE_HYDROXIDE                    CResource = "OH"
	RESOURCE_ZYNTHIUM_KEANITE             CResource = "ZK"
	RESOURCE_UTRIUM_LEMERGITE             CResource = "UL"
	RESOURCE_UTRIUM_HYDRIDE               CResource = "UH"
	RESOURCE_UTRIUM_OXIDE                 CResource = "UO"
	RESOURCE_KEANIUM_HYDRIDE              CResource = "KH"
	RESOURCE_KEANIUM_OXIDE                CResource = "KO"
	RESOURCE_LEMERGIUM_HYDRIDE            CResource = "LH"
	RESOURCE_LEMERGIUM_OXIDE              CResource = "LO"
	RESOURCE_ZYNTHIUM_HYDRIDE             CResource = "ZH"
	RESOURCE_ZYNTHIUM_OXIDE               CResource = "ZO"
	RESOURCE_GHODIUM_HYDRIDE              CResource = "GH"
	RESOURCE_GHODIUM_OXIDE                CResource = "GO"
	RESOURCE_UTRIUM_ACID                  CResource = "UH2O"
	RESOURCE_UTRIUM_ALKALIDE              CResource = "UHO2"
	RESOURCE_KEANIUM_ACID                 CResource = "KH2O"
	RESOURCE_KEANIUM_ALKALIDE             CResource = "KHO2"
	RESOURCE_LEMERGIUM_ACID               CResource = "LH2O"
	RESOURCE_LEMERGIUM_ALKALIDE           CResource = "LHO2"
	RESOURCE_ZYNTHIUM_ACID                CResource = "ZH2O"
	RESOURCE_ZYNTHIUM_ALKALIDE            CResource = "ZHO2"
	RESOURCE_GHODIUM_ACID                 CResource = "GH2O"
	RESOURCE_GHODIUM_ALKALIDE             CResource = "GHO2"
	RESOURCE_CATALYZED_UTRIUM_ACID        CResource = "XUH2O"
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE    CResource = "XUHO2"
	RESOURCE_CATALYZED_KEANIUM_ACID       CResource = "XKH2O"
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE   CResource = "XKHO2"
	RESOURCE_CATALYZED_LEMERGIUM_ACID     CResource = "XLH2O"
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE CResource = "XLHO2"
	RESOURCE_CATALYZED_ZYNTHIUM_ACID      CResource = "XZH2O"
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE  CResource = "XZHO2"
	RESOURCE_CATALYZED_GHODIUM_ACID       CResource = "XGH2O"
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE   CResource = "XGHO2"
	RESOURCE_OPS                          CResource = "ops"
	RESOURCE_UTRIUM_BAR                   CResource = "utrium_bar"
	RESOURCE_LEMERGIUM_BAR                CResource = "lemergium_bar"
	RESOURCE_ZYNTHIUM_BAR                 CResource = "zynthium_bar"
	RESOURCE_KEANIUM_BAR                  CResource = "keanium_bar"
	RESOURCE_GHODIUM_MELT                 CResource = "ghodium_melt"
	RESOURCE_OXIDANT                      CResource = "oxidant"
	RESOURCE_REDUCTANT                    CResource = "reductant"
	RESOURCE_PURIFIER                     CResource = "purifier"
	RESOURCE_BATTERY                      CResource = "battery"
	RESOURCE_COMPOSITE                    CResource = "composite"
	RESOURCE_CRYSTAL                      CResource = "crystal"
	RESOURCE_LIQUID                       CResource = "liquid"
	RESOURCE_WIRE                         CResource = "wire"
	RESOURCE_SWITCH                       CResource = "switch"
	RESOURCE_TRANSISTOR                   CResource = "transistor"
	RESOURCE_MICROCHIP                    CResource = "microchip"
	RESOURCE_CIRCUIT                      CResource = "circuit"
	RESOURCE_DEVICE                       CResource = "device"
	RESOURCE_CELL                         CResource = "cell"
	RESOURCE_PHLEGM                       CResource = "phlegm"
	RESOURCE_TISSUE                       CResource = "tissue"
	RESOURCE_MUSCLE                       CResource = "muscle"
	RESOURCE_ORGANOID                     CResource = "organoid"
	RESOURCE_ORGANISM                     CResource = "organism"
	RESOURCE_ALLOY                        CResource = "alloy"
	RESOURCE_TUBE                         CResource = "tube"
	RESOURCE_FIXTURES                     CResource = "fixtures"
	RESOURCE_FRAME                        CResource = "frame"
	RESOURCE_HYDRAULICS                   CResource = "hydraulics"
	RESOURCE_MACHINE                      CResource = "machine"
	RESOURCE_CONDENSATE                   CResource = "condensate"
	RESOURCE_CONCENTRATE                  CResource = "concentrate"
	RESOURCE_EXTRACT                      CResource = "extract"
	RESOURCE_SPIRIT                       CResource = "spirit"
	RESOURCE_EMANATION                    CResource = "emanation"
	RESOURCE_ESSENCE                      CResource = "essence"
)

const (
	MOVE          CBodyPart = "move"
	WORK          CBodyPart = "work"
	CARRY         CBodyPart = "carry"
	ATTACK        CBodyPart = "attack"
	RANGED_ATTACK CBodyPart = "ranged_attack"
	TOUGH         CBodyPart = "tough"
	HEAL          CBodyPart = "heal"
	CLAIM         CBodyPart = "claim"
)

const (
	DENSITY_LOW      CDensity = 1
	DENSITY_MODERATE CDensity = 2
	DENSITY_HIGH     CDensity = 3
	DENSITY_ULTRA    CDensity = 4
)
