package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type QAProductObservation struct {
	ObjectID       uint      `field:"OBJECTID"`
	Status         int16     `field:"STATUS"`
	GlobalID       uuid.UUID `field:"GlobalID"`
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

func (x *QAProductObservation) GetGeometry() response.Geometry  { return x.Geometry }
func (x *QAProductObservation) SetGeometry(m response.Geometry) { x.Geometry = m }
