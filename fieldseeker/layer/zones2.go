package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type Zones2 struct {
	ObjectID       uint      `field:"OBJECTID"`
	SystemName     string    `field:"SystemName"`
	Dimension      string    `field:"Dimension"`
	UnitName       string    `field:"UnitName"`
	UnitValue      string    `field:"UnitValue"`
	Description    string    `field:"Description"`
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

func (x *Zones2) GetGeometry() response.Geometry  { return x.Geometry }
func (x *Zones2) SetGeometry(m response.Geometry) { x.Geometry = m }
