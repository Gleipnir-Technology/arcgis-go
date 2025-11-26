package layer

import (
	"time"

	"github.com/google/uuid"
)

type PolygonLocationNotInUITFType int16

const (
	PolygonLocationNotInUITFTrue  PolygonLocationNotInUITFType = 1
	PolygonLocationNotInUITFFalse PolygonLocationNotInUITFType = 0
)

type PolygonLocationLocationSymbologyType string

const (
	PolygonLocationLocationSymbologyActionrequired   PolygonLocationLocationSymbologyType = "ACTION"
	PolygonLocationLocationSymbologyInactive         PolygonLocationLocationSymbologyType = "INACTIVE"
	PolygonLocationLocationSymbologyNoactionrequired PolygonLocationLocationSymbologyType = "NONE"
)

type PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType string

const (
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aFloodIrrigation             PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "flood_irrigation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aFurrowIrrigation            PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "furrow_irrigation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aDripIrrigation              PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "drip_irritation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aSprinklerIrrigation         PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "sprinkler_irrigation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aWastewaterIrrigation        PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "wastewater_irrigation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aIrrigationRunoff            PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "irrigation_runoff"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aRainwaterAccumulation       PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "rainwater_accumulation"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aLeak                        PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "leak"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aSeepage                     PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "seepage"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aStoredWater                 PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "stored_water"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aWastwaterSystem             PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "wastewater_system"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aPermanentNaturalWater       PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "permanent_natural_water"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aTemporaryNaturalWater       PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "temporary_natural_water"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aRecreationalOrnamentalWater PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "recreational_or_ornamental_water"
	PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aWaterConveyance             PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType = "water_conveyance"
)

type PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType string

const (
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdOrchard                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "orchard"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdRowCrops                 PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "row_crops"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdVineCrops                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "vine_crops"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdAgriculturalGrassesGrain PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "ag_grasses_or_grain"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdPasture                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "pasture"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdIrrigationStandpipe      PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "irrigation_standpipe"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdDitchCanal               PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "ditch"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdDairyLagoon              PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "dairy_lagoon"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdWastewaterTreatment      PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "wastewater_treatment"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdTrough                   PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "trough"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdDepression               PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "depression"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdGutterstreet             PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "gutter"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdRainGutterroof           PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "rain_gutter"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdCulvert                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "culvert"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdUtility                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "utility"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdCatchBasin               PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "catch_basin"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdStreamCreek              PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "stream_or_creek"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdSlough                   PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "slough"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdRiver                    PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "river"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdMarshWetland             PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "marsh_or_wetland"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdContainers               PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "containers"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdWateringBowl             PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "watering_bowl"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdPlantSaucer              PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "plant_saucer"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdYardDrain                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "yard_drain"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdPlantAxil                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "plant_axil"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdTreehole                 PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "treehole"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdFountainWaterFeature     PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "fountain_or_water_feature"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdBirdBath                 PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "bird_bath"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdMiscWaterAccumulation    PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "misc_water_accumulation"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdTarpCover                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "tarp_or_cover"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdSwimmingPool             PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "swimming_pool"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdAbovegroundPool          PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "aboveground_pool"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdKidPool                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "kid_pool"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdHotTubSpa                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "hot_tub"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdAppliance                PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "appliance"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdTires                    PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "tires"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdFloodedStructure         PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "flooded_structure"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdLowPoint                 PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "low_point"
	PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdUnknown                  PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType = "unknown"
)

type PolygonLocationLocationPriorityType string

const (
	PolygonLocationLocationPriorityLow    PolygonLocationLocationPriorityType = "Low"
	PolygonLocationLocationPriorityMedium PolygonLocationLocationPriorityType = "Medium"
	PolygonLocationLocationPriorityHigh   PolygonLocationLocationPriorityType = "High"
	PolygonLocationLocationPriorityNone   PolygonLocationLocationPriorityType = "None"
)

type PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type string

const (
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Residential  PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "residential"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Commercial   PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "commercial"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Industrial   PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "industrial"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Agricultural PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "agricultural"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Mixeduse     PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "mixed_use"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27PublicDomain PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "public_domain"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Natural      PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "natural"
	PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Municipal    PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type = "municipal"
)

type PolygonLocation struct {
	ObjectID                    uint                                                                          `field:"OBJECTID"`
	Name                        string                                                                        `field:"NAME"`
	Zone                        string                                                                        `field:"ZONE"`
	Habitat                     PolygonLocationPolygonLocationHABITAT45e9dde79ac84d959df8b65ba7d5dafdType     `field:"HABITAT"`
	Priority                    PolygonLocationLocationPriorityType                                           `field:"PRIORITY"`
	UseType                     PolygonLocationPolygonLocationUSETYPEe546154cb9544b9aa8e7b13e8e258b27Type     `field:"USETYPE"`
	Active                      PolygonLocationNotInUITFType                                                  `field:"ACTIVE"`
	Description                 string                                                                        `field:"DESCRIPTION"`
	AccessDescription           string                                                                        `field:"ACCESSDESC"`
	Comments                    string                                                                        `field:"COMMENTS"`
	Symbology                   PolygonLocationLocationSymbologyType                                          `field:"SYMBOLOGY"`
	ExternalID                  string                                                                        `field:"EXTERNALID"`
	Acres                       float64                                                                       `field:"ACRES"`
	NextScheduledAction         time.Time                                                                     `field:"NEXTACTIONDATESCHEDULED"`
	LarvalInspectionInterval    int16                                                                         `field:"LARVINSPECTINTERVAL"`
	Zone2                       string                                                                        `field:"ZONE2"`
	Locationnumber              int32                                                                         `field:"LOCATIONNUMBER"`
	GlobalID                    uuid.UUID                                                                     `field:"GlobalID"`
	LastInspectionDate          time.Time                                                                     `field:"LASTINSPECTDATE"`
	LastInspectionBreeding      string                                                                        `field:"LASTINSPECTBREEDING"`
	LastInspectionAverageLarvae float64                                                                       `field:"LASTINSPECTAVGLARVAE"`
	LastInspectionAveragePupae  float64                                                                       `field:"LASTINSPECTAVGPUPAE"`
	LastInspectionLarvalStages  string                                                                        `field:"LASTINSPECTLSTAGES"`
	LastInspectionAction        string                                                                        `field:"LASTINSPECTACTIONTAKEN"`
	LastInspectionFieldSpecies  string                                                                        `field:"LASTINSPECTFIELDSPECIES"`
	LastTreatmentDate           time.Time                                                                     `field:"LASTTREATDATE"`
	LastTreatmentProduct        string                                                                        `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity       float64                                                                       `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit   string                                                                        `field:"LASTTREATQTYUNIT"`
	Hectares                    float64                                                                       `field:"HECTARES"`
	LastInspectionActivity      string                                                                        `field:"LASTINSPECTACTIVITY"`
	LastTreatmentActivity       string                                                                        `field:"LASTTREATACTIVITY"`
	LastInspectionConditions    string                                                                        `field:"LASTINSPECTCONDITIONS"`
	WaterOrigin                 PolygonLocationPolygonLocationWATERORIGINe9018e925f474ff98a7cb818d848dc7aType `field:"WATERORIGIN"`
	Filter                      string                                                                        `field:"Filter"`
	CreationDate                time.Time                                                                     `field:"CreationDate"`
	Creator                     string                                                                        `field:"Creator"`
	EditDate                    time.Time                                                                     `field:"EditDate"`
	Editor                      string                                                                        `field:"Editor"`
	Jurisdiction                string                                                                        `field:"JURISDICTION"`
	ShapeArea                   float64                                                                       `field:"Shape__Area"`
	ShapeLength                 float64                                                                       `field:"Shape__Length"`
}
