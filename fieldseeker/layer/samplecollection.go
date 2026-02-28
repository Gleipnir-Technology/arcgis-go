package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type SampleCollection struct {
	ObjectID       uint      `field:"OBJECTID"`
	InspsampleID   uuid.UUID `field:"QAINSP_ID"`
	FieldSpecies   string    `field:"FIELDSPECIES"`
	Stage1Count    int16     `field:"FLARVSTG1COUNT"`
	FieldPupaCount int16     `field:"FPUPCOUNT"`
	FieldEggCount  int16     `field:"FEGGCOUNT"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	Comments       string    `field:"COMMENTS"`
	Stage2Count    int16     `field:"FLARVSTG2COUNT"`
	Stage3Count    int16     `field:"FLARVSTG3COUNT"`
	Stage4Count    int16     `field:"FLARVSTG4COUNT"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       response.Geometry
}

func (x *SampleCollection) GetGeometry() response.Geometry  { return x.Geometry }
func (x *SampleCollection) SetGeometry(m response.Geometry) { x.Geometry = m }
