package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type RodentLocation struct {
	ObjectID                     uint      `field:"OBJECTID"`
	LocationName                 string    `field:"LOCATIONNAME"`
	Zone                         string    `field:"ZONE"`
	Zone2                        string    `field:"ZONE2"`
	Habitat                      string    `field:"HABITAT"`
	Priority                     string    `field:"PRIORITY"`
	Usetype                      string    `field:"USETYPE"`
	Active                       int16     `field:"ACTIVE"`
	Description                  string    `field:"DESCRIPTION"`
	Accessdesc                   string    `field:"ACCESSDESC"`
	Comments                     string    `field:"COMMENTS"`
	Symbology                    string    `field:"SYMBOLOGY"`
	ExternalID                   string    `field:"EXTERNALID"`
	Nextactiondatescheduled      time.Time `field:"NEXTACTIONDATESCHEDULED"`
	Locationnumber               int32     `field:"LOCATIONNUMBER"`
	LastInspectionDate           time.Time `field:"LASTINSPECTDATE"`
	LastInspectionSpecies        string    `field:"LASTINSPECTSPECIES"`
	LastInspectionAction         string    `field:"LASTINSPECTACTION"`
	LastInspectionConditions     string    `field:"LASTINSPECTCONDITIONS"`
	LastInspectionRodentEvidence string    `field:"LASTINSPECTRODENTEVIDENCE"`
	GlobalID                     uuid.UUID `field:"GlobalID"`
	CreatedUser                  string    `field:"created_user"`
	CreatedDate                  time.Time `field:"created_date"`
	LastEditedUser               string    `field:"last_edited_user"`
	LastEditedDate               time.Time `field:"last_edited_date"`
	CreationDate                 time.Time `field:"CreationDate"`
	Creator                      string    `field:"Creator"`
	EditDate                     time.Time `field:"EditDate"`
	Editor                       string    `field:"Editor"`
	Jurisdiction                 string    `field:"JURISDICTION"`
	Geometry                     json.RawMessage
}

func (x *RodentLocation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *RodentLocation) SetGeometry(m json.RawMessage) { x.Geometry = m }
