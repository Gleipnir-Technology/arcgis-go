package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type AerialSprayLine struct {
	ObjectID                  uint      `field:"OBJECTID"`
	Method                    string    `field:"METHOD"`
	Comments                  string    `field:"COMMENTS"`
	Zone                      string    `field:"ZONE"`
	Reviewed                  int16     `field:"REVIEWED"`
	ReviewedBy                string    `field:"REVIEWEDBY"`
	ReviewedDate              time.Time `field:"REVIEWEDDATE"`
	Zone2                     string    `field:"ZONE2"`
	CompletedDate             time.Time `field:"COMPLETEDDATE"`
	CompletedBy               string    `field:"COMPLETEDBY"`
	Completed                 int16     `field:"COMPLETED"`
	IsSprayRoute              int16     `field:"ISSPRAYROUTE"`
	Name                      string    `field:"NAME"`
	Acres                     float64   `field:"ACRES"`
	GlobalID                  uuid.UUID `field:"GlobalID"`
	Exported                  int16     `field:"EXPORTED"`
	TargetProduct             string    `field:"TARGETPRODUCT"`
	TargetAppRate             float64   `field:"TARGETAPPRATE"`
	Hectares                  float64   `field:"HECTARES"`
	LastTreatmentActivity     string    `field:"LASTTREATACTIVITY"`
	LastTreatmentDate         time.Time `field:"LASTTREATDATE"`
	LastTreatmentProduct      string    `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity     float64   `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit string    `field:"LASTTREATQTYUNIT"`
	Priority                  string    `field:"PRIORITY"`
	DueDate                   time.Time `field:"DUEDATE"`
	CreationDate              time.Time `field:"CreationDate"`
	Creator                   string    `field:"Creator"`
	EditDate                  time.Time `field:"EditDate"`
	Editor                    string    `field:"Editor"`
	TargetSpecies             string    `field:"TARGETSPECIES"`
	ShapeArea                 float64   `field:"Shape__Area"`
	ShapeLength               float64   `field:"Shape__Length"`
	Geometry                  response.Geometry
}

func (x *AerialSprayLine) GetGeometry() response.Geometry  { return x.Geometry }
func (x *AerialSprayLine) SetGeometry(m response.Geometry) { x.Geometry = m }
