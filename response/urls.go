package response

type URLs struct {
	URLs ServerURLCollection `json:"urls"`
}
type ServerURL struct {
	HTTPS []string `json:"https"`
}
type ServerURLCollection struct {
	Features  ServerURL `json:"features"`
	Insights  ServerURL `json:"insights"`
	Notebooks ServerURL `json:"notebooks"`
	Tiles     ServerURL `json:"tiles"`
}
