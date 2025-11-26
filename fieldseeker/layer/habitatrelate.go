package layer

import (
	"time"

	"github.com/google/uuid"
)

type HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type string

const (
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Orchard                  HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "orchard"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953RowCrops                 HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "row_crops"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953VineCrops                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "vine_crops"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953AgriculturalGrassesGrain HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "ag_grasses_or_grain"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Pasture                  HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "pasture"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953IrrigationStandpipe      HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "irrigation_standpipe"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953DitchCanal               HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "ditch"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Pond                     HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "pond"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Sump                     HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "sump"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Drain                    HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "drain"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953DairyLagoon              HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "dairy_lagoon"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953WastewaterTreatment      HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "wastewater_treatment"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Trough                   HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "trough"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Depression               HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "depression"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Gutterstreet             HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "gutter"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953RainGutterroof           HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "rain_gutter"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Culvert                  HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "culvert"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Utility                  HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "utility"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953CatchBasin               HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "catch_basin"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953StreamCreek              HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "stream_or_creek"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Slough                   HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "slough"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953River                    HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "river"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953MarshWetland             HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "marsh_or_wetlands"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Containers               HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "containers"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953WateringBowl             HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "watering_bowl"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953PlantSaucer              HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "plant_saucer"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953YardDrain                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "yard_drain"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953PlantAxil                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "plant_axil"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Treehole                 HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "treehole"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953FountainWaterFeature     HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "fountain_or_water_feature"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953BirdBath                 HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "bird_bath"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953MiscWaterAccumulation    HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "misc_water_accumulation"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953TarpCover                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "tarp_or_cover"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953SwimmingPool             HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "swimming_pool"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953AbovegroundPool          HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "aboveground_pool"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953KidPool                  HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "kid_pool"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953HotTubSpa                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "hot_tub"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Appliance                HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "appliance"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Tires                    HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "tires"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953FloodedStructure         HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "flooded_structure"
	HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953LowPoint                 HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type = "low_point"
)

type HabitatRelate struct {
	ObjectID       uint                                                                      `field:"OBJECTID"`
	ForeignID      uuid.UUID                                                                 `field:"FOREIGN_ID"`
	GlobalID       uuid.UUID                                                                 `field:"GlobalID"`
	CreatedUser    string                                                                    `field:"created_user"`
	CreatedDate    time.Time                                                                 `field:"created_date"`
	LastEditedUser string                                                                    `field:"last_edited_user"`
	LastEditedDate time.Time                                                                 `field:"last_edited_date"`
	HabitatType    HabitatRelateHabitatRelateHABITATTYPE2e81cf2f550e400783cf284f3cec3953Type `field:"HABITATTYPE"`
	CreationDate   time.Time                                                                 `field:"CreationDate"`
	Creator        string                                                                    `field:"Creator"`
	EditDate       time.Time                                                                 `field:"EditDate"`
	Editor         string                                                                    `field:"Editor"`
}
