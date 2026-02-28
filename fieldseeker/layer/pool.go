package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type Pool struct {
	ObjectID                    uint      `field:"OBJECTID"`
	InspID                      uuid.UUID `field:"INSP_ID"`
	SampleID                    string    `field:"SAMPLEID"`
	Processed                   int16     `field:"PROCESSED"`
	TechIdentifyingSpeciesInLab string    `field:"IDBYTECH"`
	GlobalID                    uuid.UUID `field:"GlobalID"`
	CreatedUser                 string    `field:"created_user"`
	CreatedDate                 time.Time `field:"created_date"`
	LastEditedUser              string    `field:"last_edited_user"`
	LastEditedDate              time.Time `field:"last_edited_date"`
	CreationDate                time.Time `field:"CreationDate"`
	Creator                     string    `field:"Creator"`
	EditDate                    time.Time `field:"EditDate"`
	Editor                      string    `field:"Editor"`
	Geometry                    response.Geometry
}

func (x *Pool) GetGeometry() response.Geometry  { return x.Geometry }
func (x *Pool) SetGeometry(m response.Geometry) { x.Geometry = m }
