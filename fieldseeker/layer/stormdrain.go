package layer

import (
	"time"

	"github.com/google/uuid"
)

type StormDrainStormDrainSymbologyType string

const (
	StormDrainStormDrainSymbologyDry            StormDrainStormDrainSymbologyType = "Dry"
	StormDrainStormDrainSymbologyNeedsTreatment StormDrainStormDrainSymbologyType = "Needs Treatment"
	StormDrainStormDrainSymbologyTreated        StormDrainStormDrainSymbologyType = "Treated"
)

type StormDrain struct {
	ObjectID          uint                              `field:"OBJECTID"`
	NextTreatmentDate time.Time                         `field:"NextTreatmentDate"`
	LastTreatDate     time.Time                         `field:"LastTreatDate"`
	LastAction        string                            `field:"LastAction"`
	Symbology         StormDrainStormDrainSymbologyType `field:"Symbology"`
	GlobalID          uuid.UUID                         `field:"GlobalID"`
	CreatedUser       string                            `field:"created_user"`
	CreatedDate       time.Time                         `field:"created_date"`
	LastEditedUser    string                            `field:"last_edited_user"`
	LastEditedDate    time.Time                         `field:"last_edited_date"`
	LastStatus        string                            `field:"LastStatus"`
	Zone              string                            `field:"ZONE"`
	Zone2             string                            `field:"ZONE2"`
	CreationDate      time.Time                         `field:"CreationDate"`
	Creator           string                            `field:"Creator"`
	EditDate          time.Time                         `field:"EditDate"`
	Editor            string                            `field:"Editor"`
	Type              string                            `field:"TYPE"`
	Jurisdiction      string                            `field:"JURISDICTION"`
}
