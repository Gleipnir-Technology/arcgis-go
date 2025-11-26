package layer

import (
	"time"

	"github.com/google/uuid"
)

type RodentLocationRodentLocationSymbologyType string

const (
	RodentLocationRodentLocationSymbologyActionrequired   RodentLocationRodentLocationSymbologyType = "ACTION"
	RodentLocationRodentLocationSymbologyInactive         RodentLocationRodentLocationSymbologyType = "INACTIVE"
	RodentLocationRodentLocationSymbologyNoactionrequired RodentLocationRodentLocationSymbologyType = "NONE"
)

type RodentLocationRodentLocationHabitatType string

const (
	RodentLocationRodentLocationHabitatCommercial  RodentLocationRodentLocationHabitatType = "Commercial"
	RodentLocationRodentLocationHabitatIndustrial  RodentLocationRodentLocationHabitatType = "Industrial"
	RodentLocationRodentLocationHabitatResidential RodentLocationRodentLocationHabitatType = "Residential"
	RodentLocationRodentLocationHabitatWoodPile    RodentLocationRodentLocationHabitatType = "Wood Pile"
)

type RodentLocationLocationPriority1Type string

const (
	RodentLocationLocationPriority1Low    RodentLocationLocationPriority1Type = "Low"
	RodentLocationLocationPriority1Medium RodentLocationLocationPriority1Type = "Medium"
	RodentLocationLocationPriority1High   RodentLocationLocationPriority1Type = "High"
	RodentLocationLocationPriority1None   RodentLocationLocationPriority1Type = "None"
)

type RodentLocationLocationUseType1Type string

const (
	RodentLocationLocationUseType1Residential  RodentLocationLocationUseType1Type = "Residential"
	RodentLocationLocationUseType1Commercial   RodentLocationLocationUseType1Type = "Commercial"
	RodentLocationLocationUseType1Industrial   RodentLocationLocationUseType1Type = "Industrial"
	RodentLocationLocationUseType1Agricultural RodentLocationLocationUseType1Type = "Agricultural"
	RodentLocationLocationUseType1Mixeduse     RodentLocationLocationUseType1Type = "Mixed use"
)

type RodentLocationNotInUITF1Type int16

const (
	RodentLocationNotInUITF1True  RodentLocationNotInUITF1Type = 1
	RodentLocationNotInUITF1False RodentLocationNotInUITF1Type = 0
)

type RodentLocation struct {
	ObjectID                     uint                                      `field:"OBJECTID"`
	LocationName                 string                                    `field:"LOCATIONNAME"`
	Zone                         string                                    `field:"ZONE"`
	Zone2                        string                                    `field:"ZONE2"`
	Habitat                      RodentLocationRodentLocationHabitatType   `field:"HABITAT"`
	Priority                     RodentLocationLocationPriority1Type       `field:"PRIORITY"`
	Usetype                      RodentLocationLocationUseType1Type        `field:"USETYPE"`
	Active                       RodentLocationNotInUITF1Type              `field:"ACTIVE"`
	Description                  string                                    `field:"DESCRIPTION"`
	Accessdesc                   string                                    `field:"ACCESSDESC"`
	Comments                     string                                    `field:"COMMENTS"`
	Symbology                    RodentLocationRodentLocationSymbologyType `field:"SYMBOLOGY"`
	ExternalID                   string                                    `field:"EXTERNALID"`
	Nextactiondatescheduled      time.Time                                 `field:"NEXTACTIONDATESCHEDULED"`
	Locationnumber               int32                                     `field:"LOCATIONNUMBER"`
	LastInspectionDate           time.Time                                 `field:"LASTINSPECTDATE"`
	LastInspectionSpecies        string                                    `field:"LASTINSPECTSPECIES"`
	LastInspectionAction         string                                    `field:"LASTINSPECTACTION"`
	LastInspectionConditions     string                                    `field:"LASTINSPECTCONDITIONS"`
	LastInspectionRodentEvidence string                                    `field:"LASTINSPECTRODENTEVIDENCE"`
	GlobalID                     uuid.UUID                                 `field:"GlobalID"`
	CreatedUser                  string                                    `field:"created_user"`
	CreatedDate                  time.Time                                 `field:"created_date"`
	LastEditedUser               string                                    `field:"last_edited_user"`
	LastEditedDate               time.Time                                 `field:"last_edited_date"`
	CreationDate                 time.Time                                 `field:"CreationDate"`
	Creator                      string                                    `field:"Creator"`
	EditDate                     time.Time                                 `field:"EditDate"`
	Editor                       string                                    `field:"Editor"`
	Jurisdiction                 string                                    `field:"JURISDICTION"`
}
