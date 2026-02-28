package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type ProposedTreatmentArea struct {
	ObjectID       uint      `field:"OBJECTID"`
	DateTimeStamp  time.Time `field:"DateTimeStamp"`
	Severity       string    `field:"Severity"`
	Application    string    `field:"Application"`
	Method         string    `field:"Method"`
	Message        string    `field:"Message"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       response.Geometry
}

func (x *ProposedTreatmentArea) GetGeometry() response.Geometry  { return x.Geometry }
func (x *ProposedTreatmentArea) SetGeometry(m response.Geometry) { x.Geometry = m }
