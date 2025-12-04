package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type InspectionSampleDetail struct {
	ObjectID           uint      `field:"OBJECTID"`
	InspsampleID       uuid.UUID `field:"INSPSAMPLE_ID"`
	FieldSpecies       string    `field:"FIELDSPECIES"`
	FieldLarvaCount    int16     `field:"FLARVCOUNT"`
	FieldPupaCount     int16     `field:"FPUPCOUNT"`
	FieldEggCount      int16     `field:"FEGGCOUNT"`
	FieldLarvalStages  string    `field:"FLSTAGES"`
	FieldDominantStage string    `field:"FDOMSTAGE"`
	FieldAdultActivity string    `field:"FADULTACT"`
	LabSpecies         string    `field:"LABSPECIES"`
	LabLarvaCount      int16     `field:"LLARVCOUNT"`
	LabPupaCount       int16     `field:"LPUPCOUNT"`
	LabEggCount        int16     `field:"LEGGCOUNT"`
	LabDominantStage   string    `field:"LDOMSTAGE"`
	Comments           string    `field:"COMMENTS"`
	GlobalID           uuid.UUID `field:"GlobalID"`
	CreatedUser        string    `field:"created_user"`
	CreatedDate        time.Time `field:"created_date"`
	LastEditedUser     string    `field:"last_edited_user"`
	LastEditedDate     time.Time `field:"last_edited_date"`
	Processed          int16     `field:"PROCESSED"`
	CreationDate       time.Time `field:"CreationDate"`
	Creator            string    `field:"Creator"`
	EditDate           time.Time `field:"EditDate"`
	Editor             string    `field:"Editor"`
	Geometry           json.RawMessage
}

func (x *InspectionSampleDetail) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *InspectionSampleDetail) SetGeometry(m json.RawMessage) { x.Geometry = m }
