package layer

import (
	"time"

	"github.com/google/uuid"
)

type InspectionSampleDetailMosquitoDominantStageType string

const (
	InspectionSampleDetailMosquitoDominantStageOne       InspectionSampleDetailMosquitoDominantStageType = "1"
	InspectionSampleDetailMosquitoDominantStageTwo       InspectionSampleDetailMosquitoDominantStageType = "2"
	InspectionSampleDetailMosquitoDominantStageThree     InspectionSampleDetailMosquitoDominantStageType = "3"
	InspectionSampleDetailMosquitoDominantStageFour      InspectionSampleDetailMosquitoDominantStageType = "4"
	InspectionSampleDetailMosquitoDominantStageOneTwo    InspectionSampleDetailMosquitoDominantStageType = "1-2"
	InspectionSampleDetailMosquitoDominantStageThreeFour InspectionSampleDetailMosquitoDominantStageType = "3-4"
)

type InspectionSampleDetailMosquitoAdultActivityType string

const (
	InspectionSampleDetailMosquitoAdultActivityNone     InspectionSampleDetailMosquitoAdultActivityType = "None"
	InspectionSampleDetailMosquitoAdultActivityLight    InspectionSampleDetailMosquitoAdultActivityType = "Light"
	InspectionSampleDetailMosquitoAdultActivityModerate InspectionSampleDetailMosquitoAdultActivityType = "Moderate"
	InspectionSampleDetailMosquitoAdultActivityIntense  InspectionSampleDetailMosquitoAdultActivityType = "Intense"
)

type InspectionSampleDetailMosquitoFieldSpeciesType string

const (
	InspectionSampleDetailMosquitoFieldSpeciesAedes InspectionSampleDetailMosquitoFieldSpeciesType = "Aedes"
	InspectionSampleDetailMosquitoFieldSpeciesCulex InspectionSampleDetailMosquitoFieldSpeciesType = "Culex"
)

type InspectionSampleDetail struct {
	ObjectID           uint                                            `field:"OBJECTID"`
	InspsampleID       uuid.UUID                                       `field:"INSPSAMPLE_ID"`
	FieldSpecies       InspectionSampleDetailMosquitoFieldSpeciesType  `field:"FIELDSPECIES"`
	FieldLarvaCount    int16                                           `field:"FLARVCOUNT"`
	FieldPupaCount     int16                                           `field:"FPUPCOUNT"`
	FieldEggCount      int16                                           `field:"FEGGCOUNT"`
	FieldLarvalStages  string                                          `field:"FLSTAGES"`
	FieldDominantStage InspectionSampleDetailMosquitoDominantStageType `field:"FDOMSTAGE"`
	FieldAdultActivity InspectionSampleDetailMosquitoAdultActivityType `field:"FADULTACT"`
	LabSpecies         string                                          `field:"LABSPECIES"`
	LabLarvaCount      int16                                           `field:"LLARVCOUNT"`
	LabPupaCount       int16                                           `field:"LPUPCOUNT"`
	LabEggCount        int16                                           `field:"LEGGCOUNT"`
	LabDominantStage   InspectionSampleDetailMosquitoDominantStageType `field:"LDOMSTAGE"`
	Comments           string                                          `field:"COMMENTS"`
	GlobalID           uuid.UUID                                       `field:"GlobalID"`
	CreatedUser        string                                          `field:"created_user"`
	CreatedDate        time.Time                                       `field:"created_date"`
	LastEditedUser     string                                          `field:"last_edited_user"`
	LastEditedDate     time.Time                                       `field:"last_edited_date"`
	Processed          int16                                           `field:"PROCESSED"`
	CreationDate       time.Time                                       `field:"CreationDate"`
	Creator            string                                          `field:"Creator"`
	EditDate           time.Time                                       `field:"EditDate"`
	Editor             string                                          `field:"Editor"`
}
