package layer

import (
	"time"

	"github.com/google/uuid"
)

type TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type string

const (
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166E  TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "E"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166N  TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "N"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166NE TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "NE"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166NW TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "NW"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166S  TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "S"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166SE TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "SE"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166SW TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "SW"
	TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166W  TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type = "W"
)

type TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType string

const (
	TrapDataTrapDataLure25fe542f077f4254868176e8f436354bCO2DryIce     TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType = "CO2 (Dry Ice)"
	TrapDataTrapDataLure25fe542f077f4254868176e8f436354bCO2SugarYeast TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType = "CO2 (Sugar Yeast)"
	TrapDataTrapDataLure25fe542f077f4254868176e8f436354bBGLure        TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType = "BG-Lure"
	TrapDataTrapDataLure25fe542f077f4254868176e8f436354bGravidWater   TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType = "Gravid Water"
)

type TrapDataMosquitoTrapTypeType string

const (
	TrapDataMosquitoTrapTypeGravidTrap    TrapDataMosquitoTrapTypeType = "GRVD"
	TrapDataMosquitoTrapTypeBGSENT        TrapDataMosquitoTrapTypeType = "BGSENT"
	TrapDataMosquitoTrapTypeCO2baitedtrap TrapDataMosquitoTrapTypeType = "CO2"
)

type TrapDataNotInUITrapActivityTypeType string

const (
	TrapDataNotInUITrapActivityTypeSet      TrapDataNotInUITrapActivityTypeType = "S"
	TrapDataNotInUITrapActivityTypeRetrieve TrapDataNotInUITrapActivityTypeType = "R"
)

type TrapDataNotInUITFType int16

const (
	TrapDataNotInUITFTrue  TrapDataNotInUITFType = 1
	TrapDataNotInUITFFalse TrapDataNotInUITFType = 0
)

type TrapDataMosquitoSiteConditionType string

const (
	TrapDataMosquitoSiteConditionDry   TrapDataMosquitoSiteConditionType = "Dry"
	TrapDataMosquitoSiteConditionClean TrapDataMosquitoSiteConditionType = "Clean"
	TrapDataMosquitoSiteConditionFull  TrapDataMosquitoSiteConditionType = "Full"
	TrapDataMosquitoSiteConditionLow   TrapDataMosquitoSiteConditionType = "Low"
)

type TrapDataMosquitoTrapConditionType string

const (
	TrapDataMosquitoTrapConditionDamaged TrapDataMosquitoTrapConditionType = "Damaged"
	TrapDataMosquitoTrapConditionMissing TrapDataMosquitoTrapConditionType = "Missing"
	TrapDataMosquitoTrapConditionFanOff  TrapDataMosquitoTrapConditionType = "Fan Off"
	TrapDataMosquitoTrapConditionFanSlow TrapDataMosquitoTrapConditionType = "Fan Slow"
)

type TrapData struct {
	Objectid                    uint                                                        `field:"OBJECTID"`
	TrapType                    TrapDataMosquitoTrapTypeType                                `field:"TRAPTYPE"`
	TrapActivityType            TrapDataNotInUITrapActivityTypeType                         `field:"TRAPACTIVITYTYPE"`
	Start                       time.Time                                                   `field:"STARTDATETIME"`
	Finish                      time.Time                                                   `field:"ENDDATETIME"`
	Comments                    string                                                      `field:"COMMENTS"`
	TechIdentifyingSpeciesInLab string                                                      `field:"IDBYTECH"`
	TechSortingTrapResultsInLab string                                                      `field:"SORTBYTECH"`
	Processed                   TrapDataNotInUITFType                                       `field:"PROCESSED"`
	SiteConditions              TrapDataMosquitoSiteConditionType                           `field:"SITECOND"`
	LocationName                string                                                      `field:"LOCATIONNAME"`
	RecordStatus                int16                                                       `field:"RECORDSTATUS"`
	Reviewed                    TrapDataNotInUITFType                                       `field:"REVIEWED"`
	ReviewedBy                  string                                                      `field:"REVIEWEDBY"`
	ReviewedDate                time.Time                                                   `field:"REVIEWEDDATE"`
	TrapCondition               TrapDataMosquitoTrapConditionType                           `field:"TRAPCONDITION"`
	TrapNights                  int16                                                       `field:"TRAPNIGHTS"`
	Zone                        string                                                      `field:"ZONE"`
	Zone2                       string                                                      `field:"ZONE2"`
	GlobalID                    uuid.UUID                                                   `field:"GlobalID"`
	CreatedUser                 string                                                      `field:"created_user"`
	CreatedDate                 time.Time                                                   `field:"created_date"`
	LastEditedUser              string                                                      `field:"last_edited_user"`
	LastEditedDate              time.Time                                                   `field:"last_edited_date"`
	Srid                        uuid.UUID                                                   `field:"SRID"`
	FieldTech                   string                                                      `field:"FIELDTECH"`
	GatewaySync                 int16                                                       `field:"GATEWAYSYNC"`
	LocId                       uuid.UUID                                                   `field:"LOC_ID"`
	Voltage                     float64                                                     `field:"VOLTAGE"`
	Winddir                     TrapDataTrapDataWINDDIRc1a31e05d0b94b228800be127bb3f166Type `field:"WINDDIR"`
	Windspeed                   float64                                                     `field:"WINDSPEED"`
	Avetemp                     float64                                                     `field:"AVETEMP"`
	Raingauge                   float64                                                     `field:"RAINGAUGE"`
	LandingRate                 int16                                                       `field:"LR"`
	Field                       int32                                                       `field:"Field"`
	Vectorsurvtrapdataid        string                                                      `field:"VECTORSURVTRAPDATAID"`
	Vectorsurvtraplocationid    string                                                      `field:"VECTORSURVTRAPLOCATIONID"`
	CreationDate                time.Time                                                   `field:"CreationDate"`
	Creator                     string                                                      `field:"Creator"`
	EditDate                    time.Time                                                   `field:"EditDate"`
	Editor                      string                                                      `field:"Editor"`
	Lure                        TrapDataTrapDataLure25fe542f077f4254868176e8f436354bType    `field:"Lure"`
}
