package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type MosquitoInspection struct {
	ObjectID               uint      `field:"OBJECTID"`
	Dips                   int16     `field:"NUMDIPS"`
	Activity               string    `field:"ACTIVITY"`
	Breeding               string    `field:"BREEDING"`
	TotalLarvae            int16     `field:"TOTLARVAE"`
	TotalPupae             int16     `field:"TOTPUPAE"`
	Eggs                   int16     `field:"EGGS"`
	PositiveDips           int16     `field:"POSDIPS"`
	AdultActivity          string    `field:"ADULTACT"`
	LarvalStages           string    `field:"LSTAGES"`
	DominantStage          string    `field:"DOMSTAGE"`
	Action                 string    `field:"ACTIONTAKEN"`
	Comments               string    `field:"COMMENTS"`
	AverageTemperature     float64   `field:"AVETEMP"`
	WindSpeed              float64   `field:"WINDSPEED"`
	RainGauge              float64   `field:"RAINGAUGE"`
	Start                  time.Time `field:"STARTDATETIME"`
	Finish                 time.Time `field:"ENDDATETIME"`
	WindDirection          string    `field:"WINDDIR"`
	AverageLarvae          float64   `field:"AVGLARVAE"`
	AveragePupae           float64   `field:"AVGPUPAE"`
	Reviewed               int16     `field:"REVIEWED"`
	ReviewedBy             string    `field:"REVIEWEDBY"`
	ReviewedDate           time.Time `field:"REVIEWEDDATE"`
	LocationName           string    `field:"LOCATIONNAME"`
	Zone                   string    `field:"ZONE"`
	RecordStatus           int16     `field:"RECORDSTATUS"`
	Zone2                  string    `field:"ZONE2"`
	PersonalContact        int16     `field:"PERSONALCONTACT"`
	TireCount              int16     `field:"TIRECOUNT"`
	CatchBasinCount        int16     `field:"CBCOUNT"`
	ContainerCount         int16     `field:"CONTAINERCOUNT"`
	FieldSpecies           string    `field:"FIELDSPECIES"`
	GlobalID               uuid.UUID `field:"GlobalID"`
	CreatedUser            string    `field:"created_user"`
	CreatedDate            time.Time `field:"created_date"`
	LastEditedUser         string    `field:"last_edited_user"`
	LastEditedDate         time.Time `field:"last_edited_date"`
	LinelocID              uuid.UUID `field:"LINELOCID"`
	PointlocID             uuid.UUID `field:"POINTLOCID"`
	PolygonlocID           uuid.UUID `field:"POLYGONLOCID"`
	SrID                   uuid.UUID `field:"SRID"`
	FieldTech              string    `field:"FIELDTECH"`
	LarvaePresent          int16     `field:"LARVAEPRESENT"`
	PupaePresent           int16     `field:"PUPAEPRESENT"`
	StormDrainID           uuid.UUID `field:"SDID"`
	Conditions             string    `field:"SITECOND"`
	PositiveContainerCount int16     `field:"POSITIVECONTAINERCOUNT"`
	CreationDate           time.Time `field:"CreationDate"`
	Creator                string    `field:"Creator"`
	EditDate               time.Time `field:"EditDate"`
	Editor                 string    `field:"Editor"`
	Jurisdiction           string    `field:"JURISDICTION"`
	VisualMonitoring       int16     `field:"VISUALMONITORING"`
	VmComments             string    `field:"VMCOMMENTS"`
	AdminAction            string    `field:"adminAction"`
	PtaID                  uuid.UUID `field:"PTAID"`
	Geometry               json.RawMessage
}

func (x *MosquitoInspection) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *MosquitoInspection) SetGeometry(m json.RawMessage) { x.Geometry = m }
