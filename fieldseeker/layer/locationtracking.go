package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type LocationTracking struct {
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
	Geometry       json.RawMessage
}

func (x *LocationTracking) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *LocationTracking) SetGeometry(m json.RawMessage) { x.Geometry = m }
