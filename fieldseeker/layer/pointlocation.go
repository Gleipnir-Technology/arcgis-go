package layer

import (
	"time"

	"github.com/google/uuid"
)

type PointLocationLocationSymbologyType string

const (
	PointLocationLocationSymbologyActionrequired   PointLocationLocationSymbologyType = "ACTION"
	PointLocationLocationSymbologyInactive         PointLocationLocationSymbologyType = "INACTIVE"
	PointLocationLocationSymbologyNoactionrequired PointLocationLocationSymbologyType = "NONE"
)

type PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type string

const (
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97FloodIrrigation             PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "flood_irrigation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97FurrowIrrigation            PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "furrow_irrigation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97DripIrrigation              PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "drip_irrigation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97SprinklerIrrigation         PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "sprinkler_irrigation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97WastewaterIrrigation        PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "wastewater_irrigation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97IrrigationRunoff            PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "irrigation_runoff"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97StormwaterMunicipalRunoff   PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "stormwater_or_municipal_runoff"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97IndustrialRunoff            PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "industrial_runoff"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97RainwaterAccumulation       PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "rainwater_accumulation"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Leak                        PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "leak"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Seepage                     PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "seepage"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97StoredWater                 PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "stored_water"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97WastewaterSystem            PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "wastewater_system"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97PermanentNaturalWater       PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "permanent_natural_water"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97TemporaryNaturalWater       PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "temporary_natural_water"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97RecreationalOrnamentalWater PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "recreational_or_orenamental_water"
	PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97WaterConveyance             PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type = "water_conveyance"
)

type PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType string

const (
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fBryanFeguson      PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Bryan Feguson"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fRickAlverez       PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Rick Alverez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fAlysiaDavis       PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Alysia Davis"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fBryanRuiz         PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Bryan Ruiz"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fKoryWilson        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Kory Wilson"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fAdrianSifuentes   PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Adrian Sifuentes"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fMarcoMartinez     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Marco Martinez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fCarlosRodriguez   PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Carlos Rodriguez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fLandonMcGill      PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Landon McGill"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fTedMcGill         PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Ted McGill"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fMarioSanchez      PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Mario Sanchez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fJorgePerez        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Jorge Perez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fArturoGarciaTrejo PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Arturo Garcia-Trejo"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fLisaSalgado       PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Lisa Salgado"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fLawrenceGuzman    PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Lawrence Guzman"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fTriciaSnowden     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Tricia Snowden"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fRyanSpratt        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Ryan Spratt"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fAndreaTroupin     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Andrea Troupin"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fMarkNakata        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Mark Nakata"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fPabloOrtega       PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Pablo Ortega"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fBenjaminSperry    PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Benjamin Sperry"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fFatimaHidalgo     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Fatima Hidalgo"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fZackeryBarragan   PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Zackery Barragan"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fYajairaGodinez    PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Yajaira Godinez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fJakeMaldonado     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Jake Maldonado"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fRafaelRamirez     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Rafael Ramirez"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fCarlosPalacios    PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Carlos Palacios"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fAaronFredrick     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Aaron Fredrick"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fJoshMalone        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Josh Malone"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fAlecCaposella     PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Alec Caposella"
	PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fLauraRamos        PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType = "Laura Ramos"
)

type PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type string

const (
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5SourceRemoved                       PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Source Removed"
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5PoolMaintained                      PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Pool Maintained"
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5SourceScreenedegyarddrainscreens    PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Source Screened"
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5CropChange                          PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Crop Change"
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5ConsistentlyLoworNoMosquitoActivity PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Low or No Mosquito Activity"
	PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5ConsitentFishPresence               PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type = "Consistent Fish Presence"
)

type PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type string

const (
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Orchard                  PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "orchard"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661RowCrops                 PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "row_crops"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661VineCrops                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "vine_crops"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661AgriculturalGrassesGrain PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "ag_grasses_or_grain"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Pasture                  PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "pasture"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661IrrigationStandpipe      PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "irrigation_standpipe"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661DitchCanal               PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "ditch"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Pond                     PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "pond"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Sump                     PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "sump"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Drain                    PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "drain"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661DairyLagoon              PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "dairy_lagoon"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661WastewaterTreatment      PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "wastewater_treatment"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Trough                   PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "trough"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Depression               PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "depression"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661GutterStreet             PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "gutter"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661RainGutterRoof           PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "rain_gutter"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Culvert                  PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "culvert"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Utility                  PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "utility"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661CatchBasin               PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "catch_basin"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661StreamCreek              PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "stream_or_creek"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Slough                   PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "slough"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661River                    PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "river"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661MarshWetland             PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "marsh_or_wetlands"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Containers               PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "containers"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661WateringBowl             PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "watering_bowl"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661PlantSaucer              PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "plant_saucer"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661YardDrain                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "yard_drain"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661PlantAxil                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "plant_axil"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Treehole                 PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "treehole"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661FountainWaterFeature     PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "fountain_or_water_feature"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661BirdBath                 PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "bird_bath"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661MiscWaterAccumulation    PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "misc_water_accumulation"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661TarpCover                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "tarp_or_cover"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661SwimmingPool             PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "swimming_pool"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661AbovegroundPool          PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "aboveground_pool"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661KidPool                  PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "kid_pool"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661HotTubSpa                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "hot_tub"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Appliance                PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "appliance"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Tires                    PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "tires"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661FloodedStructure         PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "flooded_structure"
	PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661LowPoint                 PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type = "low_point"
)

type PointLocationLocationPriorityType string

const (
	PointLocationLocationPriorityLow    PointLocationLocationPriorityType = "Low"
	PointLocationLocationPriorityMedium PointLocationLocationPriorityType = "Medium"
	PointLocationLocationPriorityHigh   PointLocationLocationPriorityType = "High"
	PointLocationLocationPriorityNone   PointLocationLocationPriorityType = "None"
)

type PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType string

const (
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eResidential  PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "residential"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eCommercial   PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "commercial"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eAgricultural PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "agricultural"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eIndustrial   PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "industrial"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eMixedUse     PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "mixed_use"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89ePublicDomain PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "public_domain"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eNatural      PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "natural"
	PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eMunicipal    PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType = "municipal"
)

type PointLocationNotInUITFType int16

const (
	PointLocationNotInUITFTrue  PointLocationNotInUITFType = 1
	PointLocationNotInUITFFalse PointLocationNotInUITFType = 0
)

type PointLocation struct {
	Objectid                    uint                                                                           `field:"OBJECTID"`
	Name                        string                                                                         `field:"NAME"`
	Zone                        string                                                                         `field:"ZONE"`
	Habitat                     PointLocationPointLocationHABITATb4d8135a497949c88bb367ec7230e661Type          `field:"HABITAT"`
	Priority                    PointLocationLocationPriorityType                                              `field:"PRIORITY"`
	UseType                     PointLocationPointLocationUSETYPE58d62d18ef4f47fc8cb9874df867f89eType          `field:"USETYPE"`
	Active                      PointLocationNotInUITFType                                                     `field:"ACTIVE"`
	Description                 string                                                                         `field:"DESCRIPTION"`
	AccessDescription           string                                                                         `field:"ACCESSDESC"`
	Comments                    string                                                                         `field:"COMMENTS"`
	Symbology                   PointLocationLocationSymbologyType                                             `field:"SYMBOLOGY"`
	ExternalId                  string                                                                         `field:"EXTERNALID"`
	NextScheduledAction         time.Time                                                                      `field:"NEXTACTIONDATESCHEDULED"`
	LarvalInspectionInterval    int16                                                                          `field:"LARVINSPECTINTERVAL"`
	Zone2                       string                                                                         `field:"ZONE2"`
	Locationnumber              int32                                                                          `field:"LOCATIONNUMBER"`
	GlobalID                    uuid.UUID                                                                      `field:"GlobalID"`
	SourceType                  string                                                                         `field:"STYPE"`
	LastInspectionDate          time.Time                                                                      `field:"LASTINSPECTDATE"`
	LastInspectionBreeding      string                                                                         `field:"LASTINSPECTBREEDING"`
	LastInspectionAverageLarvae float64                                                                        `field:"LASTINSPECTAVGLARVAE"`
	LastInspectionAveragePupae  float64                                                                        `field:"LASTINSPECTAVGPUPAE"`
	LastInspectionLarvalStages  string                                                                         `field:"LASTINSPECTLSTAGES"`
	LastInspectionAction        string                                                                         `field:"LASTINSPECTACTIONTAKEN"`
	LastInspectionFieldSpecies  string                                                                         `field:"LASTINSPECTFIELDSPECIES"`
	LastTreatmentDate           time.Time                                                                      `field:"LASTTREATDATE"`
	LastTreatmentProduct        string                                                                         `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity       float64                                                                        `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit   string                                                                         `field:"LASTTREATQTYUNIT"`
	LastInspectionActivity      string                                                                         `field:"LASTINSPECTACTIVITY"`
	LastTreatmentActivity       string                                                                         `field:"LASTTREATACTIVITY"`
	LastInspectionConditions    string                                                                         `field:"LASTINSPECTCONDITIONS"`
	WaterOrigin                 PointLocationPointLocationWATERORIGIN197b22bff3eb4dad8899986460f6ea97Type      `field:"WATERORIGIN"`
	X                           float64                                                                        `field:"X"`
	Y                           float64                                                                        `field:"Y"`
	AssignedTech                PointLocationPointLocationASSIGNEDTECH9393a1622474429d85bedaa44e4c091fType     `field:"ASSIGNEDTECH"`
	CreationDate                time.Time                                                                      `field:"CreationDate"`
	Creator                     string                                                                         `field:"Creator"`
	EditDate                    time.Time                                                                      `field:"EditDate"`
	Editor                      string                                                                         `field:"Editor"`
	Jurisdiction                string                                                                         `field:"JURISDICTION"`
	ReasonForDeactivation       PointLocationPointLocationdeactivatereasondd303085b33c48948c47fa847dd9d7c5Type `field:"deactivate_reason"`
	ScalarPriority              int32                                                                          `field:"scalarPriority"`
	SourceStatus                string                                                                         `field:"sourceStatus"`
}
