package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Pool struct {
	ObjectID               uint      `field:"OBJECTID"`
	TrapDataID             uuid.UUID `field:"TRAPDATA_ID"`
	DateSent               time.Time `field:"DATESENT"`
	SurveyTech             string    `field:"SURVTECH"`
	DateTested             time.Time `field:"DATETESTED"`
	TestTech               string    `field:"TESTTECH"`
	Comments               string    `field:"COMMENTS"`
	SampleID               string    `field:"SAMPLEID"`
	Processed              int16     `field:"PROCESSED"`
	LabID                  uuid.UUID `field:"LAB_ID"`
	TestMethods            string    `field:"TESTMETHOD"`
	DiseasesTested         string    `field:"DISEASETESTED"`
	DiseasesPositive       string    `field:"DISEASEPOS"`
	GlobalID               uuid.UUID `field:"GlobalID"`
	CreatedUser            string    `field:"created_user"`
	CreatedDate            time.Time `field:"created_date"`
	LastEditedUser         string    `field:"last_edited_user"`
	LastEditedDate         time.Time `field:"last_edited_date"`
	Lab                    string    `field:"LAB"`
	PoolYear               int16     `field:"POOLYEAR"`
	GatewaySync            int16     `field:"GATEWAYSYNC"`
	VectorsurvcollectionID string    `field:"VECTORSURVCOLLECTIONID"`
	VectorsurvpoolID       string    `field:"VECTORSURVPOOLID"`
	VectorsurvtrapdataID   string    `field:"VECTORSURVTRAPDATAID"`
	CreationDate           time.Time `field:"CreationDate"`
	Creator                string    `field:"Creator"`
	EditDate               time.Time `field:"EditDate"`
	Editor                 string    `field:"Editor"`
	Geometry               json.RawMessage
}

func (x *Pool) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *Pool) SetGeometry(m json.RawMessage) { x.Geometry = m }
