package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type Zones struct {
	ObjectID       uint      `field:"OBJECTID"`
	Name           string    `field:"NAME"`
	AlternateName  string    `field:"ALTNAME"`
	PoolUse        int16     `field:"POOLUSE"`
	Active         int16     `field:"ACTIVE"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Genus          string    `field:"GENUS"`
	MappingValue   string    `field:"MAPPINGVALUE"`
	Species        string    `field:"SPECIES"`
	Priority       string    `field:"PRIORITY"`
	Geometry       response.Geometry
}

func (x *Zones) GetGeometry() response.Geometry  { return x.Geometry }
func (x *Zones) SetGeometry(m response.Geometry) { x.Geometry = m }
