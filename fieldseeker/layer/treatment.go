package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type Treatment struct {
	ObjectID         uint      `field:"OBJECTID"`
	InspID           uuid.UUID `field:"QAINSP_ID"`
	Product          string    `field:"PRODUCT"`
	QuantityFloating int16     `field:"QTYFLOATING"`
	QuantitySunk     int16     `field:"QTYSUNK"`
	QuantityDry      int16     `field:"QTYDRY"`
	GroundVisible    int16     `field:"PERCENTGROUNDVISIBLE"`
	Larvae           int16     `field:"LARVAEPRESENT"`
	PercentCoverage  int16     `field:"PERCENTCOVERAGE"`
	AcresTreated     float64   `field:"ACRESTREATED"`
	AcresNotTreated  float64   `field:"ACRESNOTTREATED"`
	GlobalID         uuid.UUID `field:"GlobalID"`
	Comments         string    `field:"COMMENTS"`
	CreatedUser      string    `field:"created_user"`
	CreatedDate      time.Time `field:"created_date"`
	LastEditedUser   string    `field:"last_edited_user"`
	LastEditedDate   time.Time `field:"last_edited_date"`
	CreationDate     time.Time `field:"CreationDate"`
	Creator          string    `field:"Creator"`
	EditDate         time.Time `field:"EditDate"`
	Editor           string    `field:"Editor"`
	Geometry         response.Geometry
}

func (x *Treatment) GetGeometry() response.Geometry  { return x.Geometry }
func (x *Treatment) SetGeometry(m response.Geometry) { x.Geometry = m }
