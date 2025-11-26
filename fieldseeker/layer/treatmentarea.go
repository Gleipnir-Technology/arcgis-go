package layer

import (
	"time"

	"github.com/google/uuid"
)

type TreatmentArea struct {
	ObjectID       uint      `field:"OBJECTID"`
	TreatID        uuid.UUID `field:"TREAT_ID"`
	SessionID      uuid.UUID `field:"SESSION_ID"`
	TreatmentDate  time.Time `field:"TREATDATE"`
	Comments       string    `field:"COMMENTS"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	Notified       int16     `field:"Notified"`
	Type           string    `field:"Type"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	ShapeArea      float64   `field:"Shape__Area"`
	ShapeLength    float64   `field:"Shape__Length"`
}
