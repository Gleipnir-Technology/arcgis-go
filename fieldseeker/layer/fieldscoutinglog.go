package layer

import (
	"time"

	"github.com/google/uuid"
)

type FieldScoutingLogFieldScoutingSymbologyType int16

const (
	FieldScoutingLogFieldScoutingSymbologyTrack         FieldScoutingLogFieldScoutingSymbologyType = 0
	FieldScoutingLogFieldScoutingSymbologyWetnolarvae   FieldScoutingLogFieldScoutingSymbologyType = 1
	FieldScoutingLogFieldScoutingSymbologyLarvaepresent FieldScoutingLogFieldScoutingSymbologyType = 2
	FieldScoutingLogFieldScoutingSymbologyDry           FieldScoutingLogFieldScoutingSymbologyType = 3
)

type FieldScoutingLog struct {
	ObjectID       uint                                       `field:"OBJECTID"`
	Status         FieldScoutingLogFieldScoutingSymbologyType `field:"STATUS"`
	GlobalID       uuid.UUID                                  `field:"GlobalID"`
	CreatedUser    string                                     `field:"created_user"`
	CreatedDate    time.Time                                  `field:"created_date"`
	LastEditedUser string                                     `field:"last_edited_user"`
	LastEditedDate time.Time                                  `field:"last_edited_date"`
	CreationDate   time.Time                                  `field:"CreationDate"`
	Creator        string                                     `field:"Creator"`
	EditDate       time.Time                                  `field:"EditDate"`
	Editor         string                                     `field:"Editor"`
}
