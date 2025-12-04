package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Treatment struct {
	ObjectID             uint      `field:"OBJECTID"`
	Activity             string    `field:"ACTIVITY"`
	AreaTreated          float64   `field:"TREATAREA"`
	AreaUnit             string    `field:"AREAUNIT"`
	Product              string    `field:"PRODUCT"`
	Quantity             float64   `field:"QTY"`
	QuantityUnit         string    `field:"QTYUNIT"`
	Method               string    `field:"METHOD"`
	EquipmentType        string    `field:"EQUIPTYPE"`
	Comments             string    `field:"COMMENTS"`
	AverageTemperature   float64   `field:"AVETEMP"`
	WindSpeed            float64   `field:"WINDSPEED"`
	WindDirection        string    `field:"WINDDIR"`
	RainGauge            float64   `field:"RAINGAUGE"`
	Start                time.Time `field:"STARTDATETIME"`
	Finish               time.Time `field:"ENDDATETIME"`
	InspID               uuid.UUID `field:"INSP_ID"`
	Reviewed             int16     `field:"REVIEWED"`
	ReviewedBy           string    `field:"REVIEWEDBY"`
	ReviewedDate         time.Time `field:"REVIEWEDDATE"`
	LocationName         string    `field:"LOCATIONNAME"`
	Zone                 string    `field:"ZONE"`
	WarningOverride      int16     `field:"WARNINGOVERRIDE"`
	RecordStatus         int16     `field:"RECORDSTATUS"`
	Zone2                string    `field:"ZONE2"`
	TreatedAcres         float64   `field:"TREATACRES"`
	TireCount            int16     `field:"TIRECOUNT"`
	CatchBasinCount      int16     `field:"CBCOUNT"`
	ContainerCount       int16     `field:"CONTAINERCOUNT"`
	GlobalID             uuid.UUID `field:"GlobalID"`
	TreatmentLength      float64   `field:"TREATMENTLENGTH"`
	TreatmentHours       float64   `field:"TREATMENTHOURS"`
	TreatmentLengthUnits string    `field:"TREATMENTLENGTHUNITS"`
	LinelocID            uuid.UUID `field:"LINELOCID"`
	PointlocID           uuid.UUID `field:"POINTLOCID"`
	PolygonlocID         uuid.UUID `field:"POLYGONLOCID"`
	SrID                 uuid.UUID `field:"SRID"`
	SdID                 uuid.UUID `field:"SDID"`
	BarrierrouteID       uuid.UUID `field:"BARRIERROUTEID"`
	UlvrouteID           uuid.UUID `field:"ULVROUTEID"`
	FieldTech            string    `field:"FIELDTECH"`
	PtaID                uuid.UUID `field:"PTAID"`
	Flowrate             float64   `field:"FLOWRATE"`
	Habitat              string    `field:"HABITAT"`
	TreatHectares        float64   `field:"TREATHECTARES"`
	InventoryLocation    string    `field:"INVLOC"`
	TempConditions       string    `field:"temp_SITECOND"`
	Conditions           string    `field:"SITECOND"`
	TotalCostProduct     float64   `field:"TotalCostProdcut"`
	CreationDate         time.Time `field:"CreationDate"`
	Creator              string    `field:"Creator"`
	EditDate             time.Time `field:"EditDate"`
	Editor               string    `field:"Editor"`
	TargetSpecies        string    `field:"TARGETSPECIES"`
	Geometry             json.RawMessage
}

func (x *Treatment) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *Treatment) SetGeometry(m json.RawMessage) { x.Geometry = m }
