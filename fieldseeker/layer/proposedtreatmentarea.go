package layer

import (
	"time"

	"github.com/google/uuid"
)

type ProposedTreatmentAreaMosquitoTreatmentMethodType string

const (
	ProposedTreatmentAreaMosquitoTreatmentMethodArgo                 ProposedTreatmentAreaMosquitoTreatmentMethodType = "Argo"
	ProposedTreatmentAreaMosquitoTreatmentMethodATV                  ProposedTreatmentAreaMosquitoTreatmentMethodType = "ATV"
	ProposedTreatmentAreaMosquitoTreatmentMethodBackpack             ProposedTreatmentAreaMosquitoTreatmentMethodType = "Backpack"
	ProposedTreatmentAreaMosquitoTreatmentMethodDrone                ProposedTreatmentAreaMosquitoTreatmentMethodType = "Drone"
	ProposedTreatmentAreaMosquitoTreatmentMethodManual               ProposedTreatmentAreaMosquitoTreatmentMethodType = "Manual"
	ProposedTreatmentAreaMosquitoTreatmentMethodTruck                ProposedTreatmentAreaMosquitoTreatmentMethodType = "Truck"
	ProposedTreatmentAreaMosquitoTreatmentMethodULV                  ProposedTreatmentAreaMosquitoTreatmentMethodType = "ULV"
	ProposedTreatmentAreaMosquitoTreatmentMethodEnhancedSurveillance ProposedTreatmentAreaMosquitoTreatmentMethodType = "Enhanced_Surveillance"
)

type ProposedTreatmentAreaNotInUITFType int16

const (
	ProposedTreatmentAreaNotInUITFTrue  ProposedTreatmentAreaNotInUITFType = 1
	ProposedTreatmentAreaNotInUITFFalse ProposedTreatmentAreaNotInUITFType = 0
)

type ProposedTreatmentAreaLocationPriorityType string

const (
	ProposedTreatmentAreaLocationPriorityLow    ProposedTreatmentAreaLocationPriorityType = "Low"
	ProposedTreatmentAreaLocationPriorityMedium ProposedTreatmentAreaLocationPriorityType = "Medium"
	ProposedTreatmentAreaLocationPriorityHigh   ProposedTreatmentAreaLocationPriorityType = "High"
	ProposedTreatmentAreaLocationPriorityNone   ProposedTreatmentAreaLocationPriorityType = "None"
)

type ProposedTreatmentArea struct {
	Objectid                  uint                                             `field:"OBJECTID"`
	Method                    ProposedTreatmentAreaMosquitoTreatmentMethodType `field:"METHOD"`
	Comments                  string                                           `field:"COMMENTS"`
	Zone                      string                                           `field:"ZONE"`
	Reviewed                  ProposedTreatmentAreaNotInUITFType               `field:"REVIEWED"`
	ReviewedBy                string                                           `field:"REVIEWEDBY"`
	ReviewedDate              time.Time                                        `field:"REVIEWEDDATE"`
	Zone2                     string                                           `field:"ZONE2"`
	CompletedDate             time.Time                                        `field:"COMPLETEDDATE"`
	CompletedBy               string                                           `field:"COMPLETEDBY"`
	Completed                 ProposedTreatmentAreaNotInUITFType               `field:"COMPLETED"`
	IsSprayRoute              ProposedTreatmentAreaNotInUITFType               `field:"ISSPRAYROUTE"`
	Name                      string                                           `field:"NAME"`
	Acres                     float64                                          `field:"ACRES"`
	GlobalID                  uuid.UUID                                        `field:"GlobalID"`
	Exported                  ProposedTreatmentAreaNotInUITFType               `field:"EXPORTED"`
	TargetProduct             string                                           `field:"TARGETPRODUCT"`
	TargetAppRate             float64                                          `field:"TARGETAPPRATE"`
	Hectares                  float64                                          `field:"HECTARES"`
	LastTreatmentActivity     string                                           `field:"LASTTREATACTIVITY"`
	LastTreatmentDate         time.Time                                        `field:"LASTTREATDATE"`
	LastTreatmentProduct      string                                           `field:"LASTTREATPRODUCT"`
	LastTreatmentQuantity     float64                                          `field:"LASTTREATQTY"`
	LastTreatmentQuantityUnit string                                           `field:"LASTTREATQTYUNIT"`
	Priority                  ProposedTreatmentAreaLocationPriorityType        `field:"PRIORITY"`
	DueDate                   time.Time                                        `field:"DUEDATE"`
	CreationDate              time.Time                                        `field:"CreationDate"`
	Creator                   string                                           `field:"Creator"`
	EditDate                  time.Time                                        `field:"EditDate"`
	Editor                    string                                           `field:"Editor"`
	TargetSpecies             string                                           `field:"TARGETSPECIES"`
	ShapeArea                 float64                                          `field:"Shape__Area"`
	ShapeLength               float64                                          `field:"Shape__Length"`
}
