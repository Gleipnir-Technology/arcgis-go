package layer

import (
	"time"

	"github.com/google/uuid"
)

type SampleLocationLocationHabitatTypeType string

const (
	SampleLocationLocationHabitatTypeCatchbasin     SampleLocationLocationHabitatTypeType = "Catch basin"
	SampleLocationLocationHabitatTypeCreek          SampleLocationLocationHabitatTypeType = "Creek"
	SampleLocationLocationHabitatTypeDitch          SampleLocationLocationHabitatTypeType = "Ditch"
	SampleLocationLocationHabitatTypeFieldPasture   SampleLocationLocationHabitatTypeType = "Field/Pasture"
	SampleLocationLocationHabitatTypePond           SampleLocationLocationHabitatTypeType = "Pond"
	SampleLocationLocationHabitatTypePondfish       SampleLocationLocationHabitatTypeType = "Pond fish"
	SampleLocationLocationHabitatTypePondmarshy     SampleLocationLocationHabitatTypeType = "Pond marshy"
	SampleLocationLocationHabitatTypePondornamental SampleLocationLocationHabitatTypeType = "Pond ornamental"
	SampleLocationLocationHabitatTypePondretention  SampleLocationLocationHabitatTypeType = "Pond retention"
	SampleLocationLocationHabitatTypePondsewage     SampleLocationLocationHabitatTypeType = "Pond sewage"
	SampleLocationLocationHabitatTypePondwoodland   SampleLocationLocationHabitatTypeType = "Pond woodland"
	SampleLocationLocationHabitatTypeTreehole       SampleLocationLocationHabitatTypeType = "Tree hole"
	SampleLocationLocationHabitatTypeSwimmingpool   SampleLocationLocationHabitatTypeType = "Swimming pool"
	SampleLocationLocationHabitatTypePark           SampleLocationLocationHabitatTypeType = "Park"
	SampleLocationLocationHabitatTypeUnknown        SampleLocationLocationHabitatTypeType = "Unknown"
)

type SampleLocationLocationPriorityType string

const (
	SampleLocationLocationPriorityLow    SampleLocationLocationPriorityType = "Low"
	SampleLocationLocationPriorityMedium SampleLocationLocationPriorityType = "Medium"
	SampleLocationLocationPriorityHigh   SampleLocationLocationPriorityType = "High"
	SampleLocationLocationPriorityNone   SampleLocationLocationPriorityType = "None"
)

type SampleLocationSampleLocationUseTypeType string

const (
	SampleLocationSampleLocationUseTypeFlockSite SampleLocationSampleLocationUseTypeType = "Flock Site"
	SampleLocationSampleLocationUseTypeDeadBird  SampleLocationSampleLocationUseTypeType = "Dead Bird"
)

type SampleLocationNotInUITFType int16

const (
	SampleLocationNotInUITFTrue  SampleLocationNotInUITFType = 1
	SampleLocationNotInUITFFalse SampleLocationNotInUITFType = 0
)

type SampleLocation struct {
	Objectid            uint                                    `field:"OBJECTID"`
	Name                string                                  `field:"NAME"`
	Zone                string                                  `field:"ZONE"`
	Habitat             SampleLocationLocationHabitatTypeType   `field:"HABITAT"`
	Priority            SampleLocationLocationPriorityType      `field:"PRIORITY"`
	UseType             SampleLocationSampleLocationUseTypeType `field:"USETYPE"`
	Active              SampleLocationNotInUITFType             `field:"ACTIVE"`
	Description         string                                  `field:"DESCRIPTION"`
	AccessDescription   string                                  `field:"ACCESSDESC"`
	Comments            string                                  `field:"COMMENTS"`
	ExternalId          string                                  `field:"EXTERNALID"`
	NextScheduledAction time.Time                               `field:"NEXTACTIONDATESCHEDULED"`
	Zone2               string                                  `field:"ZONE2"`
	Locationnumber      int32                                   `field:"LOCATIONNUMBER"`
	GlobalID            uuid.UUID                               `field:"GlobalID"`
	CreatedUser         string                                  `field:"created_user"`
	CreatedDate         time.Time                               `field:"created_date"`
	LastEditedUser      string                                  `field:"last_edited_user"`
	LastEditedDate      time.Time                               `field:"last_edited_date"`
	GatewaySync         int16                                   `field:"GATEWAYSYNC"`
	CreationDate        time.Time                               `field:"CreationDate"`
	Creator             string                                  `field:"Creator"`
	EditDate            time.Time                               `field:"EditDate"`
	Editor              string                                  `field:"Editor"`
}
