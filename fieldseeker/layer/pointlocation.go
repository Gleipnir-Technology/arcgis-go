package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PointLocation struct {
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
	NextScheduledAction         time.Time `field:"NEXTACTIONDATESCHEDULED"`
	LarvalInspectionInterval    int16     `field:"LARVINSPECTINTERVAL"`
	Zone2                       string    `field:"ZONE2"`
	Locationnumber              int32     `field:"LOCATIONNUMBER"`
	GlobalID                    uuid.UUID `field:"GlobalID"`
	SourceType                  string    `field:"STYPE"`
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
	LastInspectionActivity      string    `field:"LASTINSPECTACTIVITY"`
	LastTreatmentActivity       string    `field:"LASTTREATACTIVITY"`
	LastInspectionConditions    string    `field:"LASTINSPECTCONDITIONS"`
	WaterOrigin                 string    `field:"WATERORIGIN"`
	X                           float64   `field:"X"`
	Y                           float64   `field:"Y"`
	AssignedTech                string    `field:"ASSIGNEDTECH"`
	CreationDate                time.Time `field:"CreationDate"`
	Creator                     string    `field:"Creator"`
	EditDate                    time.Time `field:"EditDate"`
	Editor                      string    `field:"Editor"`
	Jurisdiction                string    `field:"JURISDICTION"`
	ReasonForDeactivation       string    `field:"deactivate_reason"`
	ScalarPriority              int32     `field:"scalarPriority"`
	SourceStatus                string    `field:"sourceStatus"`
	Geometry                    json.RawMessage
}

func (x *PointLocation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *PointLocation) SetGeometry(m json.RawMessage) { x.Geometry = m }
