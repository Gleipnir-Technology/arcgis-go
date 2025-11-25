package layer

import (
	"time"

	"github.com/google/uuid"
)

type QAMosquitoInspectionQAWaterSourceType string

const (
	QAMosquitoInspectionQAWaterSourceIrrigation         QAMosquitoInspectionQAWaterSourceType = "Irrigation"
	QAMosquitoInspectionQAWaterSourceManuallyControlled QAMosquitoInspectionQAWaterSourceType = "Manually Controlled"
	QAMosquitoInspectionQAWaterSourcePercolation        QAMosquitoInspectionQAWaterSourceType = "Percolation"
	QAMosquitoInspectionQAWaterSourceRainRunoff         QAMosquitoInspectionQAWaterSourceType = "Rain Runoff"
	QAMosquitoInspectionQAWaterSourceTidal              QAMosquitoInspectionQAWaterSourceType = "Tidal"
	QAMosquitoInspectionQAWaterSourceWaterTable         QAMosquitoInspectionQAWaterSourceType = "Water Table"
)

type QAMosquitoInspectionQAWaterConditionsType string

const (
	QAMosquitoInspectionQAWaterConditionsRustmaterial         QAMosquitoInspectionQAWaterConditionsType = "rust material"
	QAMosquitoInspectionQAWaterConditionsClear                QAMosquitoInspectionQAWaterConditionsType = "Clear"
	QAMosquitoInspectionQAWaterConditionsCloudyfines          QAMosquitoInspectionQAWaterConditionsType = "Cloudy/fines"
	QAMosquitoInspectionQAWaterConditionsFloatingdebris       QAMosquitoInspectionQAWaterConditionsType = "Floating debris"
	QAMosquitoInspectionQAWaterConditionsSubmergeddecomdebris QAMosquitoInspectionQAWaterConditionsType = "Submerged/decom. debris"
)

type QAMosquitoInspectionMosquitoActionType string

const (
	QAMosquitoInspectionMosquitoActionTreatment        QAMosquitoInspectionMosquitoActionType = "Treatment"
	QAMosquitoInspectionMosquitoActionCoveredcontainer QAMosquitoInspectionMosquitoActionType = "Covered container"
	QAMosquitoInspectionMosquitoActionCleareddebris    QAMosquitoInspectionMosquitoActionType = "Cleared debris"
	QAMosquitoInspectionMosquitoActionMaintenance      QAMosquitoInspectionMosquitoActionType = "Maintenance"
)

type QAMosquitoInspectionNotInUITFType int16

const (
	QAMosquitoInspectionNotInUITFTrue  QAMosquitoInspectionNotInUITFType = 1
	QAMosquitoInspectionNotInUITFFalse QAMosquitoInspectionNotInUITFType = 0
)

type QAMosquitoInspectionQASiteTypeType string

const (
	QAMosquitoInspectionQASiteTypeDetentionPond QAMosquitoInspectionQASiteTypeType = "Detention Pond"
	QAMosquitoInspectionQASiteTypeDitch         QAMosquitoInspectionQASiteTypeType = "Ditch"
	QAMosquitoInspectionQASiteTypeLowArea       QAMosquitoInspectionQASiteTypeType = "Low Area"
	QAMosquitoInspectionQASiteTypeMangroveEdge  QAMosquitoInspectionQASiteTypeType = "Mangrove Edge"
	QAMosquitoInspectionQASiteTypePond          QAMosquitoInspectionQASiteTypeType = "Pond"
	QAMosquitoInspectionQASiteTypePondEdge      QAMosquitoInspectionQASiteTypeType = "Pond Edge"
	QAMosquitoInspectionQASiteTypeSwale         QAMosquitoInspectionQASiteTypeType = "Swale"
)

type QAMosquitoInspectionQALarvaeReasonType string

const (
	QAMosquitoInspectionQALarvaeReasonMissedArea      QAMosquitoInspectionQALarvaeReasonType = "Missed Area"
	QAMosquitoInspectionQALarvaeReasonNewSite         QAMosquitoInspectionQALarvaeReasonType = "New Site"
	QAMosquitoInspectionQALarvaeReasonNotVisited      QAMosquitoInspectionQALarvaeReasonType = "Not Visited"
	QAMosquitoInspectionQALarvaeReasonRateLow         QAMosquitoInspectionQALarvaeReasonType = "Rate Low"
	QAMosquitoInspectionQALarvaeReasonTreatedRecently QAMosquitoInspectionQALarvaeReasonType = "Treated Recently"
	QAMosquitoInspectionQALarvaeReasonUnknown         QAMosquitoInspectionQALarvaeReasonType = "Unknown"
	QAMosquitoInspectionQALarvaeReasonWrongProduct    QAMosquitoInspectionQALarvaeReasonType = "Wrong Product"
)

type QAMosquitoInspectionQAVegetationType string

const (
	QAMosquitoInspectionQAVegetationAlgae             QAMosquitoInspectionQAVegetationType = "Algae"
	QAMosquitoInspectionQAVegetationCattails          QAMosquitoInspectionQAVegetationType = "Cattails"
	QAMosquitoInspectionQAVegetationDuckweed          QAMosquitoInspectionQAVegetationType = "Duckweed"
	QAMosquitoInspectionQAVegetationGlasswort         QAMosquitoInspectionQAVegetationType = "Glasswort"
	QAMosquitoInspectionQAVegetationGrassonedge       QAMosquitoInspectionQAVegetationType = "Grass on edge"
	QAMosquitoInspectionQAVegetationMangrove          QAMosquitoInspectionQAVegetationType = "Mangrove"
	QAMosquitoInspectionQAVegetationMosquitofern      QAMosquitoInspectionQAVegetationType = "Mosquito fern"
	QAMosquitoInspectionQAVegetationMuskgrass         QAMosquitoInspectionQAVegetationType = "Muskgrass"
	QAMosquitoInspectionQAVegetationMyriophyllum      QAMosquitoInspectionQAVegetationType = "Myriophyllum"
	QAMosquitoInspectionQAVegetationOther             QAMosquitoInspectionQAVegetationType = "Other"
	QAMosquitoInspectionQAVegetationRottingvegetation QAMosquitoInspectionQAVegetationType = "Rotting vegetation"
	QAMosquitoInspectionQAVegetationSaltwort          QAMosquitoInspectionQAVegetationType = "Saltwort"
	QAMosquitoInspectionQAVegetationSedges            QAMosquitoInspectionQAVegetationType = "Sedges"
)

type QAMosquitoInspectionQAWaterMovementType string

const (
	QAMosquitoInspectionQAWaterMovementFast     QAMosquitoInspectionQAWaterMovementType = "Fast"
	QAMosquitoInspectionQAWaterMovementMedium   QAMosquitoInspectionQAWaterMovementType = "Medium"
	QAMosquitoInspectionQAWaterMovementNone     QAMosquitoInspectionQAWaterMovementType = "None"
	QAMosquitoInspectionQAWaterMovementSlow     QAMosquitoInspectionQAWaterMovementType = "Slow"
	QAMosquitoInspectionQAWaterMovementVerySlow QAMosquitoInspectionQAWaterMovementType = "Very Slow"
)

type QAMosquitoInspectionQASoilConditionType string

const (
	QAMosquitoInspectionQASoilConditionCracked      QAMosquitoInspectionQASoilConditionType = "Cracked"
	QAMosquitoInspectionQASoilConditionDry          QAMosquitoInspectionQASoilConditionType = "Dry"
	QAMosquitoInspectionQASoilConditionInundated    QAMosquitoInspectionQASoilConditionType = "Inundated"
	QAMosquitoInspectionQASoilConditionSaturated    QAMosquitoInspectionQASoilConditionType = "Saturated"
	QAMosquitoInspectionQASoilConditionSurfaceMoist QAMosquitoInspectionQASoilConditionType = "Surface Moist"
)

type QAMosquitoInspectionQAWaterDurationType string

const (
	QAMosquitoInspectionQAWaterDurationAboutmonth       QAMosquitoInspectionQAWaterDurationType = "~month"
	QAMosquitoInspectionQAWaterDurationAboutweek        QAMosquitoInspectionQAWaterDurationType = "~week"
	QAMosquitoInspectionQAWaterDurationLessThanOneweek  QAMosquitoInspectionQAWaterDurationType = "<1 week"
	QAMosquitoInspectionQAWaterDurationLessThanday      QAMosquitoInspectionQAWaterDurationType = "<day"
	QAMosquitoInspectionQAWaterDurationLessThanmonth    QAMosquitoInspectionQAWaterDurationType = "<month"
	QAMosquitoInspectionQAWaterDurationGreaterThanmonth QAMosquitoInspectionQAWaterDurationType = ">month"
	QAMosquitoInspectionQAWaterDurationGreaterThanweek  QAMosquitoInspectionQAWaterDurationType = ">week"
)

type QAMosquitoInspectionQABreedingPotentialType string

const (
	QAMosquitoInspectionQABreedingPotentialHigh   QAMosquitoInspectionQABreedingPotentialType = "High"
	QAMosquitoInspectionQABreedingPotentialLow    QAMosquitoInspectionQABreedingPotentialType = "Low"
	QAMosquitoInspectionQABreedingPotentialMedium QAMosquitoInspectionQABreedingPotentialType = "Medium"
	QAMosquitoInspectionQABreedingPotentialRare   QAMosquitoInspectionQABreedingPotentialType = "Rare"
)

type QAMosquitoInspectionQAMosquitoHabitatType string

const (
	QAMosquitoInspectionQAMosquitoHabitatDepressions       QAMosquitoInspectionQAMosquitoHabitatType = "Depressions"
	QAMosquitoInspectionQAMosquitoHabitatDetrituspresent   QAMosquitoInspectionQAMosquitoHabitatType = "Detritus present"
	QAMosquitoInspectionQAMosquitoHabitatFast              QAMosquitoInspectionQAMosquitoHabitatType = "Fast"
	QAMosquitoInspectionQAMosquitoHabitatFewpredators      QAMosquitoInspectionQAMosquitoHabitatType = "Few predators"
	QAMosquitoInspectionQAMosquitoHabitatFluctuatinglevels QAMosquitoInspectionQAMosquitoHabitatType = "Fluctuating levels"
	QAMosquitoInspectionQAMosquitoHabitatH20LessThanSix    QAMosquitoInspectionQAMosquitoHabitatType = "H20<6"
	QAMosquitoInspectionQAMosquitoHabitatLowwavepotential  QAMosquitoInspectionQAMosquitoHabitatType = "Low wave potential"
	QAMosquitoInspectionQAMosquitoHabitatNofish            QAMosquitoInspectionQAMosquitoHabitatType = "No fish"
	QAMosquitoInspectionQAMosquitoHabitatShallowedges      QAMosquitoInspectionQAMosquitoHabitatType = "Shallow edges"
	QAMosquitoInspectionQAMosquitoHabitatStillwateredges   QAMosquitoInspectionQAMosquitoHabitatType = "Still water edges"
	QAMosquitoInspectionQAMosquitoHabitatStillwaterwhole   QAMosquitoInspectionQAMosquitoHabitatType = "Still water whole"
	QAMosquitoInspectionQAMosquitoHabitatVegonedges        QAMosquitoInspectionQAMosquitoHabitatType = "Veg. on edges"
)

type QAMosquitoInspectionQAAquaticOrganismsType string

const (
	QAMosquitoInspectionQAAquaticOrganismsFish   QAMosquitoInspectionQAAquaticOrganismsType = "fish"
	QAMosquitoInspectionQAAquaticOrganismsScuds  QAMosquitoInspectionQAAquaticOrganismsType = "scuds"
	QAMosquitoInspectionQAAquaticOrganismsSnails QAMosquitoInspectionQAAquaticOrganismsType = "snails"
)

type QAMosquitoInspectionQASourceReductionType string

const (
	QAMosquitoInspectionQASourceReductionOnetractorLessThanday QAMosquitoInspectionQASourceReductionType = "1 tractor < day"
	QAMosquitoInspectionQASourceReductionAdjustfloodirrigation QAMosquitoInspectionQASourceReductionType = "adjust flood irrigation"
	QAMosquitoInspectionQASourceReductionAdjustturfirrigation  QAMosquitoInspectionQASourceReductionType = "adjust turf irrigation"
	QAMosquitoInspectionQASourceReductionClearoutflow          QAMosquitoInspectionQASourceReductionType = "clear outflow"
	QAMosquitoInspectionQASourceReductionCutditch              QAMosquitoInspectionQASourceReductionType = "cut ditch"
	QAMosquitoInspectionQASourceReductionHandgrading           QAMosquitoInspectionQASourceReductionType = "hand grading"
	QAMosquitoInspectionQASourceReductionLaserleveling         QAMosquitoInspectionQASourceReductionType = "laser leveling"
	QAMosquitoInspectionQASourceReductionMultipleloadssoil     QAMosquitoInspectionQASourceReductionType = "multiple loads soil"
	QAMosquitoInspectionQASourceReductionNone                  QAMosquitoInspectionQASourceReductionType = "none"
)

type QAMosquitoInspection struct {
	Objectid                  uint                                        `field:"OBJECTID"`
	PositiveDips              int16                                       `field:"POSDIPS"`
	Action                    QAMosquitoInspectionMosquitoActionType      `field:"ACTIONTAKEN"`
	Comments                  string                                      `field:"COMMENTS"`
	AverageTemperature        float64                                     `field:"AVETEMP"`
	WindSpeed                 float64                                     `field:"WINDSPEED"`
	RainGauge                 float64                                     `field:"RAINGAUGE"`
	GlobalID                  uuid.UUID                                   `field:"GlobalID"`
	Start                     time.Time                                   `field:"STARTDATETIME"`
	Finish                    time.Time                                   `field:"ENDDATETIME"`
	WindDirection             string                                      `field:"WINDDIR"`
	Reviewed                  QAMosquitoInspectionNotInUITFType           `field:"REVIEWED"`
	Reviewedby                string                                      `field:"REVIEWEDBY"`
	Revieweddate              time.Time                                   `field:"REVIEWEDDATE"`
	Locationname              string                                      `field:"LOCATIONNAME"`
	Zone                      string                                      `field:"ZONE"`
	Recordstatus              int16                                       `field:"RECORDSTATUS"`
	Zone2                     string                                      `field:"ZONE2"`
	LandingRate               int16                                       `field:"LR"`
	NegativeDips              int16                                       `field:"NEGDIPS"`
	TotalAcres                float64                                     `field:"TOTALACRES"`
	AcresBreeding             float64                                     `field:"ACRESBREEDING"`
	FishPresent               QAMosquitoInspectionNotInUITFType           `field:"FISH"`
	SiteType                  QAMosquitoInspectionQASiteTypeType          `field:"SITETYPE"`
	BreedingPotential         QAMosquitoInspectionQABreedingPotentialType `field:"BREEDINGPOTENTIAL"`
	MovingWater               QAMosquitoInspectionNotInUITFType           `field:"MOVINGWATER"`
	NoEvidenceOfWaterEver     QAMosquitoInspectionNotInUITFType           `field:"NOWATEREVER"`
	MosquitoHabitatIndicators QAMosquitoInspectionQAMosquitoHabitatType   `field:"MOSQUITOHABITAT"`
	HabitatValue              int16                                       `field:"HABVALUE1"`
	Habvalue1percent          int16                                       `field:"HABVALUE1PERCENT"`
	HabitatValue2             int16                                       `field:"HABVALUE2"`
	Habvalue2percent          int16                                       `field:"HABVALUE2PERCENT"`
	Potential                 int16                                       `field:"POTENTIAL"`
	LarvaePresent             QAMosquitoInspectionNotInUITFType           `field:"LARVAEPRESENT"`
	LarvaeInsideTreatedArea   QAMosquitoInspectionNotInUITFType           `field:"LARVAEINSIDETREATEDAREA"`
	LarvaeOutsideTreatedArea  QAMosquitoInspectionNotInUITFType           `field:"LARVAEOUTSIDETREATEDAREA"`
	ReasonLarvaePresent       QAMosquitoInspectionQALarvaeReasonType      `field:"LARVAEREASON"`
	AquaticOrganisms          QAMosquitoInspectionQAAquaticOrganismsType  `field:"AQUATICORGANISMS"`
	Vegetation                QAMosquitoInspectionQAVegetationType        `field:"VEGETATION"`
	SourceReduction           QAMosquitoInspectionQASourceReductionType   `field:"SOURCEREDUCTION"`
	WaterPresent              QAMosquitoInspectionNotInUITFType           `field:"WATERPRESENT"`
	WaterMovement             QAMosquitoInspectionQAWaterMovementType     `field:"WATERMOVEMENT1"`
	Watermovement1percent     int16                                       `field:"WATERMOVEMENT1PERCENT"`
	WaterMovement2            QAMosquitoInspectionQAWaterMovementType     `field:"WATERMOVEMENT2"`
	Watermovement2percent     int16                                       `field:"WATERMOVEMENT2PERCENT"`
	SoilConditions            QAMosquitoInspectionQASoilConditionType     `field:"SOILCONDITIONS"`
	HowLongWaterPresent       QAMosquitoInspectionQAWaterDurationType     `field:"WATERDURATION"`
	WaterSource               QAMosquitoInspectionQAWaterSourceType       `field:"WATERSOURCE"`
	WaterConditions           QAMosquitoInspectionQAWaterConditionsType   `field:"WATERCONDITIONS"`
	AdultActivity             QAMosquitoInspectionNotInUITFType           `field:"ADULTACTIVITY"`
	Linelocid                 uuid.UUID                                   `field:"LINELOCID"`
	Pointlocid                uuid.UUID                                   `field:"POINTLOCID"`
	Polygonlocid              uuid.UUID                                   `field:"POLYGONLOCID"`
	CreatedUser               string                                      `field:"created_user"`
	CreatedDate               time.Time                                   `field:"created_date"`
	LastEditedUser            string                                      `field:"last_edited_user"`
	LastEditedDate            time.Time                                   `field:"last_edited_date"`
	FieldTech                 string                                      `field:"FIELDTECH"`
	CreationDate              time.Time                                   `field:"CreationDate"`
	Creator                   string                                      `field:"Creator"`
	EditDate                  time.Time                                   `field:"EditDate"`
	Editor                    string                                      `field:"Editor"`
}
