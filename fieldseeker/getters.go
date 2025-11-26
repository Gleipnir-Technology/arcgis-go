package fieldseeker

import (
	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker/layer"
)

func (fs *FieldSeeker) LocationTracking(offset uint) ([]*layer.LocationTracking, error) {
	return featureToStruct[layer.LocationTracking](fs, LayerLocationTracking, offset)
}
func (fs *FieldSeeker) Tracklog(offset uint) ([]*layer.Tracklog, error) {
	return featureToStruct[layer.Tracklog](fs, LayerTracklog, offset)
}
func (fs *FieldSeeker) ServiceRequest(offset uint) ([]*layer.ServiceRequest, error) {
	return featureToStruct[layer.ServiceRequest](fs, LayerServiceRequest, offset)
}
func (fs *FieldSeeker) TrapLocation(offset uint) ([]*layer.TrapLocation, error) {
	return featureToStruct[layer.TrapLocation](fs, LayerTrapLocation, offset)
}
func (fs *FieldSeeker) LandingCountLocation(offset uint) ([]*layer.LandingCountLocation, error) {
	return featureToStruct[layer.LandingCountLocation](fs, LayerLandingCountLocation, offset)
}
func (fs *FieldSeeker) SampleLocation(offset uint) ([]*layer.SampleLocation, error) {
	return featureToStruct[layer.SampleLocation](fs, LayerSampleLocation, offset)
}
func (fs *FieldSeeker) ContainerRelate(offset uint) ([]*layer.ContainerRelate, error) {
	return featureToStruct[layer.ContainerRelate](fs, LayerContainerRelate, offset)
}
func (fs *FieldSeeker) HabitatRelate(offset uint) ([]*layer.HabitatRelate, error) {
	return featureToStruct[layer.HabitatRelate](fs, LayerHabitatRelate, offset)
}
func (fs *FieldSeeker) PoolDetail(offset uint) ([]*layer.PoolDetail, error) {
	return featureToStruct[layer.PoolDetail](fs, LayerPoolDetail, offset)
}
func (fs *FieldSeeker) Pool(offset uint) ([]*layer.Pool, error) {
	return featureToStruct[layer.Pool](fs, LayerPool, offset)
}

func (fs *FieldSeeker) SpeciesAbundance(offset uint) ([]*layer.SpeciesAbundance, error) {
	return featureToStruct[layer.SpeciesAbundance](fs, LayerSpeciesAbundance, offset)
}
func (fs *FieldSeeker) PointLocation(offset uint) ([]*layer.PointLocation, error) {
	return featureToStruct[layer.PointLocation](fs, LayerPointLocation, offset)
}
func (fs *FieldSeeker) InspectionSample(offset uint) ([]*layer.InspectionSample, error) {
	return featureToStruct[layer.InspectionSample](fs, LayerInspectionSample, offset)
}
func (fs *FieldSeeker) InspectionSampleDetail(offset uint) ([]*layer.InspectionSampleDetail, error) {
	return featureToStruct[layer.InspectionSampleDetail](fs, LayerInspectionSampleDetail, offset)
}
func (fs *FieldSeeker) MosquitoInspection(offset uint) ([]*layer.MosquitoInspection, error) {
	return featureToStruct[layer.MosquitoInspection](fs, LayerMosquitoInspection, offset)
}
func (fs *FieldSeeker) TrapData(offset uint) ([]*layer.TrapData, error) {
	return featureToStruct[layer.TrapData](fs, LayerTrapData, offset)
}
func (fs *FieldSeeker) LandingCount(offset uint) ([]*layer.LandingCount, error) {
	return featureToStruct[layer.LandingCount](fs, LayerLandingCount, offset)
}
func (fs *FieldSeeker) TimeCard(offset uint) ([]*layer.TimeCard, error) {
	return featureToStruct[layer.TimeCard](fs, LayerTimeCard, offset)
}
func (fs *FieldSeeker) Treatment(offset uint) ([]*layer.Treatment, error) {
	return featureToStruct[layer.Treatment](fs, LayerTreatment, offset)
}
func (fs *FieldSeeker) SampleCollection(offset uint) ([]*layer.SampleCollection, error) {
	return featureToStruct[layer.SampleCollection](fs, LayerSampleCollection, offset)
}
func (fs *FieldSeeker) StormDrain(offset uint) ([]*layer.StormDrain, error) {
	return featureToStruct[layer.StormDrain](fs, LayerStormDrain, offset)
}
func (fs *FieldSeeker) QAProductObservation(offset uint) ([]*layer.QAProductObservation, error) {
	return featureToStruct[layer.QAProductObservation](fs, LayerQAProductObservation, offset)
}
func (fs *FieldSeeker) QALarvCount(offset uint) ([]*layer.QALarvCount, error) {
	return featureToStruct[layer.QALarvCount](fs, LayerQALarvCount, offset)
}
func (fs *FieldSeeker) QAMosquitoInspection(offset uint) ([]*layer.QAMosquitoInspection, error) {
	return featureToStruct[layer.QAMosquitoInspection](fs, LayerQAMosquitoInspection, offset)
}
func (fs *FieldSeeker) FieldScoutingLog(offset uint) ([]*layer.FieldScoutingLog, error) {
	return featureToStruct[layer.FieldScoutingLog](fs, LayerFieldScoutingLog, offset)
}
func (fs *FieldSeeker) BarrierSpray(offset uint) ([]*layer.BarrierSpray, error) {
	return featureToStruct[layer.BarrierSpray](fs, LayerBarrierSpray, offset)
}
func (fs *FieldSeeker) BarrierSprayRoute(offset uint) ([]*layer.BarrierSprayRoute, error) {
	return featureToStruct[layer.BarrierSprayRoute](fs, LayerBarrierSprayRoute, offset)
}
func (fs *FieldSeeker) LineLocation(offset uint) ([]*layer.LineLocation, error) {
	return featureToStruct[layer.LineLocation](fs, LayerLineLocation, offset)
}
func (fs *FieldSeeker) ULVSprayRoute(offset uint) ([]*layer.ULVSprayRoute, error) {
	return featureToStruct[layer.ULVSprayRoute](fs, LayerULVSprayRoute, offset)
}
func (fs *FieldSeeker) OfflineMapAreas(offset uint) ([]*layer.OfflineMapAreas, error) {
	return featureToStruct[layer.OfflineMapAreas](fs, LayerOfflineMapAreas, offset)
}
func (fs *FieldSeeker) TreatmentArea(offset uint) ([]*layer.TreatmentArea, error) {
	return featureToStruct[layer.TreatmentArea](fs, LayerTreatmentArea, offset)
}
func (fs *FieldSeeker) RestrictedArea(offset uint) ([]*layer.RestrictedArea, error) {
	return featureToStruct[layer.RestrictedArea](fs, LayerRestrictedArea, offset)
}
func (fs *FieldSeeker) ProposedTreatmentArea(offset uint) ([]*layer.ProposedTreatmentArea, error) {
	return featureToStruct[layer.ProposedTreatmentArea](fs, LayerProposedTreatmentArea, offset)
}
func (fs *FieldSeeker) PolygonLocation(offset uint) ([]*layer.PolygonLocation, error) {
	return featureToStruct[layer.PolygonLocation](fs, LayerPolygonLocation, offset)
}
func (fs *FieldSeeker) Zones2(offset uint) ([]*layer.Zones2, error) {
	return featureToStruct[layer.Zones2](fs, LayerZones2, offset)
}
func (fs *FieldSeeker) Zones(offset uint) ([]*layer.Zones, error) {
	return featureToStruct[layer.Zones](fs, LayerZones, offset)
}
func (fs *FieldSeeker) AerialSpraySession(offset uint) ([]*layer.AerialSpraySession, error) {
	return featureToStruct[layer.AerialSpraySession](fs, LayerAerialSpraySession, offset)
}
func (fs *FieldSeeker) AerialSprayLine(offset uint) ([]*layer.AerialSprayLine, error) {
	return featureToStruct[layer.AerialSprayLine](fs, LayerAerialSprayLine, offset)
}
func (fs *FieldSeeker) PoolBuffer(offset uint) ([]*layer.PoolBuffer, error) {
	return featureToStruct[layer.PoolBuffer](fs, LayerPoolBuffer, offset)
}
func (fs *FieldSeeker) RodentInspection(offset uint) ([]*layer.RodentInspection, error) {
	return featureToStruct[layer.RodentInspection](fs, LayerRodentInspection, offset)
}
func (fs *FieldSeeker) RodentLocation(offset uint) ([]*layer.RodentLocation, error) {
	return featureToStruct[layer.RodentLocation](fs, LayerRodentLocation, offset)
}
