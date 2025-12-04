package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type StormDrain struct {
	ObjectID          uint      `field:"OBJECTID"`
	NextTreatmentDate time.Time `field:"NextTreatmentDate"`
	LastTreatDate     time.Time `field:"LastTreatDate"`
	LastAction        string    `field:"LastAction"`
	Symbology         string    `field:"Symbology"`
	GlobalID          uuid.UUID `field:"GlobalID"`
	CreatedUser       string    `field:"created_user"`
	CreatedDate       time.Time `field:"created_date"`
	LastEditedUser    string    `field:"last_edited_user"`
	LastEditedDate    time.Time `field:"last_edited_date"`
	LastStatus        string    `field:"LastStatus"`
	Zone              string    `field:"ZONE"`
	Zone2             string    `field:"ZONE2"`
	CreationDate      time.Time `field:"CreationDate"`
	Creator           string    `field:"Creator"`
	EditDate          time.Time `field:"EditDate"`
	Editor            string    `field:"Editor"`
	Type              string    `field:"TYPE"`
	Jurisdiction      string    `field:"JURISDICTION"`
	Geometry          json.RawMessage
}

func (x *StormDrain) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *StormDrain) SetGeometry(m json.RawMessage) { x.Geometry = m }
