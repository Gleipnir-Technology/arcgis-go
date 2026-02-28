package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type Tracklog struct {
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
	CreationDate        time.Time `field:"CreationDate"`
	Creator             string    `field:"Creator"`
	EditDate            time.Time `field:"EditDate"`
	Editor              string    `field:"Editor"`
	Geometry            response.Geometry
}

func (x *Tracklog) GetGeometry() response.Geometry  { return x.Geometry }
func (x *Tracklog) SetGeometry(m response.Geometry) { x.Geometry = m }
