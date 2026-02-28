package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type InspectionSampleDetail struct {
	ObjectID           uint      `field:"OBJECTID"`
	LocID              uuid.UUID `field:"LOC_ID"`
	Start              time.Time `field:"STARTDATETIME"`
	Finish             time.Time `field:"ENDDATETIME"`
	Conditions         string    `field:"SITECOND"`
	Activity           string    `field:"ACTIVITY"`
	Breeding           string    `field:"BREEDING"`
	AdultActivity      string    `field:"ADULTACT"`
	LandingRate        int16     `field:"LANDINGRATE"`
	Seconds            int16     `field:"LANDINGTIME"`
	DominantSpecies    string    `field:"DOMSPECIES"`
	Comments           string    `field:"COMMENTS"`
	AverageTemperature float64   `field:"AVETEMP"`
	WindSpeed          float64   `field:"WINDSPEED"`
	RainGauge          float64   `field:"RAINGAUGE"`
	WindDirection      string    `field:"WINDDIR"`
	Reviewed           int16     `field:"REVIEWED"`
	ReviewedBy         string    `field:"REVIEWEDBY"`
	ReviewedDate       time.Time `field:"REVIEWEDDATE"`
	LocationName       string    `field:"LOCATIONNAME"`
	Zone               string    `field:"ZONE"`
	RecordStatus       int16     `field:"RECORDSTATUS"`
	Zone2              string    `field:"ZONE2"`
	GlobalID           uuid.UUID `field:"GlobalID"`
	CreatedUser        string    `field:"created_user"`
	CreatedDate        time.Time `field:"created_date"`
	LastEditedUser     string    `field:"last_edited_user"`
	LastEditedDate     time.Time `field:"last_edited_date"`
	SrID               uuid.UUID `field:"SRID"`
	FieldTech          string    `field:"FIELDTECH"`
	CreationDate       time.Time `field:"CreationDate"`
	Creator            string    `field:"Creator"`
	EditDate           time.Time `field:"EditDate"`
	Editor             string    `field:"Editor"`
	Geometry           response.Geometry
}

func (x *InspectionSampleDetail) GetGeometry() response.Geometry  { return x.Geometry }
func (x *InspectionSampleDetail) SetGeometry(m response.Geometry) { x.Geometry = m }
