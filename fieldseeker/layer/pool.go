package layer

import (
	"time"

	"github.com/google/uuid"
)

type PoolPoolTESTMETHOD670efbfba86d41ba8e2d3cab5d749e7fType string

const (
	PoolPoolTESTMETHOD670efbfba86d41ba8e2d3cab5d749e7fRTPCR PoolPoolTESTMETHOD670efbfba86d41ba8e2d3cab5d749e7fType = "RT-PCR"
)

type PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type string

const (
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97WNV   PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "WNV"
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97SLEV  PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "SLEV"
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97WEEV  PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "WEEV"
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97DENV  PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "DENV"
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97ZIKV  PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "ZIKV"
	PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97CHIKV PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type = "CHIKV"
)

type PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType string

const (
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faWNV     PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "WNV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faSLEV    PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "SLEV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faWEEV    PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "WEEV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faDENV    PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "DENV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faZIKV    PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "ZIKV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faCHIKV   PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "CHIKV"
	PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faWNVSLEV PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType = "WNV/SLEV"
)

type PoolMosquitoLabNameType string

const (
	PoolMosquitoLabNameInternalLab PoolMosquitoLabNameType = "Internal Lab"
	PoolMosquitoLabNameStateLab    PoolMosquitoLabNameType = "State Lab"
)

type PoolNotInUITFType int16

const (
	PoolNotInUITFTrue  PoolNotInUITFType = 1
	PoolNotInUITFFalse PoolNotInUITFType = 0
)

type Pool struct {
	ObjectID               uint                                                      `field:"OBJECTID"`
	TrapDataID             uuid.UUID                                                 `field:"TRAPDATA_ID"`
	DateSent               time.Time                                                 `field:"DATESENT"`
	SurveyTech             string                                                    `field:"SURVTECH"`
	DateTested             time.Time                                                 `field:"DATETESTED"`
	TestTech               string                                                    `field:"TESTTECH"`
	Comments               string                                                    `field:"COMMENTS"`
	SampleID               string                                                    `field:"SAMPLEID"`
	Processed              PoolNotInUITFType                                         `field:"PROCESSED"`
	LabID                  uuid.UUID                                                 `field:"LAB_ID"`
	TestMethods            PoolPoolTESTMETHOD670efbfba86d41ba8e2d3cab5d749e7fType    `field:"TESTMETHOD"`
	DiseasesTested         PoolPoolDISEASETESTED0f02232949c04c7e8de820b9b515ed97Type `field:"DISEASETESTED"`
	DiseasesPositive       PoolPoolDISEASEPOS6889f8dd00074874aa726907e78497faType    `field:"DISEASEPOS"`
	GlobalID               uuid.UUID                                                 `field:"GlobalID"`
	CreatedUser            string                                                    `field:"created_user"`
	CreatedDate            time.Time                                                 `field:"created_date"`
	LastEditedUser         string                                                    `field:"last_edited_user"`
	LastEditedDate         time.Time                                                 `field:"last_edited_date"`
	Lab                    PoolMosquitoLabNameType                                   `field:"LAB"`
	PoolYear               int16                                                     `field:"POOLYEAR"`
	GatewaySync            int16                                                     `field:"GATEWAYSYNC"`
	VectorsurvcollectionID string                                                    `field:"VECTORSURVCOLLECTIONID"`
	VectorsurvpoolID       string                                                    `field:"VECTORSURVPOOLID"`
	VectorsurvtrapdataID   string                                                    `field:"VECTORSURVTRAPDATAID"`
	CreationDate           time.Time                                                 `field:"CreationDate"`
	Creator                string                                                    `field:"Creator"`
	EditDate               time.Time                                                 `field:"EditDate"`
	Editor                 string                                                    `field:"Editor"`
}
