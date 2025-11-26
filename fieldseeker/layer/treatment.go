package layer

import (
	"time"

	"github.com/google/uuid"
)

type TreatmentNotInUIWindDirectionType string

const (
	TreatmentNotInUIWindDirectionN  TreatmentNotInUIWindDirectionType = "N"
	TreatmentNotInUIWindDirectionNE TreatmentNotInUIWindDirectionType = "NE"
	TreatmentNotInUIWindDirectionE  TreatmentNotInUIWindDirectionType = "E"
	TreatmentNotInUIWindDirectionSE TreatmentNotInUIWindDirectionType = "SE"
	TreatmentNotInUIWindDirectionS  TreatmentNotInUIWindDirectionType = "S"
	TreatmentNotInUIWindDirectionSW TreatmentNotInUIWindDirectionType = "SW"
	TreatmentNotInUIWindDirectionW  TreatmentNotInUIWindDirectionType = "W"
	TreatmentNotInUIWindDirectionNW TreatmentNotInUIWindDirectionType = "NW"
)

type TreatmentNotInUITFType int16

const (
	TreatmentNotInUITFTrue  TreatmentNotInUITFType = 1
	TreatmentNotInUITFFalse TreatmentNotInUITFType = 0
)

type TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type string

const (
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Orchard                  TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "orchard"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16RowCrops                 TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "row_crops"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16VineCrops                TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "vine_crops"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16AgriculturalGrassesGrain TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "ag_grass_or_grain"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Pasture                  TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "pasture"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16IrrigationStandpipe      TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "irrigation_standpipe"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16DitchCanal               TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "ditch"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Pond                     TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "pond"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Sump                     TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "sump"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Drain                    TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "drain"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16DairyLagoon              TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "dairy_lagoon"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16WastewaterTreatment      TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "wastewater_treatment"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Trough                   TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "trough"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Depression               TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "depression"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16GutterStreet             TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "gutter"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16RainGutterRoof           TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "rain_gutter"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Culvert                  TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "culvert"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Utility                  TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "Utility"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16CatchBasin               TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "catch_basin"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16StreamCreek              TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "stream_or_creek"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Slough                   TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "slough"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16River                    TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "river"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16MarshWetlands            TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "marsh_or_wetlands"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Containers               TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "containers"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16WateringBowl             TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "watering_bowl"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16PlantSaucer              TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "plant_saucer"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16YardDrain                TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "yard_drain"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16PlantAxil                TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "plant_axil"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Treehole                 TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "treehole"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16FountainWaterFeature     TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "foutain_or_water_feature"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16BirdBath                 TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "bird_bath"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16MiscWaterAccumulation    TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "misc_water_accumulation"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16TarpCover                TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "tarp_or_cover"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16SwimmingPool             TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "swimming_pool"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16AbovegroundPool          TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "aboveground_pool"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16KidPool                  TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "kid_pool"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16HotTubSpa                TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "hot_tub"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Applicance               TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "applicance"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16FloodedStructure         TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "flooded_structure"
	TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16LowPoint                 TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type = "low_point"
)

type TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type string

const (
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Dry          TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "Dry"
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Flowing      TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "Flowing"
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Maintained   TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "Maintained"
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Unmaintained TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "Unmaintained"
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169HighOrganic  TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "High Organic"
	TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169FishPresent  TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type = "Fish Present"
)

type TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type string

const (
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Dry          TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Dry"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Flowing      TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Flowing"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Maintained   TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Maintained"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Unmaintained TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Unmaintained"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571HighOrganic  TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "High Organic"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571FishPresent  TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Fish Present"
	TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Stagnant     TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type = "Stagnant"
)

type TreatmentMosquitoActivityType string

const (
	TreatmentMosquitoActivityRoutineinspection TreatmentMosquitoActivityType = "Routine inspection"
	TreatmentMosquitoActivityPretreatment      TreatmentMosquitoActivityType = "Pre-treatment"
	TreatmentMosquitoActivityMaintenance       TreatmentMosquitoActivityType = "Maintenance"
	TreatmentMosquitoActivityULV               TreatmentMosquitoActivityType = "ULV"
	TreatmentMosquitoActivityBARRIER           TreatmentMosquitoActivityType = "BARRIER"
	TreatmentMosquitoActivityLOGIN             TreatmentMosquitoActivityType = "LOGIN"
	TreatmentMosquitoActivityTREATSD           TreatmentMosquitoActivityType = "TREATSD"
	TreatmentMosquitoActivitySD                TreatmentMosquitoActivityType = "SD"
	TreatmentMosquitoActivitySITEVISIT         TreatmentMosquitoActivityType = "SITEVISIT"
	TreatmentMosquitoActivityONLINE            TreatmentMosquitoActivityType = "ONLINE"
	TreatmentMosquitoActivitySYNC              TreatmentMosquitoActivityType = "SYNC"
	TreatmentMosquitoActivityCREATESR          TreatmentMosquitoActivityType = "CREATESR"
	TreatmentMosquitoActivityLC                TreatmentMosquitoActivityType = "LC"
	TreatmentMosquitoActivityACCEPTSR          TreatmentMosquitoActivityType = "ACCEPTSR"
	TreatmentMosquitoActivityPOINT             TreatmentMosquitoActivityType = "POINT"
	TreatmentMosquitoActivityDOWNLOAD          TreatmentMosquitoActivityType = "DOWNLOAD"
	TreatmentMosquitoActivityCOMPLETESR        TreatmentMosquitoActivityType = "COMPLETESR"
	TreatmentMosquitoActivityPOLYGON           TreatmentMosquitoActivityType = "POLYGON"
	TreatmentMosquitoActivityTRAP              TreatmentMosquitoActivityType = "TRAP"
	TreatmentMosquitoActivitySAMPLE            TreatmentMosquitoActivityType = "SAMPLE"
	TreatmentMosquitoActivityQA                TreatmentMosquitoActivityType = "QA"
	TreatmentMosquitoActivityPTA               TreatmentMosquitoActivityType = "PTA"
	TreatmentMosquitoActivityFIELDSCOUTING     TreatmentMosquitoActivityType = "FIELDSCOUTING"
	TreatmentMosquitoActivityOFFLINE           TreatmentMosquitoActivityType = "OFFLINE"
	TreatmentMosquitoActivityLINE              TreatmentMosquitoActivityType = "LINE"
	TreatmentMosquitoActivityTRAPLOCATION      TreatmentMosquitoActivityType = "TRAPLOCATION"
	TreatmentMosquitoActivitySAMPLELOCATION    TreatmentMosquitoActivityType = "SAMPLELOCATION"
	TreatmentMosquitoActivityLCLOCATION        TreatmentMosquitoActivityType = "LCLOCATION"
)

type TreatmentMosquitoProductAreaUnitType string

const (
	TreatmentMosquitoProductAreaUnitAcre TreatmentMosquitoProductAreaUnitType = "acre"
	TreatmentMosquitoProductAreaUnitSqft TreatmentMosquitoProductAreaUnitType = "sq ft"
)

type TreatmentMosquitoProductMeasureUnitType string

const (
	TreatmentMosquitoProductMeasureUnitBriquet TreatmentMosquitoProductMeasureUnitType = "briquet"
	TreatmentMosquitoProductMeasureUnitDryoz   TreatmentMosquitoProductMeasureUnitType = "dry oz"
	TreatmentMosquitoProductMeasureUnitEach    TreatmentMosquitoProductMeasureUnitType = "each"
	TreatmentMosquitoProductMeasureUnitFloz    TreatmentMosquitoProductMeasureUnitType = "fl oz"
	TreatmentMosquitoProductMeasureUnitGal     TreatmentMosquitoProductMeasureUnitType = "gal"
	TreatmentMosquitoProductMeasureUnitLb      TreatmentMosquitoProductMeasureUnitType = "lb"
	TreatmentMosquitoProductMeasureUnitPacket  TreatmentMosquitoProductMeasureUnitType = "packet"
	TreatmentMosquitoProductMeasureUnitPouch   TreatmentMosquitoProductMeasureUnitType = "pouch"
)

type TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type string

const (
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Argo                 TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Argo"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621ATV                  TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "ATV"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Backpack             TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Backpack"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Drone                TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Drone"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Manual               TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Manual"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Truck                TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Truck"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621ULV                  TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "ULV"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621WALS                 TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "WALS"
	TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621AdministrativeAction TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type = "Administrative Action"
)

type TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type string

const (
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0BackpackOne         TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Backpack #1"
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0A1MistSprayerTThree TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "A1  Mist Sprayer (T-3) "
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0SpreaderTwo         TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Spreader #2"
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Guardian73          TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Guardian #73 "
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0ULV74Grizzly        TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "ULV #74 (Grizzly)"
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0ClarkULVSprayer71   TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Clark ULV Sprayer #71"
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0ClarkULVSprayer72   TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Clark ULV Sprayer #72"
	TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Spraybottle         TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type = "Spray bottle"
)

type Treatment struct {
	ObjectID             uint                                                            `field:"OBJECTID"`
	Activity             TreatmentMosquitoActivityType                                   `field:"ACTIVITY"`
	AreaTreated          float64                                                         `field:"TREATAREA"`
	AreaUnit             TreatmentMosquitoProductAreaUnitType                            `field:"AREAUNIT"`
	Product              string                                                          `field:"PRODUCT"`
	Quantity             float64                                                         `field:"QTY"`
	QuantityUnit         TreatmentMosquitoProductMeasureUnitType                         `field:"QTYUNIT"`
	Method               TreatmentTreatmentMETHODd558ca3ccf43440c8160758253967621Type    `field:"METHOD"`
	EquipmentType        TreatmentTreatmentEQUIPTYPE45694d79ff2142ccbe4fa0d1def4fba0Type `field:"EQUIPTYPE"`
	Comments             string                                                          `field:"COMMENTS"`
	AverageTemperature   float64                                                         `field:"AVETEMP"`
	WindSpeed            float64                                                         `field:"WINDSPEED"`
	WindDirection        TreatmentNotInUIWindDirectionType                               `field:"WINDDIR"`
	RainGauge            float64                                                         `field:"RAINGAUGE"`
	Start                time.Time                                                       `field:"STARTDATETIME"`
	Finish               time.Time                                                       `field:"ENDDATETIME"`
	InspID               uuid.UUID                                                       `field:"INSP_ID"`
	Reviewed             TreatmentNotInUITFType                                          `field:"REVIEWED"`
	ReviewedBy           string                                                          `field:"REVIEWEDBY"`
	ReviewedDate         time.Time                                                       `field:"REVIEWEDDATE"`
	LocationName         string                                                          `field:"LOCATIONNAME"`
	Zone                 string                                                          `field:"ZONE"`
	WarningOverride      TreatmentNotInUITFType                                          `field:"WARNINGOVERRIDE"`
	RecordStatus         int16                                                           `field:"RECORDSTATUS"`
	Zone2                string                                                          `field:"ZONE2"`
	TreatedAcres         float64                                                         `field:"TREATACRES"`
	TireCount            int16                                                           `field:"TIRECOUNT"`
	CatchBasinCount      int16                                                           `field:"CBCOUNT"`
	ContainerCount       int16                                                           `field:"CONTAINERCOUNT"`
	GlobalID             uuid.UUID                                                       `field:"GlobalID"`
	TreatmentLength      float64                                                         `field:"TREATMENTLENGTH"`
	TreatmentHours       float64                                                         `field:"TREATMENTHOURS"`
	TreatmentLengthUnits string                                                          `field:"TREATMENTLENGTHUNITS"`
	LinelocID            uuid.UUID                                                       `field:"LINELOCID"`
	PointlocID           uuid.UUID                                                       `field:"POINTLOCID"`
	PolygonlocID         uuid.UUID                                                       `field:"POLYGONLOCID"`
	SrID                 uuid.UUID                                                       `field:"SRID"`
	SdID                 uuid.UUID                                                       `field:"SDID"`
	BarrierrouteID       uuid.UUID                                                       `field:"BARRIERROUTEID"`
	UlvrouteID           uuid.UUID                                                       `field:"ULVROUTEID"`
	FieldTech            string                                                          `field:"FIELDTECH"`
	PtaID                uuid.UUID                                                       `field:"PTAID"`
	Flowrate             float64                                                         `field:"FLOWRATE"`
	Habitat              TreatmentTreatmentHABITAT0afee7ebf9ea47078483cccfe60f0d16Type   `field:"HABITAT"`
	TreatHectares        float64                                                         `field:"TREATHECTARES"`
	InventoryLocation    string                                                          `field:"INVLOC"`
	TempConditions       TreatmentTreatmentSITECONDf812e1f64dcb4dc9a75da9d00abe6169Type  `field:"temp_SITECOND"`
	Conditions           TreatmentTreatmentSITECOND5a15bf36fa124280b961f31cd1a9b571Type  `field:"SITECOND"`
	TotalCostProduct     float64                                                         `field:"TotalCostProdcut"`
	CreationDate         time.Time                                                       `field:"CreationDate"`
	Creator              string                                                          `field:"Creator"`
	EditDate             time.Time                                                       `field:"EditDate"`
	Editor               string                                                          `field:"Editor"`
	TargetSpecies        string                                                          `field:"TARGETSPECIES"`
}
