package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type QAMosquitoInspection struct {
	ObjectID                  uint      `field:"OBJECTID"`
	PositiveDips              int16     `field:"POSDIPS"`
	Action                    string    `field:"ACTIONTAKEN"`
	Comments                  string    `field:"COMMENTS"`
	AverageTemperature        float64   `field:"AVETEMP"`
	WindSpeed                 float64   `field:"WINDSPEED"`
	RainGauge                 float64   `field:"RAINGAUGE"`
	GlobalID                  uuid.UUID `field:"GlobalID"`
	Start                     time.Time `field:"STARTDATETIME"`
	Finish                    time.Time `field:"ENDDATETIME"`
	WindDirection             string    `field:"WINDDIR"`
	Reviewed                  int16     `field:"REVIEWED"`
	Reviewedby                string    `field:"REVIEWEDBY"`
	Revieweddate              time.Time `field:"REVIEWEDDATE"`
	Locationname              string    `field:"LOCATIONNAME"`
	Zone                      string    `field:"ZONE"`
	Recordstatus              int16     `field:"RECORDSTATUS"`
	Zone2                     string    `field:"ZONE2"`
	LandingRate               int16     `field:"LR"`
	NegativeDips              int16     `field:"NEGDIPS"`
	TotalAcres                float64   `field:"TOTALACRES"`
	AcresBreeding             float64   `field:"ACRESBREEDING"`
	FishPresent               int16     `field:"FISH"`
	SiteType                  string    `field:"SITETYPE"`
	BreedingPotential         string    `field:"BREEDINGPOTENTIAL"`
	MovingWater               int16     `field:"MOVINGWATER"`
	NoEvidenceOfWaterEver     int16     `field:"NOWATEREVER"`
	MosquitoHabitatIndicators string    `field:"MOSQUITOHABITAT"`
	HabitatValue              int16     `field:"HABVALUE1"`
	Habvalue1percent          int16     `field:"HABVALUE1PERCENT"`
	HabitatValue2             int16     `field:"HABVALUE2"`
	Habvalue2percent          int16     `field:"HABVALUE2PERCENT"`
	Potential                 int16     `field:"POTENTIAL"`
	LarvaePresent             int16     `field:"LARVAEPRESENT"`
	LarvaeInsideTreatedArea   int16     `field:"LARVAEINSIDETREATEDAREA"`
	LarvaeOutsideTreatedArea  int16     `field:"LARVAEOUTSIDETREATEDAREA"`
	ReasonLarvaePresent       string    `field:"LARVAEREASON"`
	AquaticOrganisms          string    `field:"AQUATICORGANISMS"`
	Vegetation                string    `field:"VEGETATION"`
	SourceReduction           string    `field:"SOURCEREDUCTION"`
	WaterPresent              int16     `field:"WATERPRESENT"`
	WaterMovement             string    `field:"WATERMOVEMENT1"`
	Watermovement1percent     int16     `field:"WATERMOVEMENT1PERCENT"`
	WaterMovement2            string    `field:"WATERMOVEMENT2"`
	Watermovement2percent     int16     `field:"WATERMOVEMENT2PERCENT"`
	SoilConditions            string    `field:"SOILCONDITIONS"`
	HowLongWaterPresent       string    `field:"WATERDURATION"`
	WaterSource               string    `field:"WATERSOURCE"`
	WaterConditions           string    `field:"WATERCONDITIONS"`
	AdultActivity             int16     `field:"ADULTACTIVITY"`
	LinelocID                 uuid.UUID `field:"LINELOCID"`
	PointlocID                uuid.UUID `field:"POINTLOCID"`
	PolygonlocID              uuid.UUID `field:"POLYGONLOCID"`
	CreatedUser               string    `field:"created_user"`
	CreatedDate               time.Time `field:"created_date"`
	LastEditedUser            string    `field:"last_edited_user"`
	LastEditedDate            time.Time `field:"last_edited_date"`
	FieldTech                 string    `field:"FIELDTECH"`
	CreationDate              time.Time `field:"CreationDate"`
	Creator                   string    `field:"Creator"`
	EditDate                  time.Time `field:"EditDate"`
	Editor                    string    `field:"Editor"`
	Geometry                  json.RawMessage
}

func (x *QAMosquitoInspection) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *QAMosquitoInspection) SetGeometry(m json.RawMessage) { x.Geometry = m }
