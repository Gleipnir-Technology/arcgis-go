package layer

import (
	"time"

	"github.com/google/uuid"
)

type LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type string

const (
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Orchard                  LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "orchard"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3RowCrops                 LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "row_crops"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3VineCrops                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "vine_crops"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3AgriculturalGrassesGrain LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "ag_grasses_or_grain"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Pasture                  LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "pasture"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3IrrigationStandpipe      LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "irrigation_standpipe"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3DitchCanal               LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "ditch"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Pond                     LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "pond"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Sump                     LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "sump"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Drain                    LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "drain"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3DairyLagoon              LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "dairy_lagoon"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3WastewaterTreatment      LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "wastewater_treatment"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Trough                   LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "trough"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Depression               LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "depression"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Gutterstreet             LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "gutter"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3RainGutterroof           LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "rain_gutter"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Culvert                  LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "culvert"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Utility                  LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "utility"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3CatchBasin               LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "catch_basin"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3StreamCreek              LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "stream_or_creek"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Slough                   LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "slough"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3River                    LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "river"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3MarshWetland             LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "marsh_or_wetland"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Containers               LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "containers"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3WateringBowl             LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "watering_bowl"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3PlantSaucer              LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "plant_saucer"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3YardDrain                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "yard_drain"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3PlantAxil                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "plant_axil"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Treehole                 LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "treehole"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3FountainWaterFeature     LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "fountain_or_water_feature"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3BirdBath                 LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "bird_bath"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3MiscWaterAccumulation    LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "misc_water_accumulation"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3TarpCover                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "tarp_or_cover"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3SwimmingPool             LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "swimming_pool"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3AbovegroundPool          LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "aboveground_pool"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3KidPool                  LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "kid_pool"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3HotTubSpa                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "hot_tub"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Appliance                LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "appliance"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Tires                    LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "tires"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3FloodedStructure         LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "flooded_structure"
	LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3LowPoint                 LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type = "low_point"
)

type LineLocationLocationPriorityType string

const (
	LineLocationLocationPriorityLow    LineLocationLocationPriorityType = "Low"
	LineLocationLocationPriorityMedium LineLocationLocationPriorityType = "Medium"
	LineLocationLocationPriorityHigh   LineLocationLocationPriorityType = "High"
	LineLocationLocationPriorityNone   LineLocationLocationPriorityType = "None"
)

type LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type string

const (
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Residential  LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "residential"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Commercial   LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "commercial"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Industrial   LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "industrial"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Agricultural LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "agricultural"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Mixeduse     LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "mixed_use"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02PublicDomain LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "public_domain"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Natural      LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "natural"
	LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Municipal    LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type = "municipal"
)

type LineLocationNotInUITFType int16

const (
	LineLocationNotInUITFTrue  LineLocationNotInUITFType = 1
	LineLocationNotInUITFFalse LineLocationNotInUITFType = 0
)

type LineLocationLocationSymbologyType string

const (
	LineLocationLocationSymbologyActionrequired   LineLocationLocationSymbologyType = "ACTION"
	LineLocationLocationSymbologyInactive         LineLocationLocationSymbologyType = "INACTIVE"
	LineLocationLocationSymbologyNoactionrequired LineLocationLocationSymbologyType = "NONE"
)

type LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type string

const (
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008FloodIrrigation             LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "flood_irrigation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008FurrowIrrigation            LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "furrow_irrigation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008DripIrrigation              LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "drip_irrigation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008SprinklerIrrigation         LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "sprinkler_irrigation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008WastewaterIrrigation        LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "wastewater_irrigation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008IrrigationRunoff            LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "irrigation_runoff"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008RainwaterAccumulation       LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "rainwater_accumulation"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Leak                        LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "leak"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Seepage                     LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "seepage"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008StoredWater                 LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "stored_water"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008WastwaterSystem             LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "wastewater_system"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008PermanentNaturalWater       LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "permanent_natural_water"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008TemporaryNaturalWater       LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "temporary_natural_water"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008RecreationalOrnamentalWater LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "recreational_or_ornamental_water"
	LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008WaterConveyance             LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type = "water_conveyance"
)

type LineLocation struct {
	Objectid                    uint                                                                    `field:"OBJECTID"`
	Name                        string                                                                  `field:"NAME"`
	Zone                        string                                                                  `field:"ZONE"`
	Habitat                     LineLocationLineLocationHABITATfc51bdc4f1954df58206d69ce14182f3Type     `field:"HABITAT"`
	Priority                    LineLocationLocationPriorityType                                        `field:"PRIORITY"`
	UseType                     LineLocationLineLocationUSETYPE2aeca2e60d2f455c86fc34895dc80a02Type     `field:"USETYPE"`
	Active                      LineLocationNotInUITFType                                               `field:"ACTIVE"`
	Description                 string                                                                  `field:"DESCRIPTION"`
	AccessDescription           string                                                                  `field:"ACCESSDESC"`
	Comments                    string                                                                  `field:"COMMENTS"`
	Symbology                   LineLocationLocationSymbologyType                                       `field:"SYMBOLOGY"`
	ExternalId                  string                                                                  `field:"EXTERNALID"`
	Acres                       float64                                                                 `field:"ACRES"`
	NextScheduledAction         time.Time                                                               `field:"NEXTACTIONDATESCHEDULED"`
	LarvalInspectionInterval    int16                                                                   `field:"LARVINSPECTINTERVAL"`
	Length                      float64                                                                 `field:"LENGTH_FT"`
	Width                       float64                                                                 `field:"WIDTH_FT"`
	Zone2                       string                                                                  `field:"ZONE2"`
	Locationnumber              int32                                                                   `field:"LOCATIONNUMBER"`
	GlobalID                    uuid.UUID                                                               `field:"GlobalID"`
	CreatedUser                 string                                                                  `field:"created_user"`
	CreatedDate                 time.Time                                                               `field:"created_date"`
	LastEditedUser              string                                                                  `field:"last_edited_user"`
	LastEditedDate              time.Time                                                               `field:"last_edited_date"`
	LastInspectionDate          time.Time                                                               `field:"LASTINSPECTDATE"`
	LastInspectionBreeding      string                                                                  `field:"LASTINSPECTBREEDING"`
	LastInspectionAverageLarvae float64                                                                 `field:"LASTINSPECTAVGLARVAE"`
	LastInspectionAveragePupae  float64                                                                 `field:"LASTINSPECTAVGPUPAE"`
	LastInspectionLarvalStages  string                                                                  `field:"LASTINSPECTLSTAGES"`
	LastInspectionAction        string                                                                  `field:"LASTINSPECTACTIONTAKEN"`
	LastInspectionFieldSpecies  string                                                                  `field:"LASTINSPECTFIELDSPECIES"`
	LastTreatmentDate           time.Time                                                               `field:"LASTTREATDATE"`
	LastTreatmentProduct        string                                                                  `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity       float64                                                                 `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit   string                                                                  `field:"LASTTREATQTYUNIT"`
	Hectares                    float64                                                                 `field:"HECTARES"`
	LastInspectionActivity      string                                                                  `field:"LASTINSPECTACTIVITY"`
	LastTreatmentActivity       string                                                                  `field:"LASTTREATACTIVITY"`
	LengthMeters                float64                                                                 `field:"LENGTH_METERS"`
	WidthMeters                 float64                                                                 `field:"WIDTH_METERS"`
	LastInspectionConditions    string                                                                  `field:"LASTINSPECTCONDITIONS"`
	WaterOrigin                 LineLocationLineLocationWATERORIGIN84723d92306a46f48ef169b55a916008Type `field:"WATERORIGIN"`
	CreationDate                time.Time                                                               `field:"CreationDate"`
	Creator                     string                                                                  `field:"Creator"`
	EditDate                    time.Time                                                               `field:"EditDate"`
	Editor                      string                                                                  `field:"Editor"`
	Jurisdiction                string                                                                  `field:"JURISDICTION"`
	ShapeLength                 float64                                                                 `field:"Shape__Length"`
}
