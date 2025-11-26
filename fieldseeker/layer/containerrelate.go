package layer

import (
	"time"

	"github.com/google/uuid"
)

type ContainerRelateMosquitoContainerTypeType string

const (
	ContainerRelateMosquitoContainerTypeAquarium         ContainerRelateMosquitoContainerTypeType = "Aquarium"
	ContainerRelateMosquitoContainerTypeFlowerpot        ContainerRelateMosquitoContainerTypeType = "Flower pot"
	ContainerRelateMosquitoContainerTypeFivegallonbucket ContainerRelateMosquitoContainerTypeType = "5 gallon bucket"
	ContainerRelateMosquitoContainerTypeFountain         ContainerRelateMosquitoContainerTypeType = "Fountain"
	ContainerRelateMosquitoContainerTypeBirdbath         ContainerRelateMosquitoContainerTypeType = "Bird bath"
)

type ContainerRelate struct {
	ObjectID       uint                                     `field:"OBJECTID"`
	GlobalID       uuid.UUID                                `field:"GlobalID"`
	CreatedUser    string                                   `field:"created_user"`
	CreatedDate    time.Time                                `field:"created_date"`
	LastEditedUser string                                   `field:"last_edited_user"`
	LastEditedDate time.Time                                `field:"last_edited_date"`
	InspsampleID   uuid.UUID                                `field:"INSPSAMPLEID"`
	MosquitoinspID uuid.UUID                                `field:"MOSQUITOINSPID"`
	TreatmentID    uuid.UUID                                `field:"TREATMENTID"`
	ContainerType  ContainerRelateMosquitoContainerTypeType `field:"CONTAINERTYPE"`
	CreationDate   time.Time                                `field:"CreationDate"`
	Creator        string                                   `field:"Creator"`
	EditDate       time.Time                                `field:"EditDate"`
	Editor         string                                   `field:"Editor"`
}
