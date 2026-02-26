package response

type TileInfo struct {
	Rows               int              `json:"rows"`
	Cols               int              `json:"cols"`
	DPI                int              `json:"dpi"`
	Format             string           `json:"format"`
	CompressionQuality int              `json:"compressionQuality"`
	Origin             Point            `json:"origin"`
	SpatialReference   SpatialReference `json:"spatialReference"`
	LODs               []LOD            `json:"lods"`
}
