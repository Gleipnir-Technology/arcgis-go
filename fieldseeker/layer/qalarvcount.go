package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type QALarvCount struct {
	ObjectID       uint      `field:"OBJECTID"`
	TreatID        uuid.UUID `field:"TREAT_ID"`
	SessionID      uuid.UUID `field:"SESSION_ID"`
	TreatmentDate  time.Time `field:"TREATDATE"`
	Comments       string    `field:"COMMENTS"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	Type           string    `field:"TYPE"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	ShapeLength    float64   `field:"Shape__Length"`
	Geometry       response.Geometry
}

func (x *QALarvCount) GetGeometry() response.Geometry  { return x.Geometry }
func (x *QALarvCount) SetGeometry(m response.Geometry) { x.Geometry = m }
