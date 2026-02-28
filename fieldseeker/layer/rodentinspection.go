package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type RodentInspection struct {
	ObjectID       uint      `field:"OBJECTID"`
	Accuracym      float64   `field:"Accuracy"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	FieldTech      string    `field:"FIELDTECH"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       response.Geometry
}

func (x *RodentInspection) GetGeometry() response.Geometry  { return x.Geometry }
func (x *RodentInspection) SetGeometry(m response.Geometry) { x.Geometry = m }
