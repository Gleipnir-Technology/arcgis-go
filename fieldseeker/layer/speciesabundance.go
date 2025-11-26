package layer

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesAbundanceNotInUITFType int16

const (
	SpeciesAbundanceNotInUITFTrue  SpeciesAbundanceNotInUITFType = 1
	SpeciesAbundanceNotInUITFFalse SpeciesAbundanceNotInUITFType = 0
)

type SpeciesAbundance struct {
	ObjectID        uint                          `field:"OBJECTID"`
	TrapDataID      uuid.UUID                     `field:"TRAPDATA_ID"`
	Species         string                        `field:"SPECIES"`
	Males           int16                         `field:"MALES"`
	Unknown         int16                         `field:"UNKNOWN"`
	BloodedFemales  int16                         `field:"BLOODEDFEM"`
	GravidFemales   int16                         `field:"GRAVIDFEM"`
	Larvae          int16                         `field:"LARVAE"`
	PoolsToGenerate int16                         `field:"POOLSTOGEN"`
	Processed       SpeciesAbundanceNotInUITFType `field:"PROCESSED"`
	GlobalID        uuid.UUID                     `field:"GlobalID"`
	CreatedUser     string                        `field:"created_user"`
	CreatedDate     time.Time                     `field:"created_date"`
	LastEditedUser  string                        `field:"last_edited_user"`
	LastEditedDate  time.Time                     `field:"last_edited_date"`
	Pupae           int16                         `field:"PUPAE"`
	Eggs            int16                         `field:"EGGS"`
	Females         int32                         `field:"FEMALES"`
	TotalAdults     int32                         `field:"TOTAL"`
	CreationDate    time.Time                     `field:"CreationDate"`
	Creator         string                        `field:"Creator"`
	EditDate        time.Time                     `field:"EditDate"`
	Editor          string                        `field:"Editor"`
	YearWeek        int32                         `field:"yearWeek"`
	GlobalZScore    float64                       `field:"globalZScore"`
	R7Score         float64                       `field:"r7Score"`
	R8Score         float64                       `field:"r8Score"`
	H3r7            string                        `field:"h3r7"`
	H3r8            string                        `field:"h3r8"`
}
