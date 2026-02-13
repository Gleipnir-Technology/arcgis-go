package response

type Basemap struct {
	BasemapLayers []LayerResource `json:"baseMapLayers"`
	Title         string          `json:"title"`
}
