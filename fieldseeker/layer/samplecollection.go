package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SampleCollection struct {
	ObjectID               uint      `field:"OBJECTID"`
	LocID                  uuid.UUID `field:"LOC_ID"`
	Start                  time.Time `field:"STARTDATETIME"`
	Finish                 time.Time `field:"ENDDATETIME"`
	Conditions             string    `field:"SITECOND"`
	SampleID               string    `field:"SAMPLEID"`
	SurveillanceTechnician string    `field:"SURVTECH"`
	Sent                   time.Time `field:"DATESENT"`
	Tested                 time.Time `field:"DATETESTED"`
	TestTechnician         string    `field:"TESTTECH"`
	Comments               string    `field:"COMMENTS"`
	Processed              int16     `field:"PROCESSED"`
	SampleType             string    `field:"SAMPLETYPE"`
	SampleCondition        string    `field:"SAMPLECOND"`
	Species                string    `field:"SPECIES"`
	Sex                    string    `field:"SEX"`
	AverageTemperature     float64   `field:"AVETEMP"`
	WindSpeed              float64   `field:"WINDSPEED"`
	WindDirection          string    `field:"WINDDIR"`
	RainGauge              float64   `field:"RAINGAUGE"`
	Activity               string    `field:"ACTIVITY"`
	TestMethod             string    `field:"TESTMETHOD"`
	DiseaseTested          string    `field:"DISEASETESTED"`
	DiseasePositive        string    `field:"DISEASEPOS"`
	Reviewed               int16     `field:"REVIEWED"`
	ReviewedBy             string    `field:"REVIEWEDBY"`
	ReviewedDate           time.Time `field:"REVIEWEDDATE"`
	LocationName           string    `field:"LOCATIONNAME"`
	Zone                   string    `field:"ZONE"`
	RecordStatus           int16     `field:"RECORDSTATUS"`
	Zone2                  string    `field:"ZONE2"`
	GlobalID               uuid.UUID `field:"GlobalID"`
	CreatedUser            string    `field:"created_user"`
	CreatedDate            time.Time `field:"created_date"`
	LastEditedUser         string    `field:"last_edited_user"`
	LastEditedDate         time.Time `field:"last_edited_date"`
	Lab                    string    `field:"LAB"`
	FieldTech              string    `field:"FIELDTECH"`
	FlockID                uuid.UUID `field:"FLOCKID"`
	SampleCount            int16     `field:"SAMPLECOUNT"`
	ChickenID              uuid.UUID `field:"CHICKENID"`
	GatewaySync            int16     `field:"GATEWAYSYNC"`
	CreationDate           time.Time `field:"CreationDate"`
	Creator                string    `field:"Creator"`
	EditDate               time.Time `field:"EditDate"`
	Editor                 string    `field:"Editor"`
	Geometry               json.RawMessage
}

func (x *SampleCollection) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *SampleCollection) SetGeometry(m json.RawMessage) { x.Geometry = m }
