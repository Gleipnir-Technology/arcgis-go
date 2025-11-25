package layer

import (
	"time"

	"github.com/google/uuid"
)

type TrapLocationTrapLocationHABITAT5c349680f5ff40b1aeca88c17993e8f3Type string

const (
	TrapLocationTrapLocationHABITAT5c349680f5ff40b1aeca88c17993e8f3Trap TrapLocationTrapLocationHABITAT5c349680f5ff40b1aeca88c17993e8f3Type = "Trap"
)

type TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType string

const (
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fLow      TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "Low"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fMedium   TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "Medium"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fHigh     TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "High"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fNone     TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "None"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fProject  TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "Project"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fFixed    TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "Fixed"
	TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fResponse TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType = "Response "
)

type TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type string

const (
	TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193FixedTrapping    TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type = "Fixed Trapping"
	TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193ResponseTrapping TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type = "Response Trapping"
	TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193ServiceRequest   TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type = "Service Request"
	TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193ProjectTrap      TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type = "Project Trap"
)

type TrapLocationNotInUITFType int16

const (
	TrapLocationNotInUITFTrue  TrapLocationNotInUITFType = 1
	TrapLocationNotInUITFFalse TrapLocationNotInUITFType = 0
)

type TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aType string

const (
	TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aHomeownerpreference TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aType = "homeowner preference"
	TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aNolongerneeded      TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aType = "no longer needed"
)

type TrapLocation struct {
	Objectid            uint                                                                   `field:"OBJECTID"`
	Name                string                                                                 `field:"NAME"`
	Zone                string                                                                 `field:"ZONE"`
	Habitat             TrapLocationTrapLocationHABITAT5c349680f5ff40b1aeca88c17993e8f3Type    `field:"HABITAT"`
	Priority            TrapLocationTrapLocationPRIORITY680fb011063b41d59f39271c959b857fType   `field:"PRIORITY"`
	UseType             TrapLocationTrapLocationUSETYPE5e0eff9231fb404c98cc53c1d49a2193Type    `field:"USETYPE"`
	Active              TrapLocationNotInUITFType                                              `field:"ACTIVE"`
	Description         string                                                                 `field:"DESCRIPTION"`
	AccessDescription   TrapLocationTrapLocationACCESSDESC154cbd1045244e3a8ca0f099ec86556aType `field:"ACCESSDESC"`
	Comments            string                                                                 `field:"COMMENTS"`
	ExternalId          string                                                                 `field:"EXTERNALID"`
	NextScheduledAction time.Time                                                              `field:"NEXTACTIONDATESCHEDULED"`
	Zone2               string                                                                 `field:"ZONE2"`
	Locationnumber      int32                                                                  `field:"LOCATIONNUMBER"`
	GlobalID            uuid.UUID                                                              `field:"GlobalID"`
	CreatedUser         string                                                                 `field:"created_user"`
	CreatedDate         time.Time                                                              `field:"created_date"`
	LastEditedUser      string                                                                 `field:"last_edited_user"`
	LastEditedDate      time.Time                                                              `field:"last_edited_date"`
	GatewaySync         int16                                                                  `field:"GATEWAYSYNC"`
	Route               int32                                                                  `field:"route"`
	SetDayOfWeek        int32                                                                  `field:"set_dow"`
	RouteOrder          int32                                                                  `field:"route_order"`
	Vectorsurvsiteid    string                                                                 `field:"VECTORSURVSITEID"`
	CreationDate        time.Time                                                              `field:"CreationDate"`
	Creator             string                                                                 `field:"Creator"`
	EditDate            time.Time                                                              `field:"EditDate"`
	Editor              string                                                                 `field:"Editor"`
	H3r7                string                                                                 `field:"h3r7"`
	H3r8                string                                                                 `field:"h3r8"`
}
