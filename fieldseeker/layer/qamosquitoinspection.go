package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type QAMosquitoInspection struct {
	ObjectID       uint      `field:"OBJECTID"`
	RouteName      string    `field:"NAME"`
	Zone           string    `field:"ZONE"`
	Zone2          string    `field:"ZONE2"`
	ExternalID     string    `field:"EXTERNAL_ID"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	ShapeLength    float64   `field:"Shape__Length"`
	Geometry       response.Geometry
}

func (x *QAMosquitoInspection) GetGeometry() response.Geometry  { return x.Geometry }
func (x *QAMosquitoInspection) SetGeometry(m response.Geometry) { x.Geometry = m }
