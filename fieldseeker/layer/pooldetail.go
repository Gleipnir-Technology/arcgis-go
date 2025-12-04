package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PoolDetail struct {
	ObjectID       uint      `field:"OBJECTID"`
	TrapDataID     uuid.UUID `field:"TRAPDATA_ID"`
	PoolID         uuid.UUID `field:"POOL_ID"`
	Species        string    `field:"SPECIES"`
	Females        int16     `field:"FEMALES"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       json.RawMessage
}

func (x *PoolDetail) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *PoolDetail) SetGeometry(m json.RawMessage) { x.Geometry = m }
