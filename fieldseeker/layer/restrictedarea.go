package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type RestrictedArea struct {
	ObjectID       uint      `field:"OBJECTID"`
	Name           string    `field:"NAME"`
	Requiredtype   string    `field:"REQUIREDTYPE"`
	Value          string    `field:"VALUE"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	DisplayName    string    `field:"DISPLAY"`
	Visible        int16     `field:"VISIBLE"`
	Geometry       response.Geometry
}

func (x *RestrictedArea) GetGeometry() response.Geometry  { return x.Geometry }
func (x *RestrictedArea) SetGeometry(m response.Geometry) { x.Geometry = m }
