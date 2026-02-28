package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type PoolBuffer struct {
	ObjectID          uint      `field:"OBJECTID"`
	Type              string    `field:"TYPE"`
	Name              string    `field:"NAME"`
	DateAdded         time.Time `field:"DATEADDED"`
	Comments          string    `field:"COMMENTS"`
	Zone              string    `field:"ZONE"`
	Zone2             string    `field:"ZONE2"`
	DoNotEnter        int16     `field:"DONOTENTER"`
	DoNotSpray        int16     `field:"DONOTSPRAY"`
	CustomerWarning   string    `field:"CUSTOMERWARNING"`
	TechnicianWarning string    `field:"TECHNICIANWARNING"`
	LastContactDate   time.Time `field:"LASTCONTACTDATE"`
	GlobalID          uuid.UUID `field:"GlobalID"`
	CreatedUser       string    `field:"created_user"`
	CreatedDate       time.Time `field:"created_date"`
	LastEditedUser    string    `field:"last_edited_user"`
	LastEditedDate    time.Time `field:"last_edited_date"`
	Address           string    `field:"ADDRESS"`
	PhoneNumber       string    `field:"PHONENUMBER"`
	CreationDate      time.Time `field:"CreationDate"`
	Creator           string    `field:"Creator"`
	EditDate          time.Time `field:"EditDate"`
	Editor            string    `field:"Editor"`
	ShapeArea         float64   `field:"Shape__Area"`
	ShapeLength       float64   `field:"Shape__Length"`
	Geometry          response.Geometry
}

func (x *PoolBuffer) GetGeometry() response.Geometry  { return x.Geometry }
func (x *PoolBuffer) SetGeometry(m response.Geometry) { x.Geometry = m }
