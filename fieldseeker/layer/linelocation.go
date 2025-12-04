package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type LineLocation struct {
	ObjectID                    uint      `field:"OBJECTID"`
	Name                        string    `field:"NAME"`
	Zone                        string    `field:"ZONE"`
	Habitat                     string    `field:"HABITAT"`
	Priority                    string    `field:"PRIORITY"`
	UseType                     string    `field:"USETYPE"`
	Active                      int16     `field:"ACTIVE"`
	Description                 string    `field:"DESCRIPTION"`
	AccessDescription           string    `field:"ACCESSDESC"`
	Comments                    string    `field:"COMMENTS"`
	Symbology                   string    `field:"SYMBOLOGY"`
	ExternalID                  string    `field:"EXTERNALID"`
	Acres                       float64   `field:"ACRES"`
	NextScheduledAction         time.Time `field:"NEXTACTIONDATESCHEDULED"`
	LarvalInspectionInterval    int16     `field:"LARVINSPECTINTERVAL"`
	Length                      float64   `field:"LENGTH_FT"`
	Width                       float64   `field:"WIDTH_FT"`
	Zone2                       string    `field:"ZONE2"`
	Locationnumber              int32     `field:"LOCATIONNUMBER"`
	GlobalID                    uuid.UUID `field:"GlobalID"`
	CreatedUser                 string    `field:"created_user"`
	CreatedDate                 time.Time `field:"created_date"`
	LastEditedUser              string    `field:"last_edited_user"`
	LastEditedDate              time.Time `field:"last_edited_date"`
	LastInspectionDate          time.Time `field:"LASTINSPECTDATE"`
	LastInspectionBreeding      string    `field:"LASTINSPECTBREEDING"`
	LastInspectionAverageLarvae float64   `field:"LASTINSPECTAVGLARVAE"`
	LastInspectionAveragePupae  float64   `field:"LASTINSPECTAVGPUPAE"`
	LastInspectionLarvalStages  string    `field:"LASTINSPECTLSTAGES"`
	LastInspectionAction        string    `field:"LASTINSPECTACTIONTAKEN"`
	LastInspectionFieldSpecies  string    `field:"LASTINSPECTFIELDSPECIES"`
	LastTreatmentDate           time.Time `field:"LASTTREATDATE"`
	LastTreatmentProduct        string    `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity       float64   `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit   string    `field:"LASTTREATQTYUNIT"`
	Hectares                    float64   `field:"HECTARES"`
	LastInspectionActivity      string    `field:"LASTINSPECTACTIVITY"`
	LastTreatmentActivity       string    `field:"LASTTREATACTIVITY"`
	LengthMeters                float64   `field:"LENGTH_METERS"`
	WidthMeters                 float64   `field:"WIDTH_METERS"`
	LastInspectionConditions    string    `field:"LASTINSPECTCONDITIONS"`
	WaterOrigin                 string    `field:"WATERORIGIN"`
	CreationDate                time.Time `field:"CreationDate"`
	Creator                     string    `field:"Creator"`
	EditDate                    time.Time `field:"EditDate"`
	Editor                      string    `field:"Editor"`
	Jurisdiction                string    `field:"JURISDICTION"`
	ShapeLength                 float64   `field:"Shape__Length"`
	Geometry                    json.RawMessage
}

func (x *LineLocation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *LineLocation) SetGeometry(m json.RawMessage) { x.Geometry = m }
