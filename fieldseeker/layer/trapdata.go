package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type TrapData struct {
	ObjectID                    uint      `field:"OBJECTID"`
	TrapType                    string    `field:"TRAPTYPE"`
	TrapActivityType            string    `field:"TRAPACTIVITYTYPE"`
	Start                       time.Time `field:"STARTDATETIME"`
	Finish                      time.Time `field:"ENDDATETIME"`
	Comments                    string    `field:"COMMENTS"`
	TechIdentifyingSpeciesInLab string    `field:"IDBYTECH"`
	TechSortingTrapResultsInLab string    `field:"SORTBYTECH"`
	Processed                   int16     `field:"PROCESSED"`
	SiteConditions              string    `field:"SITECOND"`
	LocationName                string    `field:"LOCATIONNAME"`
	RecordStatus                int16     `field:"RECORDSTATUS"`
	Reviewed                    int16     `field:"REVIEWED"`
	ReviewedBy                  string    `field:"REVIEWEDBY"`
	ReviewedDate                time.Time `field:"REVIEWEDDATE"`
	TrapCondition               string    `field:"TRAPCONDITION"`
	TrapNights                  int16     `field:"TRAPNIGHTS"`
	Zone                        string    `field:"ZONE"`
	Zone2                       string    `field:"ZONE2"`
	GlobalID                    uuid.UUID `field:"GlobalID"`
	CreatedUser                 string    `field:"created_user"`
	CreatedDate                 time.Time `field:"created_date"`
	LastEditedUser              string    `field:"last_edited_user"`
	LastEditedDate              time.Time `field:"last_edited_date"`
	SrID                        uuid.UUID `field:"SRID"`
	FieldTech                   string    `field:"FIELDTECH"`
	GatewaySync                 int16     `field:"GATEWAYSYNC"`
	LocID                       uuid.UUID `field:"LOC_ID"`
	Voltage                     float64   `field:"VOLTAGE"`
	Winddir                     string    `field:"WINDDIR"`
	Windspeed                   float64   `field:"WINDSPEED"`
	Avetemp                     float64   `field:"AVETEMP"`
	Raingauge                   float64   `field:"RAINGAUGE"`
	LandingRate                 int16     `field:"LR"`
	Field                       int32     `field:"Field"`
	VectorsurvtrapdataID        string    `field:"VECTORSURVTRAPDATAID"`
	VectorsurvtraplocationID    string    `field:"VECTORSURVTRAPLOCATIONID"`
	CreationDate                time.Time `field:"CreationDate"`
	Creator                     string    `field:"Creator"`
	EditDate                    time.Time `field:"EditDate"`
	Editor                      string    `field:"Editor"`
	Lure                        string    `field:"Lure"`
	Geometry                    json.RawMessage
}

func (x *TrapData) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *TrapData) SetGeometry(m json.RawMessage) { x.Geometry = m }
