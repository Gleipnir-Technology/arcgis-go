package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type BarrierSpray struct {
	ObjectID       uint      `field:"OBJECTID"`
	Comments       string    `field:"COMMENTS"`
	Zone           string    `field:"ZONE"`
	Zone2          string    `field:"ZONE2"`
	RouteName      string    `field:"NAME"`
	Acres          float64   `field:"ACRES"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	Scheduled      int16     `field:"SCHEDULED"`
	ScheduledDate  time.Time `field:"SCHEDULEDDATE"`
	Notified       int16     `field:"Notified"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	ShapeArea      float64   `field:"Shape__Area"`
	ShapeLength    float64   `field:"Shape__Length"`
	Geometry       response.Geometry
}

func (x *BarrierSpray) GetGeometry() response.Geometry  { return x.Geometry }
func (x *BarrierSpray) SetGeometry(m response.Geometry) { x.Geometry = m }
