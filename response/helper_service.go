package response

type Service struct {
	URL string `json:"url"`
}
type HelperServices struct {
	Geometry Service `json:"geometry"`
}
