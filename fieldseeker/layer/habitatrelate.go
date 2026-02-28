package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type HabitatRelate struct {
	ObjectID        uint      `field:"OBJECTID"`
	TrapDataID      uuid.UUID `field:"TRAPDATA_ID"`
	Species         string    `field:"SPECIES"`
	Males           int16     `field:"MALES"`
	Unknown         int16     `field:"UNKNOWN"`
	BloodedFemales  int16     `field:"BLOODEDFEM"`
	GravidFemales   int16     `field:"GRAVIDFEM"`
	Larvae          int16     `field:"LARVAE"`
	PoolsToGenerate int16     `field:"POOLSTOGEN"`
	Processed       int16     `field:"PROCESSED"`
	GlobalID        uuid.UUID `field:"GlobalID"`
	CreatedUser     string    `field:"created_user"`
	CreatedDate     time.Time `field:"created_date"`
	LastEditedUser  string    `field:"last_edited_user"`
	LastEditedDate  time.Time `field:"last_edited_date"`
	Pupae           int16     `field:"PUPAE"`
	Eggs            int16     `field:"EGGS"`
	Females         int32     `field:"FEMALES"`
	TotalAdults     int32     `field:"TOTAL"`
	CreationDate    time.Time `field:"CreationDate"`
	Creator         string    `field:"Creator"`
	EditDate        time.Time `field:"EditDate"`
	Editor          string    `field:"Editor"`
	YearWeek        int32     `field:"yearWeek"`
	Geometry        response.Geometry
}

func (x *HabitatRelate) GetGeometry() response.Geometry  { return x.Geometry }
func (x *HabitatRelate) SetGeometry(m response.Geometry) { x.Geometry = m }
