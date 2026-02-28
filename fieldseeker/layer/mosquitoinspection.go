package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type MosquitoInspection struct {
	ObjectID       uint      `field:"OBJECTID"`
	Activity       string    `field:"ACTIVITY"`
	Start          time.Time `field:"STARTDATETIME"`
	Finish         time.Time `field:"ENDDATETIME"`
	Comments       string    `field:"COMMENTS"`
	ExternalID     string    `field:"EXTERNALID"`
	EquipmentType  string    `field:"EQUIPTYPE"`
	LocationName   string    `field:"LOCATIONNAME"`
	Zone           string    `field:"ZONE"`
	Zone2          string    `field:"ZONE2"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	LinelocID      uuid.UUID `field:"LINELOCID"`
	PointlocID     uuid.UUID `field:"POINTLOCID"`
	PolygonlocID   uuid.UUID `field:"POLYGONLOCID"`
	LclocID        uuid.UUID `field:"LCLOCID"`
	SamplelocID    uuid.UUID `field:"SAMPLELOCID"`
	SrID           uuid.UUID `field:"SRID"`
	TraplocID      uuid.UUID `field:"TRAPLOCID"`
	FieldTech      string    `field:"FIELDTECH"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	RodentlocID    uuid.UUID `field:"RODENTLOCID"`
	Geometry       response.Geometry
}

func (x *MosquitoInspection) GetGeometry() response.Geometry  { return x.Geometry }
func (x *MosquitoInspection) SetGeometry(m response.Geometry) { x.Geometry = m }
