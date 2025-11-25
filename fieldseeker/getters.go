package fieldseeker

import (
	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker/layer"
)

func (fs *FieldSeeker) LocationTracking(offset uint) ([]*layer.LocationTracking, error) {
	return featureToStruct[layer.LocationTracking](fs, "LocationTracking", 0, offset)
}
func (fs *FieldSeeker) Tracklog(offset uint) ([]*layer.Tracklog, error) {
	return featureToStruct[layer.Tracklog](fs, "Tracklog", 1, offset)
}
func (fs *FieldSeeker) ServiceRequest(offset uint) ([]*layer.ServiceRequest, error) {
	return featureToStruct[layer.ServiceRequest](fs, "ServiceRequest", 2, offset)
}
func (fs *FieldSeeker) TrapLocation(offset uint) ([]*layer.TrapLocation, error) {
	return featureToStruct[layer.TrapLocation](fs, "TrapLocation", 3, offset)
}
func (fs *FieldSeeker) LandingCountLocation(offset uint) ([]*layer.LandingCountLocation, error) {
	return featureToStruct[layer.LandingCountLocation](fs, "LandingCountLocation", 4, offset)
}
func (fs *FieldSeeker) SampleLocation(offset uint) ([]*layer.SampleLocation, error) {
	return featureToStruct[layer.SampleLocation](fs, "SampleLocation", 5, offset)
}
func (fs *FieldSeeker) ContainerRelate(offset uint) ([]*layer.ContainerRelate, error) {
	return featureToStruct[layer.ContainerRelate](fs, "ContainerRelate", 6, offset)
}
func (fs *FieldSeeker) HabitatRelate(offset uint) ([]*layer.HabitatRelate, error) {
	return featureToStruct[layer.HabitatRelate](fs, "HabitatRelate", 7, offset)
}
func (fs *FieldSeeker) PoolDetail(offset uint) ([]*layer.PoolDetail, error) {
	return featureToStruct[layer.PoolDetail](fs, "PoolDetail", 8, offset)
}
func (fs *FieldSeeker) Pool(offset uint) ([]*layer.Pool, error) {
	return featureToStruct[layer.Pool](fs, "Pool", 9, offset)
}

func (fs *FieldSeeker) SpeciesAbundance(offset uint) ([]*layer.SpeciesAbundance, error) {
	return featureToStruct[layer.SpeciesAbundance](fs, "SpeciesAbundance", 10, offset)
}
func (fs *FieldSeeker) PointLocation(offset uint) ([]*layer.PointLocation, error) {
	return featureToStruct[layer.PointLocation](fs, "PointLocation", 11, offset)
}
func (fs *FieldSeeker) InspectionSample(offset uint) ([]*layer.InspectionSample, error) {
	return featureToStruct[layer.InspectionSample](fs, "InspectionSample", 12, offset)
}
func (fs *FieldSeeker) InspectionSampleDetail(offset uint) ([]*layer.InspectionSampleDetail, error) {
	return featureToStruct[layer.InspectionSampleDetail](fs, "InspectionSampleDetail", 13, offset)
}
func (fs *FieldSeeker) MosquitoInspection(offset uint) ([]*layer.MosquitoInspection, error) {
	return featureToStruct[layer.MosquitoInspection](fs, "MosquitoInspection", 14, offset)
}
func (fs *FieldSeeker) TrapData(offset uint) ([]*layer.TrapData, error) {
	return featureToStruct[layer.TrapData](fs, "TrapData", 15, offset)
}
func (fs *FieldSeeker) LandingCount(offset uint) ([]*layer.LandingCount, error) {
	return featureToStruct[layer.LandingCount](fs, "LandingCount", 16, offset)
}
func (fs *FieldSeeker) TimeCard(offset uint) ([]*layer.TimeCard, error) {
	return featureToStruct[layer.TimeCard](fs, "TimeCard", 17, offset)
}
func (fs *FieldSeeker) Treatment(offset uint) ([]*layer.Treatment, error) {
	return featureToStruct[layer.Treatment](fs, "Treatment", 18, offset)
}
func (fs *FieldSeeker) SampleCollection(offset uint) ([]*layer.SampleCollection, error) {
	return featureToStruct[layer.SampleCollection](fs, "SampleCollection", 19, offset)
}
func (fs *FieldSeeker) StormDrain(offset uint) ([]*layer.StormDrain, error) {
	return featureToStruct[layer.StormDrain](fs, "StormDrain", 20, offset)
}
func (fs *FieldSeeker) QAProductObservation(offset uint) ([]*layer.QAProductObservation, error) {
	return featureToStruct[layer.QAProductObservation](fs, "QAProductObservation", 21, offset)
}
func (fs *FieldSeeker) QALarvCount(offset uint) ([]*layer.QALarvCount, error) {
	return featureToStruct[layer.QALarvCount](fs, "QALarvCount", 22, offset)
}
func (fs *FieldSeeker) QAMosquitoInspection(offset uint) ([]*layer.QAMosquitoInspection, error) {
	return featureToStruct[layer.QAMosquitoInspection](fs, "QAMosquitoInspection", 23, offset)
}
func (fs *FieldSeeker) FieldScoutingLog(offset uint) ([]*layer.FieldScoutingLog, error) {
	return featureToStruct[layer.FieldScoutingLog](fs, "FieldScoutingLog", 24, offset)
}
func (fs *FieldSeeker) BarrierSpray(offset uint) ([]*layer.BarrierSpray, error) {
	return featureToStruct[layer.BarrierSpray](fs, "BarrierSpray", 25, offset)
}
func (fs *FieldSeeker) BarrierSprayRoute(offset uint) ([]*layer.BarrierSprayRoute, error) {
	return featureToStruct[layer.BarrierSprayRoute](fs, "BarrierSprayRoute", 26, offset)
}
func (fs *FieldSeeker) LineLocation(offset uint) ([]*layer.LineLocation, error) {
	return featureToStruct[layer.LineLocation](fs, "LineLocation", 27, offset)
}
func (fs *FieldSeeker) ULVSprayRoute(offset uint) ([]*layer.ULVSprayRoute, error) {
	return featureToStruct[layer.ULVSprayRoute](fs, "ULVSprayRoute", 28, offset)
}
func (fs *FieldSeeker) OfflineMapAreas(offset uint) ([]*layer.OfflineMapAreas, error) {
	return featureToStruct[layer.OfflineMapAreas](fs, "OfflineMapAreas", 29, offset)
}
func (fs *FieldSeeker) TreatmentArea(offset uint) ([]*layer.TreatmentArea, error) {
	return featureToStruct[layer.TreatmentArea](fs, "TreatmentArea", 30, offset)
}
func (fs *FieldSeeker) RestrictedArea(offset uint) ([]*layer.RestrictedArea, error) {
	return featureToStruct[layer.RestrictedArea](fs, "RestrictedArea", 31, offset)
}
func (fs *FieldSeeker) ProposedTreatmentArea(offset uint) ([]*layer.ProposedTreatmentArea, error) {
	return featureToStruct[layer.ProposedTreatmentArea](fs, "ProposedTreatmentArea", 32, offset)
}
func (fs *FieldSeeker) PolygonLocation(offset uint) ([]*layer.PolygonLocation, error) {
	return featureToStruct[layer.PolygonLocation](fs, "PolygonLocation", 33, offset)
}
func (fs *FieldSeeker) Zones2(offset uint) ([]*layer.Zones2, error) {
	return featureToStruct[layer.Zones2](fs, "Zones2", 34, offset)
}
func (fs *FieldSeeker) Zones(offset uint) ([]*layer.Zones, error) {
	return featureToStruct[layer.Zones](fs, "Zones", 35, offset)
}
func (fs *FieldSeeker) AerialSpraySession(offset uint) ([]*layer.AerialSpraySession, error) {
	return featureToStruct[layer.AerialSpraySession](fs, "AerialSpraySession", 45, offset)
}
func (fs *FieldSeeker) AerialSprayLine(offset uint) ([]*layer.AerialSprayLine, error) {
	return featureToStruct[layer.AerialSprayLine](fs, "AerialSprayLine", 46, offset)
}
func (fs *FieldSeeker) PoolBuffer(offset uint) ([]*layer.PoolBuffer, error) {
	return featureToStruct[layer.PoolBuffer](fs, "PoolBuffer", 51, offset)
}
func (fs *FieldSeeker) RodentInspection(offset uint) ([]*layer.RodentInspection, error) {
	return featureToStruct[layer.RodentInspection](fs, "RodentInspection", 52, offset)
}
func (fs *FieldSeeker) RodentLocation(offset uint) ([]*layer.RodentLocation, error) {
	return featureToStruct[layer.RodentLocation](fs, "RodentLocation", 53, offset)
}
