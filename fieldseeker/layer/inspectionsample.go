package layer

import (
	"time"

	"github.com/google/uuid"
)

type InspectionSampleNotInUITFType int16

const (
	InspectionSampleNotInUITFTrue  InspectionSampleNotInUITFType = 1
	InspectionSampleNotInUITFFalse InspectionSampleNotInUITFType = 0
)

type InspectionSample struct {
	ObjectID                    uint                          `field:"OBJECTID"`
	InspID                      uuid.UUID                     `field:"INSP_ID"`
	SampleID                    string                        `field:"SAMPLEID"`
	Processed                   InspectionSampleNotInUITFType `field:"PROCESSED"`
	TechIdentifyingSpeciesInLab string                        `field:"IDBYTECH"`
	GlobalID                    uuid.UUID                     `field:"GlobalID"`
	CreatedUser                 string                        `field:"created_user"`
	CreatedDate                 time.Time                     `field:"created_date"`
	LastEditedUser              string                        `field:"last_edited_user"`
	LastEditedDate              time.Time                     `field:"last_edited_date"`
	CreationDate                time.Time                     `field:"CreationDate"`
	Creator                     string                        `field:"Creator"`
	EditDate                    time.Time                     `field:"EditDate"`
	Editor                      string                        `field:"Editor"`
}
