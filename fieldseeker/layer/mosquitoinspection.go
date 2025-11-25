package layer

import (
	"time"

	"github.com/google/uuid"
)

type MosquitoInspectionMosquitoAdultActivityType string

const (
	MosquitoInspectionMosquitoAdultActivityNone     MosquitoInspectionMosquitoAdultActivityType = "None"
	MosquitoInspectionMosquitoAdultActivityLight    MosquitoInspectionMosquitoAdultActivityType = "Light"
	MosquitoInspectionMosquitoAdultActivityModerate MosquitoInspectionMosquitoAdultActivityType = "Moderate"
	MosquitoInspectionMosquitoAdultActivityIntense  MosquitoInspectionMosquitoAdultActivityType = "Intense"
)

type MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType string

const (
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeOne       MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "1"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeTwo       MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "2"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeThree     MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "3"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeFour      MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "4"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeFive      MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "5"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeOneTwo    MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "1-2"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeTwoThree  MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "2-3"
	MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeThreeFour MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType = "3-4"
)

type MosquitoInspectionMosquitoFieldSpeciesType string

const (
	MosquitoInspectionMosquitoFieldSpeciesAedes MosquitoInspectionMosquitoFieldSpeciesType = "Aedes"
	MosquitoInspectionMosquitoFieldSpeciesCulex MosquitoInspectionMosquitoFieldSpeciesType = "Culex"
)

type MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270Type string

const (
	MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270Yes MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270Type = "yes"
	MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270No  MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270Type = "no"
)

type MosquitoInspectionMosquitoActivityType string

const (
	MosquitoInspectionMosquitoActivityRoutineinspection MosquitoInspectionMosquitoActivityType = "Routine inspection"
	MosquitoInspectionMosquitoActivityPretreatment      MosquitoInspectionMosquitoActivityType = "Pre-treatment"
	MosquitoInspectionMosquitoActivityMaintenance       MosquitoInspectionMosquitoActivityType = "Maintenance"
	MosquitoInspectionMosquitoActivityULV               MosquitoInspectionMosquitoActivityType = "ULV"
	MosquitoInspectionMosquitoActivityBARRIER           MosquitoInspectionMosquitoActivityType = "BARRIER"
	MosquitoInspectionMosquitoActivityLOGIN             MosquitoInspectionMosquitoActivityType = "LOGIN"
	MosquitoInspectionMosquitoActivityTREATSD           MosquitoInspectionMosquitoActivityType = "TREATSD"
	MosquitoInspectionMosquitoActivitySD                MosquitoInspectionMosquitoActivityType = "SD"
	MosquitoInspectionMosquitoActivitySITEVISIT         MosquitoInspectionMosquitoActivityType = "SITEVISIT"
	MosquitoInspectionMosquitoActivityONLINE            MosquitoInspectionMosquitoActivityType = "ONLINE"
	MosquitoInspectionMosquitoActivitySYNC              MosquitoInspectionMosquitoActivityType = "SYNC"
	MosquitoInspectionMosquitoActivityCREATESR          MosquitoInspectionMosquitoActivityType = "CREATESR"
	MosquitoInspectionMosquitoActivityLC                MosquitoInspectionMosquitoActivityType = "LC"
	MosquitoInspectionMosquitoActivityACCEPTSR          MosquitoInspectionMosquitoActivityType = "ACCEPTSR"
	MosquitoInspectionMosquitoActivityPOINT             MosquitoInspectionMosquitoActivityType = "POINT"
	MosquitoInspectionMosquitoActivityDOWNLOAD          MosquitoInspectionMosquitoActivityType = "DOWNLOAD"
	MosquitoInspectionMosquitoActivityCOMPLETESR        MosquitoInspectionMosquitoActivityType = "COMPLETESR"
	MosquitoInspectionMosquitoActivityPOLYGON           MosquitoInspectionMosquitoActivityType = "POLYGON"
	MosquitoInspectionMosquitoActivityTRAP              MosquitoInspectionMosquitoActivityType = "TRAP"
	MosquitoInspectionMosquitoActivitySAMPLE            MosquitoInspectionMosquitoActivityType = "SAMPLE"
	MosquitoInspectionMosquitoActivityQA                MosquitoInspectionMosquitoActivityType = "QA"
	MosquitoInspectionMosquitoActivityPTA               MosquitoInspectionMosquitoActivityType = "PTA"
	MosquitoInspectionMosquitoActivityFIELDSCOUTING     MosquitoInspectionMosquitoActivityType = "FIELDSCOUTING"
	MosquitoInspectionMosquitoActivityOFFLINE           MosquitoInspectionMosquitoActivityType = "OFFLINE"
	MosquitoInspectionMosquitoActivityLINE              MosquitoInspectionMosquitoActivityType = "LINE"
	MosquitoInspectionMosquitoActivityTRAPLOCATION      MosquitoInspectionMosquitoActivityType = "TRAPLOCATION"
	MosquitoInspectionMosquitoActivitySAMPLELOCATION    MosquitoInspectionMosquitoActivityType = "SAMPLELOCATION"
	MosquitoInspectionMosquitoActivityLCLOCATION        MosquitoInspectionMosquitoActivityType = "LCLOCATION"
)

type MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type string

const (
	MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5TreatmentPesticide                 MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type = "Treatment"
	MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5TreatmentMechanicalorBiological    MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type = "Mechanical or Biological Treatment"
	MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5ResidentScheduleRequest            MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type = "Resident Schedule Request"
	MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5AdministrativeFlyerNoticeorWarrent MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type = "Administrative"
)

type MosquitoInspectionNotInUIWindDirectionType string

const (
	MosquitoInspectionNotInUIWindDirectionN  MosquitoInspectionNotInUIWindDirectionType = "N"
	MosquitoInspectionNotInUIWindDirectionNE MosquitoInspectionNotInUIWindDirectionType = "NE"
	MosquitoInspectionNotInUIWindDirectionE  MosquitoInspectionNotInUIWindDirectionType = "E"
	MosquitoInspectionNotInUIWindDirectionSE MosquitoInspectionNotInUIWindDirectionType = "SE"
	MosquitoInspectionNotInUIWindDirectionS  MosquitoInspectionNotInUIWindDirectionType = "S"
	MosquitoInspectionNotInUIWindDirectionSW MosquitoInspectionNotInUIWindDirectionType = "SW"
	MosquitoInspectionNotInUIWindDirectionW  MosquitoInspectionNotInUIWindDirectionType = "W"
	MosquitoInspectionNotInUIWindDirectionNW MosquitoInspectionNotInUIWindDirectionType = "NW"
)

type MosquitoInspectionNotInUITFType int16

const (
	MosquitoInspectionNotInUITFTrue  MosquitoInspectionNotInUITFType = 1
	MosquitoInspectionNotInUITFFalse MosquitoInspectionNotInUITFType = 0
)

type MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type string

const (
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Dry                   MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Dry"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Flowing               MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Flowing"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34MaintainedPoolOnly    MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Maintained"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34UnmaintainedPoolOnly  MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Unmaintained"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34HighOrganicDirtyWater MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "High Organic"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Unknown               MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Unknown"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34StagnantWaterBlocked  MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Stagnant"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34NeedsMonitoring       MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Needs Monitoring"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34DryingOut             MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Drying Out"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34AppearsVacant         MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Appears Vacant"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34EntryDenied           MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Entry Denied"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34PoolRemoved           MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "Pool Removed"
	MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34FalsePoolNotasource   MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type = "False Pool"
)

type MosquitoInspectionMosquitoBreedingType string

const (
	MosquitoInspectionMosquitoBreedingNone     MosquitoInspectionMosquitoBreedingType = "None"
	MosquitoInspectionMosquitoBreedingLight    MosquitoInspectionMosquitoBreedingType = "Light"
	MosquitoInspectionMosquitoBreedingModerate MosquitoInspectionMosquitoBreedingType = "Moderate"
	MosquitoInspectionMosquitoBreedingIntense  MosquitoInspectionMosquitoBreedingType = "Intense"
)

type MosquitoInspection struct {
	Objectid               uint                                                                                `field:"OBJECTID"`
	Dips                   int16                                                                               `field:"NUMDIPS"`
	Activity               MosquitoInspectionMosquitoActivityType                                              `field:"ACTIVITY"`
	Breeding               MosquitoInspectionMosquitoBreedingType                                              `field:"BREEDING"`
	TotalLarvae            int16                                                                               `field:"TOTLARVAE"`
	TotalPupae             int16                                                                               `field:"TOTPUPAE"`
	Eggs                   int16                                                                               `field:"EGGS"`
	PositiveDips           int16                                                                               `field:"POSDIPS"`
	AdultActivity          MosquitoInspectionMosquitoAdultActivityType                                         `field:"ADULTACT"`
	LarvalStages           string                                                                              `field:"LSTAGES"`
	DominantStage          MosquitoInspectionMosquitoInspectionDOMSTAGEb7a6c36bccde49a292020de4812cf5aeType    `field:"DOMSTAGE"`
	Action                 MosquitoInspectionMosquitoInspectionACTIONTAKEN252243d69b0b44ddbdc229c04ec3a8d5Type `field:"ACTIONTAKEN"`
	Comments               string                                                                              `field:"COMMENTS"`
	AverageTemperature     float64                                                                             `field:"AVETEMP"`
	WindSpeed              float64                                                                             `field:"WINDSPEED"`
	RainGauge              float64                                                                             `field:"RAINGAUGE"`
	Start                  time.Time                                                                           `field:"STARTDATETIME"`
	Finish                 time.Time                                                                           `field:"ENDDATETIME"`
	WindDirection          MosquitoInspectionNotInUIWindDirectionType                                          `field:"WINDDIR"`
	AverageLarvae          float64                                                                             `field:"AVGLARVAE"`
	AveragePupae           float64                                                                             `field:"AVGPUPAE"`
	Reviewed               MosquitoInspectionNotInUITFType                                                     `field:"REVIEWED"`
	ReviewedBy             string                                                                              `field:"REVIEWEDBY"`
	ReviewedDate           time.Time                                                                           `field:"REVIEWEDDATE"`
	LocationName           string                                                                              `field:"LOCATIONNAME"`
	Zone                   string                                                                              `field:"ZONE"`
	RecordStatus           int16                                                                               `field:"RECORDSTATUS"`
	Zone2                  string                                                                              `field:"ZONE2"`
	PersonalContact        MosquitoInspectionNotInUITFType                                                     `field:"PERSONALCONTACT"`
	TireCount              int16                                                                               `field:"TIRECOUNT"`
	CatchBasinCount        int16                                                                               `field:"CBCOUNT"`
	ContainerCount         int16                                                                               `field:"CONTAINERCOUNT"`
	FieldSpecies           MosquitoInspectionMosquitoFieldSpeciesType                                          `field:"FIELDSPECIES"`
	GlobalID               uuid.UUID                                                                           `field:"GlobalID"`
	CreatedUser            string                                                                              `field:"created_user"`
	CreatedDate            time.Time                                                                           `field:"created_date"`
	LastEditedUser         string                                                                              `field:"last_edited_user"`
	LastEditedDate         time.Time                                                                           `field:"last_edited_date"`
	Linelocid              uuid.UUID                                                                           `field:"LINELOCID"`
	Pointlocid             uuid.UUID                                                                           `field:"POINTLOCID"`
	Polygonlocid           uuid.UUID                                                                           `field:"POLYGONLOCID"`
	Srid                   uuid.UUID                                                                           `field:"SRID"`
	FieldTech              string                                                                              `field:"FIELDTECH"`
	LarvaePresent          MosquitoInspectionNotInUITFType                                                     `field:"LARVAEPRESENT"`
	PupaePresent           MosquitoInspectionNotInUITFType                                                     `field:"PUPAEPRESENT"`
	StormDrainId           uuid.UUID                                                                           `field:"SDID"`
	Conditions             MosquitoInspectionMosquitoInspectionSITECONDdb7350bc81e5401e858fcd3e5e5d8a34Type    `field:"SITECOND"`
	PositiveContainerCount int16                                                                               `field:"POSITIVECONTAINERCOUNT"`
	CreationDate           time.Time                                                                           `field:"CreationDate"`
	Creator                string                                                                              `field:"Creator"`
	EditDate               time.Time                                                                           `field:"EditDate"`
	Editor                 string                                                                              `field:"Editor"`
	Jurisdiction           string                                                                              `field:"JURISDICTION"`
	VisualMonitoring       MosquitoInspectionNotInUITFType                                                     `field:"VISUALMONITORING"`
	VmComments             string                                                                              `field:"VMCOMMENTS"`
	AdminAction            MosquitoInspectionMosquitoInspectionadminActionb74ae1bbc98b40f68cfa40e4fd16c270Type `field:"adminAction"`
	Ptaid                  uuid.UUID                                                                           `field:"PTAID"`
}
