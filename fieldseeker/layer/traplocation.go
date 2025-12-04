package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type TrapLocation struct {
	ObjectID            uint      `field:"OBJECTID"`
	Name                string    `field:"NAME"`
	Zone                string    `field:"ZONE"`
	Habitat             string    `field:"HABITAT"`
	Priority            string    `field:"PRIORITY"`
	UseType             string    `field:"USETYPE"`
	Active              int16     `field:"ACTIVE"`
	Description         string    `field:"DESCRIPTION"`
	AccessDescription   string    `field:"ACCESSDESC"`
	Comments            string    `field:"COMMENTS"`
	ExternalID          string    `field:"EXTERNALID"`
	NextScheduledAction time.Time `field:"NEXTACTIONDATESCHEDULED"`
	Zone2               string    `field:"ZONE2"`
	Locationnumber      int32     `field:"LOCATIONNUMBER"`
	GlobalID            uuid.UUID `field:"GlobalID"`
	CreatedUser         string    `field:"created_user"`
	CreatedDate         time.Time `field:"created_date"`
	LastEditedUser      string    `field:"last_edited_user"`
	LastEditedDate      time.Time `field:"last_edited_date"`
	GatewaySync         int16     `field:"GATEWAYSYNC"`
	Route               int32     `field:"route"`
	SetDayOfWeek        int32     `field:"set_dow"`
	RouteOrder          int32     `field:"route_order"`
	VectorsurvsiteID    string    `field:"VECTORSURVSITEID"`
	CreationDate        time.Time `field:"CreationDate"`
	Creator             string    `field:"Creator"`
	EditDate            time.Time `field:"EditDate"`
	Editor              string    `field:"Editor"`
	H3r7                string    `field:"h3r7"`
	H3r8                string    `field:"h3r8"`
	Geometry            json.RawMessage
}

func (x *TrapLocation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *TrapLocation) SetGeometry(m json.RawMessage) { x.Geometry = m }
