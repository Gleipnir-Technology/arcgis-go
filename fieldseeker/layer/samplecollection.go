package layer

import (
	"time"

	"github.com/google/uuid"
)

type SampleCollectionMosquitoSiteConditionType string

const (
	SampleCollectionMosquitoSiteConditionDry   SampleCollectionMosquitoSiteConditionType = "Dry"
	SampleCollectionMosquitoSiteConditionClean SampleCollectionMosquitoSiteConditionType = "Clean"
	SampleCollectionMosquitoSiteConditionFull  SampleCollectionMosquitoSiteConditionType = "Full"
	SampleCollectionMosquitoSiteConditionLow   SampleCollectionMosquitoSiteConditionType = "Low"
)

type SampleCollectionMosquitoSampleTypeType string

const (
	SampleCollectionMosquitoSampleTypeBlood    SampleCollectionMosquitoSampleTypeType = "Blood"
	SampleCollectionMosquitoSampleTypeTissue   SampleCollectionMosquitoSampleTypeType = "Tissue"
	SampleCollectionMosquitoSampleTypeSpecimen SampleCollectionMosquitoSampleTypeType = "Specimen"
	SampleCollectionMosquitoSampleTypeCarcass  SampleCollectionMosquitoSampleTypeType = "Carcass"
)

type SampleCollectionMosquitoActivityType string

const (
	SampleCollectionMosquitoActivityRoutineinspection SampleCollectionMosquitoActivityType = "Routine inspection"
	SampleCollectionMosquitoActivityPretreatment      SampleCollectionMosquitoActivityType = "Pre-treatment"
	SampleCollectionMosquitoActivityMaintenance       SampleCollectionMosquitoActivityType = "Maintenance"
	SampleCollectionMosquitoActivityULV               SampleCollectionMosquitoActivityType = "ULV"
	SampleCollectionMosquitoActivityBARRIER           SampleCollectionMosquitoActivityType = "BARRIER"
	SampleCollectionMosquitoActivityLOGIN             SampleCollectionMosquitoActivityType = "LOGIN"
	SampleCollectionMosquitoActivityTREATSD           SampleCollectionMosquitoActivityType = "TREATSD"
	SampleCollectionMosquitoActivitySD                SampleCollectionMosquitoActivityType = "SD"
	SampleCollectionMosquitoActivitySITEVISIT         SampleCollectionMosquitoActivityType = "SITEVISIT"
	SampleCollectionMosquitoActivityONLINE            SampleCollectionMosquitoActivityType = "ONLINE"
	SampleCollectionMosquitoActivitySYNC              SampleCollectionMosquitoActivityType = "SYNC"
	SampleCollectionMosquitoActivityCREATESR          SampleCollectionMosquitoActivityType = "CREATESR"
	SampleCollectionMosquitoActivityLC                SampleCollectionMosquitoActivityType = "LC"
	SampleCollectionMosquitoActivityACCEPTSR          SampleCollectionMosquitoActivityType = "ACCEPTSR"
	SampleCollectionMosquitoActivityPOINT             SampleCollectionMosquitoActivityType = "POINT"
	SampleCollectionMosquitoActivityDOWNLOAD          SampleCollectionMosquitoActivityType = "DOWNLOAD"
	SampleCollectionMosquitoActivityCOMPLETESR        SampleCollectionMosquitoActivityType = "COMPLETESR"
	SampleCollectionMosquitoActivityPOLYGON           SampleCollectionMosquitoActivityType = "POLYGON"
	SampleCollectionMosquitoActivityTRAP              SampleCollectionMosquitoActivityType = "TRAP"
	SampleCollectionMosquitoActivitySAMPLE            SampleCollectionMosquitoActivityType = "SAMPLE"
	SampleCollectionMosquitoActivityQA                SampleCollectionMosquitoActivityType = "QA"
	SampleCollectionMosquitoActivityPTA               SampleCollectionMosquitoActivityType = "PTA"
	SampleCollectionMosquitoActivityFIELDSCOUTING     SampleCollectionMosquitoActivityType = "FIELDSCOUTING"
	SampleCollectionMosquitoActivityOFFLINE           SampleCollectionMosquitoActivityType = "OFFLINE"
	SampleCollectionMosquitoActivityLINE              SampleCollectionMosquitoActivityType = "LINE"
	SampleCollectionMosquitoActivityTRAPLOCATION      SampleCollectionMosquitoActivityType = "TRAPLOCATION"
	SampleCollectionMosquitoActivitySAMPLELOCATION    SampleCollectionMosquitoActivityType = "SAMPLELOCATION"
	SampleCollectionMosquitoActivityLCLOCATION        SampleCollectionMosquitoActivityType = "LCLOCATION"
)

type SampleCollectionMosquitoDiseaseType string

const (
	SampleCollectionMosquitoDiseaseEEE    SampleCollectionMosquitoDiseaseType = "EEE"
	SampleCollectionMosquitoDiseaseWNV    SampleCollectionMosquitoDiseaseType = "WNV"
	SampleCollectionMosquitoDiseaseDengue SampleCollectionMosquitoDiseaseType = "Dengue"
	SampleCollectionMosquitoDiseaseZika   SampleCollectionMosquitoDiseaseType = "Zika"
)

type SampleCollectionNotInUITFType int16

const (
	SampleCollectionNotInUITFTrue  SampleCollectionNotInUITFType = 1
	SampleCollectionNotInUITFFalse SampleCollectionNotInUITFType = 0
)

type SampleCollectionMosquitoSampleConditionType string

const (
	SampleCollectionMosquitoSampleConditionUsable   SampleCollectionMosquitoSampleConditionType = "Usable"
	SampleCollectionMosquitoSampleConditionUnusable SampleCollectionMosquitoSampleConditionType = "Unusable"
)

type SampleCollectionMosquitoSampleSpeciesType string

const (
	SampleCollectionMosquitoSampleSpeciesChicken  SampleCollectionMosquitoSampleSpeciesType = "Chicken"
	SampleCollectionMosquitoSampleSpeciesWildbird SampleCollectionMosquitoSampleSpeciesType = "Wild bird"
	SampleCollectionMosquitoSampleSpeciesHorse    SampleCollectionMosquitoSampleSpeciesType = "Horse"
	SampleCollectionMosquitoSampleSpeciesHuman    SampleCollectionMosquitoSampleSpeciesType = "Human"
)

type SampleCollectionNotInUISexType string

const (
	SampleCollectionNotInUISexMale    SampleCollectionNotInUISexType = "M"
	SampleCollectionNotInUISexFemale  SampleCollectionNotInUISexType = "F"
	SampleCollectionNotInUISexUnknown SampleCollectionNotInUISexType = "U"
)

type SampleCollectionNotInUIWindDirectionType string

const (
	SampleCollectionNotInUIWindDirectionN  SampleCollectionNotInUIWindDirectionType = "N"
	SampleCollectionNotInUIWindDirectionNE SampleCollectionNotInUIWindDirectionType = "NE"
	SampleCollectionNotInUIWindDirectionE  SampleCollectionNotInUIWindDirectionType = "E"
	SampleCollectionNotInUIWindDirectionSE SampleCollectionNotInUIWindDirectionType = "SE"
	SampleCollectionNotInUIWindDirectionS  SampleCollectionNotInUIWindDirectionType = "S"
	SampleCollectionNotInUIWindDirectionSW SampleCollectionNotInUIWindDirectionType = "SW"
	SampleCollectionNotInUIWindDirectionW  SampleCollectionNotInUIWindDirectionType = "W"
	SampleCollectionNotInUIWindDirectionNW SampleCollectionNotInUIWindDirectionType = "NW"
)

type SampleCollectionMosquitoTestMethodType string

const (
	SampleCollectionMosquitoTestMethodRAMP    SampleCollectionMosquitoTestMethodType = "RAMP"
	SampleCollectionMosquitoTestMethodVecTest SampleCollectionMosquitoTestMethodType = "VecTest"
	SampleCollectionMosquitoTestMethodELISA   SampleCollectionMosquitoTestMethodType = "ELISA"
	SampleCollectionMosquitoTestMethodRTPCR   SampleCollectionMosquitoTestMethodType = "RT-PCR"
)

type SampleCollectionMosquitoLabNameType string

const (
	SampleCollectionMosquitoLabNameInternalLab SampleCollectionMosquitoLabNameType = "Internal Lab"
	SampleCollectionMosquitoLabNameStateLab    SampleCollectionMosquitoLabNameType = "State Lab"
)

type SampleCollection struct {
	ObjectID               uint                                        `field:"OBJECTID"`
	LocID                  uuid.UUID                                   `field:"LOC_ID"`
	Start                  time.Time                                   `field:"STARTDATETIME"`
	Finish                 time.Time                                   `field:"ENDDATETIME"`
	Conditions             SampleCollectionMosquitoSiteConditionType   `field:"SITECOND"`
	SampleID               string                                      `field:"SAMPLEID"`
	SurveillanceTechnician string                                      `field:"SURVTECH"`
	Sent                   time.Time                                   `field:"DATESENT"`
	Tested                 time.Time                                   `field:"DATETESTED"`
	TestTechnician         string                                      `field:"TESTTECH"`
	Comments               string                                      `field:"COMMENTS"`
	Processed              SampleCollectionNotInUITFType               `field:"PROCESSED"`
	SampleType             SampleCollectionMosquitoSampleTypeType      `field:"SAMPLETYPE"`
	SampleCondition        SampleCollectionMosquitoSampleConditionType `field:"SAMPLECOND"`
	Species                SampleCollectionMosquitoSampleSpeciesType   `field:"SPECIES"`
	Sex                    SampleCollectionNotInUISexType              `field:"SEX"`
	AverageTemperature     float64                                     `field:"AVETEMP"`
	WindSpeed              float64                                     `field:"WINDSPEED"`
	WindDirection          SampleCollectionNotInUIWindDirectionType    `field:"WINDDIR"`
	RainGauge              float64                                     `field:"RAINGAUGE"`
	Activity               SampleCollectionMosquitoActivityType        `field:"ACTIVITY"`
	TestMethod             SampleCollectionMosquitoTestMethodType      `field:"TESTMETHOD"`
	DiseaseTested          SampleCollectionMosquitoDiseaseType         `field:"DISEASETESTED"`
	DiseasePositive        SampleCollectionMosquitoDiseaseType         `field:"DISEASEPOS"`
	Reviewed               SampleCollectionNotInUITFType               `field:"REVIEWED"`
	ReviewedBy             string                                      `field:"REVIEWEDBY"`
	ReviewedDate           time.Time                                   `field:"REVIEWEDDATE"`
	LocationName           string                                      `field:"LOCATIONNAME"`
	Zone                   string                                      `field:"ZONE"`
	RecordStatus           int16                                       `field:"RECORDSTATUS"`
	Zone2                  string                                      `field:"ZONE2"`
	GlobalID               uuid.UUID                                   `field:"GlobalID"`
	CreatedUser            string                                      `field:"created_user"`
	CreatedDate            time.Time                                   `field:"created_date"`
	LastEditedUser         string                                      `field:"last_edited_user"`
	LastEditedDate         time.Time                                   `field:"last_edited_date"`
	Lab                    SampleCollectionMosquitoLabNameType         `field:"LAB"`
	FieldTech              string                                      `field:"FIELDTECH"`
	FlockID                uuid.UUID                                   `field:"FLOCKID"`
	SampleCount            int16                                       `field:"SAMPLECOUNT"`
	ChickenID              uuid.UUID                                   `field:"CHICKENID"`
	GatewaySync            int16                                       `field:"GATEWAYSYNC"`
	CreationDate           time.Time                                   `field:"CreationDate"`
	Creator                string                                      `field:"Creator"`
	EditDate               time.Time                                   `field:"EditDate"`
	Editor                 string                                      `field:"Editor"`
}
