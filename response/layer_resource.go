package response

type LayerResource struct {
	ID        string `json:"id"`
	LayerType string `json:"layerType"`
	//ResourceInfo ResourceInfo `json:"resourceInfo"`
	URL        string `json:"url"`
	Visibility bool   `json:"visibility"`
	Opacity    int    `json:"opacity"`
	Title      string `json:"title"`
}
