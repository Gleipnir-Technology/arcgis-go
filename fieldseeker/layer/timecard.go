package layer

import (
	"time"

	"github.com/google/uuid"
)

type TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type string

const (
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Routineinspection TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "Routine inspection"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Pretreatment      TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "Pre-treatment"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Maintenance       TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "Maintenance"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366ULV               TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "ULV"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366BARRIER           TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "BARRIER"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366LOGIN             TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "LOGIN"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366TREATSD           TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "TREATSD"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366SD                TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "SD"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366SITEVISIT         TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "SITEVISIT"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366ONLINE            TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "ONLINE"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366SYNC              TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "SYNC"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366CREATESR          TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "CREATESR"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366LC                TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "LC"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366ACCEPTSR          TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "ACCEPTSR"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366POINT             TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "POINT"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366DOWNLOAD          TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "DOWNLOAD"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366COMPLETESR        TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "COMPLETESR"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366POLYGON           TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "POLYGON"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366TRAP              TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "TRAP"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366SAMPLE            TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "SAMPLE"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366QA                TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "QA"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366PTA               TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "PTA"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366FIELDSCOUTING     TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "FIELDSCOUTING"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366OFFLINE           TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "OFFLINE"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366LINE              TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "LINE"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366TRAPLOCATION      TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "TRAPLOCATION"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366SAMPLELOCATION    TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "SAMPLELOCATION"
	TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366LCLOCATION        TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type = "LCLOCATION"
)

type TimeCardTimeCardEquipmentTypeType string

const (
	TimeCardTimeCardEquipmentTypeSpreader TimeCardTimeCardEquipmentTypeType = "Spreader"
	TimeCardTimeCardEquipmentTypeATV      TimeCardTimeCardEquipmentTypeType = "ATV"
	TimeCardTimeCardEquipmentTypeTruck    TimeCardTimeCardEquipmentTypeType = "Truck"
)

type TimeCard struct {
	ObjectID       uint                                                         `field:"OBJECTID"`
	Activity       TimeCardTimeCardACTIVITY451e67260c084304a35457170dc13366Type `field:"ACTIVITY"`
	Start          time.Time                                                    `field:"STARTDATETIME"`
	Finish         time.Time                                                    `field:"ENDDATETIME"`
	Comments       string                                                       `field:"COMMENTS"`
	ExternalID     string                                                       `field:"EXTERNALID"`
	EquipmentType  TimeCardTimeCardEquipmentTypeType                            `field:"EQUIPTYPE"`
	LocationName   string                                                       `field:"LOCATIONNAME"`
	Zone           string                                                       `field:"ZONE"`
	Zone2          string                                                       `field:"ZONE2"`
	GlobalID       uuid.UUID                                                    `field:"GlobalID"`
	CreatedUser    string                                                       `field:"created_user"`
	CreatedDate    time.Time                                                    `field:"created_date"`
	LastEditedUser string                                                       `field:"last_edited_user"`
	LastEditedDate time.Time                                                    `field:"last_edited_date"`
	LinelocID      uuid.UUID                                                    `field:"LINELOCID"`
	PointlocID     uuid.UUID                                                    `field:"POINTLOCID"`
	PolygonlocID   uuid.UUID                                                    `field:"POLYGONLOCID"`
	LclocID        uuid.UUID                                                    `field:"LCLOCID"`
	SamplelocID    uuid.UUID                                                    `field:"SAMPLELOCID"`
	SrID           uuid.UUID                                                    `field:"SRID"`
	TraplocID      uuid.UUID                                                    `field:"TRAPLOCID"`
	FieldTech      string                                                       `field:"FIELDTECH"`
	CreationDate   time.Time                                                    `field:"CreationDate"`
	Creator        string                                                       `field:"Creator"`
	EditDate       time.Time                                                    `field:"EditDate"`
	Editor         string                                                       `field:"Editor"`
	RodentlocID    uuid.UUID                                                    `field:"RODENTLOCID"`
}
