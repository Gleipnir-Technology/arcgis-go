package response

type QueryResult struct {
	Features          []Feature        `json:"features"`
	Fields            []Field          `json:"fields"`
	GeometryType      string           `json:"geometryType"`
	GlobalIDFieldName string           `json:"globalIdFieldName"`
	ObjectIDFieldName string           `json:"objectIdFieldName"`
	SpatialReference  SpatialReference `json:"spatialReference"`
	UniqueIDField     UniqueIdField    `json:"uniqueIdField"`
}
