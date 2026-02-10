package fieldseeker

import (
	"context"

	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker/layer"
)

func (fs *FieldSeeker) LocationTracking(ctx context.Context, offset uint) ([]*layer.LocationTracking, error) {
	return featureToStruct[layer.LocationTracking](ctx, fs, LayerLocationTracking, offset)
}
func (fs *FieldSeeker) Tracklog(ctx context.Context, offset uint) ([]*layer.Tracklog, error) {
	return featureToStruct[layer.Tracklog](ctx, fs, LayerTracklog, offset)
}
func (fs *FieldSeeker) ServiceRequest(ctx context.Context, offset uint) ([]*layer.ServiceRequest, error) {
	return featureToStruct[layer.ServiceRequest](ctx, fs, LayerServiceRequest, offset)
}
func (fs *FieldSeeker) TrapLocation(ctx context.Context, offset uint) ([]*layer.TrapLocation, error) {
	return featureToStruct[layer.TrapLocation](ctx, fs, LayerTrapLocation, offset)
}
func (fs *FieldSeeker) LandingCountLocation(ctx context.Context, offset uint) ([]*layer.LandingCountLocation, error) {
	return featureToStruct[layer.LandingCountLocation](ctx, fs, LayerLandingCountLocation, offset)
}
func (fs *FieldSeeker) SampleLocation(ctx context.Context, offset uint) ([]*layer.SampleLocation, error) {
	return featureToStruct[layer.SampleLocation](ctx, fs, LayerSampleLocation, offset)
}
func (fs *FieldSeeker) ContainerRelate(ctx context.Context, offset uint) ([]*layer.ContainerRelate, error) {
	return featureToStruct[layer.ContainerRelate](ctx, fs, LayerContainerRelate, offset)
}
func (fs *FieldSeeker) HabitatRelate(ctx context.Context, offset uint) ([]*layer.HabitatRelate, error) {
	return featureToStruct[layer.HabitatRelate](ctx, fs, LayerHabitatRelate, offset)
}
func (fs *FieldSeeker) PoolDetail(ctx context.Context, offset uint) ([]*layer.PoolDetail, error) {
	return featureToStruct[layer.PoolDetail](ctx, fs, LayerPoolDetail, offset)
}
func (fs *FieldSeeker) Pool(ctx context.Context, offset uint) ([]*layer.Pool, error) {
	return featureToStruct[layer.Pool](ctx, fs, LayerPool, offset)
}

func (fs *FieldSeeker) SpeciesAbundance(ctx context.Context, offset uint) ([]*layer.SpeciesAbundance, error) {
	return featureToStruct[layer.SpeciesAbundance](ctx, fs, LayerSpeciesAbundance, offset)
}
func (fs *FieldSeeker) PointLocation(ctx context.Context, offset uint) ([]*layer.PointLocation, error) {
	return featureToStruct[layer.PointLocation](ctx, fs, LayerPointLocation, offset)
}
func (fs *FieldSeeker) InspectionSample(ctx context.Context, offset uint) ([]*layer.InspectionSample, error) {
	return featureToStruct[layer.InspectionSample](ctx, fs, LayerInspectionSample, offset)
}
func (fs *FieldSeeker) InspectionSampleDetail(ctx context.Context, offset uint) ([]*layer.InspectionSampleDetail, error) {
	return featureToStruct[layer.InspectionSampleDetail](ctx, fs, LayerInspectionSampleDetail, offset)
}
func (fs *FieldSeeker) MosquitoInspection(ctx context.Context, offset uint) ([]*layer.MosquitoInspection, error) {
	return featureToStruct[layer.MosquitoInspection](ctx, fs, LayerMosquitoInspection, offset)
}
func (fs *FieldSeeker) TrapData(ctx context.Context, offset uint) ([]*layer.TrapData, error) {
	return featureToStruct[layer.TrapData](ctx, fs, LayerTrapData, offset)
}
func (fs *FieldSeeker) LandingCount(ctx context.Context, offset uint) ([]*layer.LandingCount, error) {
	return featureToStruct[layer.LandingCount](ctx, fs, LayerLandingCount, offset)
}
func (fs *FieldSeeker) TimeCard(ctx context.Context, offset uint) ([]*layer.TimeCard, error) {
	return featureToStruct[layer.TimeCard](ctx, fs, LayerTimeCard, offset)
}
func (fs *FieldSeeker) Treatment(ctx context.Context, offset uint) ([]*layer.Treatment, error) {
	return featureToStruct[layer.Treatment](ctx, fs, LayerTreatment, offset)
}
func (fs *FieldSeeker) SampleCollection(ctx context.Context, offset uint) ([]*layer.SampleCollection, error) {
	return featureToStruct[layer.SampleCollection](ctx, fs, LayerSampleCollection, offset)
}
func (fs *FieldSeeker) StormDrain(ctx context.Context, offset uint) ([]*layer.StormDrain, error) {
	return featureToStruct[layer.StormDrain](ctx, fs, LayerStormDrain, offset)
}
func (fs *FieldSeeker) QAProductObservation(ctx context.Context, offset uint) ([]*layer.QAProductObservation, error) {
	return featureToStruct[layer.QAProductObservation](ctx, fs, LayerQAProductObservation, offset)
}
func (fs *FieldSeeker) QALarvCount(ctx context.Context, offset uint) ([]*layer.QALarvCount, error) {
	return featureToStruct[layer.QALarvCount](ctx, fs, LayerQALarvCount, offset)
}
func (fs *FieldSeeker) QAMosquitoInspection(ctx context.Context, offset uint) ([]*layer.QAMosquitoInspection, error) {
	return featureToStruct[layer.QAMosquitoInspection](ctx, fs, LayerQAMosquitoInspection, offset)
}
func (fs *FieldSeeker) FieldScoutingLog(ctx context.Context, offset uint) ([]*layer.FieldScoutingLog, error) {
	return featureToStruct[layer.FieldScoutingLog](ctx, fs, LayerFieldScoutingLog, offset)
}
func (fs *FieldSeeker) BarrierSpray(ctx context.Context, offset uint) ([]*layer.BarrierSpray, error) {
	return featureToStruct[layer.BarrierSpray](ctx, fs, LayerBarrierSpray, offset)
}
func (fs *FieldSeeker) BarrierSprayRoute(ctx context.Context, offset uint) ([]*layer.BarrierSprayRoute, error) {
	return featureToStruct[layer.BarrierSprayRoute](ctx, fs, LayerBarrierSprayRoute, offset)
}
func (fs *FieldSeeker) LineLocation(ctx context.Context, offset uint) ([]*layer.LineLocation, error) {
	return featureToStruct[layer.LineLocation](ctx, fs, LayerLineLocation, offset)
}
func (fs *FieldSeeker) ULVSprayRoute(ctx context.Context, offset uint) ([]*layer.ULVSprayRoute, error) {
	return featureToStruct[layer.ULVSprayRoute](ctx, fs, LayerULVSprayRoute, offset)
}
func (fs *FieldSeeker) OfflineMapAreas(ctx context.Context, offset uint) ([]*layer.OfflineMapAreas, error) {
	return featureToStruct[layer.OfflineMapAreas](ctx, fs, LayerOfflineMapAreas, offset)
}
func (fs *FieldSeeker) TreatmentArea(ctx context.Context, offset uint) ([]*layer.TreatmentArea, error) {
	return featureToStruct[layer.TreatmentArea](ctx, fs, LayerTreatmentArea, offset)
}
func (fs *FieldSeeker) RestrictedArea(ctx context.Context, offset uint) ([]*layer.RestrictedArea, error) {
	return featureToStruct[layer.RestrictedArea](ctx, fs, LayerRestrictedArea, offset)
}
func (fs *FieldSeeker) ProposedTreatmentArea(ctx context.Context, offset uint) ([]*layer.ProposedTreatmentArea, error) {
	return featureToStruct[layer.ProposedTreatmentArea](ctx, fs, LayerProposedTreatmentArea, offset)
}
func (fs *FieldSeeker) PolygonLocation(ctx context.Context, offset uint) ([]*layer.PolygonLocation, error) {
	return featureToStruct[layer.PolygonLocation](ctx, fs, LayerPolygonLocation, offset)
}
func (fs *FieldSeeker) Zones2(ctx context.Context, offset uint) ([]*layer.Zones2, error) {
	return featureToStruct[layer.Zones2](ctx, fs, LayerZones2, offset)
}
func (fs *FieldSeeker) Zones(ctx context.Context, offset uint) ([]*layer.Zones, error) {
	return featureToStruct[layer.Zones](ctx, fs, LayerZones, offset)
}
func (fs *FieldSeeker) AerialSpraySession(ctx context.Context, offset uint) ([]*layer.AerialSpraySession, error) {
	return featureToStruct[layer.AerialSpraySession](ctx, fs, LayerAerialSpraySession, offset)
}
func (fs *FieldSeeker) AerialSprayLine(ctx context.Context, offset uint) ([]*layer.AerialSprayLine, error) {
	return featureToStruct[layer.AerialSprayLine](ctx, fs, LayerAerialSprayLine, offset)
}
func (fs *FieldSeeker) PoolBuffer(ctx context.Context, offset uint) ([]*layer.PoolBuffer, error) {
	return featureToStruct[layer.PoolBuffer](ctx, fs, LayerPoolBuffer, offset)
}
func (fs *FieldSeeker) RodentInspection(ctx context.Context, offset uint) ([]*layer.RodentInspection, error) {
	return featureToStruct[layer.RodentInspection](ctx, fs, LayerRodentInspection, offset)
}
func (fs *FieldSeeker) RodentLocation(ctx context.Context, offset uint) ([]*layer.RodentLocation, error) {
	return featureToStruct[layer.RodentLocation](ctx, fs, LayerRodentLocation, offset)
}
