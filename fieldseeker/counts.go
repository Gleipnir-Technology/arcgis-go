package fieldseeker

func (fs *FieldSeeker) LayerCount(layer_id uint) (uint, error) {
	return 0, nil
}

func (fs *FieldSeeker) LocationTrackingCount() (uint, error) {
	return fs.LayerCount(fs.layerToID[LayerLocationTracking])
}
