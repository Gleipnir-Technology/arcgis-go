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

// EditFieldsInfo contains information about edit tracking fields
type EditFieldsInfo struct {
	CreationDateField       *string                  `json:"creationDateField,omitempty"`
	CreatorField            *string                  `json:"creatorField,omitempty"`
	EditDateField           *string                  `json:"editDateField,omitempty"`
	EditorField             *string                  `json:"editorField,omitempty"`
	Realm                   *string                  `json:"realm,omitempty"`
	DateFieldsTimeReference *DateFieldsTimeReference `json:"dateFieldsTimeReference,omitempty"`
}

// DateFieldsTimeReference contains time zone information
type DateFieldsTimeReference struct {
	TimeZone               *string `json:"timeZone,omitempty"`
	TimeZoneIANA           *string `json:"timeZoneIANA,omitempty"`
	RespectsDaylightSaving *bool   `json:"respectsDaylightSaving,omitempty"`
}

// OwnershipBasedAccessControlForFeatures defines ownership-based access control
type OwnershipBasedAccessControlForFeatures struct {
	AllowOthersToUpdate *bool `json:"allowOthersToUpdate,omitempty"`
	AllowOthersToDelete *bool `json:"allowOthersToDelete,omitempty"`
	AllowOthersToQuery  *bool `json:"allowOthersToQuery,omitempty"`
}

// Relationship defines a relationship between tables
type Relationship struct {
	ID                          *int    `json:"id,omitempty"`
	Name                        *string `json:"name,omitempty"`
	RelatedTableID              *int    `json:"relatedTableId,omitempty"`
	Cardinality                 *string `json:"cardinality,omitempty"`
	Role                        *string `json:"role,omitempty"`
	KeyField                    *string `json:"keyField,omitempty"`
	Composite                   *bool   `json:"composite,omitempty"`
	CatalogID                   *string `json:"catalogID,omitempty"`
	RelationshipTableID         *int    `json:"relationshipTableId,omitempty"`
	KeyFieldInRelationshipTable *string `json:"keyFieldInRelationshipTable,omitempty"`
}

// ArchivingInfo contains archiving information
type ArchivingInfo struct {
	SupportsQueryWithHistoricMoment *bool  `json:"supportsQueryWithHistoricMoment,omitempty"`
	StartArchivingMoment            *int64 `json:"startArchivingMoment,omitempty"`
}

// GeometryProperties contains geometry-related properties
type GeometryProperties struct {
	ShapeAreaFieldName   *string `json:"shapeAreaFieldName,omitempty"`
	ShapeLengthFieldName *string `json:"shapeLengthFieldName,omitempty"`
	Units                *string `json:"units,omitempty"`
}

// AdvancedQueryCapabilities defines advanced query capabilities
type AdvancedQueryCapabilities struct {
	SupportsPagination                    *bool       `json:"supportsPagination,omitempty"`
	SupportsTrueCurve                     *bool       `json:"supportsTrueCurve,omitempty"`
	SupportsQueryWithDistance             *bool       `json:"supportsQueryWithDistance,omitempty"`
	SupportsLod                           *bool       `json:"supportsLod,omitempty"`
	SupportsReturningQueryExtent          *bool       `json:"supportsReturningQueryExtent,omitempty"`
	SupportsStatistics                    *bool       `json:"supportsStatistics,omitempty"`
	SupportsHavingClause                  *bool       `json:"supportsHavingClause,omitempty"`
	SupportsOrderBy                       *bool       `json:"supportsOrderBy,omitempty"`
	SupportsDistinct                      *bool       `json:"supportsDistinct,omitempty"`
	SupportsCountDistinct                 *bool       `json:"supportsCountDistinct,omitempty"`
	SupportsPaginationOnAggregatedQueries *bool       `json:"supportsPaginationOnAggregatedQueries,omitempty"`
	SupportsQueryWithResultType           *bool       `json:"supportsQueryWithResultType,omitempty"`
	SupportsReturningGeometryCentroid     *bool       `json:"supportsReturningGeometryCentroid,omitempty"`
	SupportsSQLExpression                 *bool       `json:"supportsSqlExpression,omitempty"`
	SupportsOutFieldsSQLExpression        *bool       `json:"supportsOutFieldsSqlExpression,omitempty"`
	SupportsTopFeaturesQuery              *bool       `json:"supportsTopFeaturesQuery,omitempty"`
	SupportsOrderByOnlyOnLayerFields      *bool       `json:"supportsOrderByOnlyOnLayerFields,omitempty"`
	SupportsQueryWithDatumTransformation  *bool       `json:"supportsQueryWithDatumTransformation,omitempty"`
	SupportsPercentileStatistics          *bool       `json:"supportsPercentileStatistics,omitempty"`
	SupportsQueryAttachments              *bool       `json:"supportsQueryAttachments,omitempty"`
	SupportsQueryAttachmentsWithReturnURL *bool       `json:"supportsQueryAttachmentsWithReturnUrl,omitempty"`
	SupportsQueryAnalytic                 *bool       `json:"supportsQueryAnalytic,omitempty"`
	SupportedMultipatchOptions            []string    `json:"supportedMultipatchOptions,omitempty"`
	SupportsCurrentUserQueries            *bool       `json:"supportsCurrentUserQueries,omitempty"`
	SupportsFullTextSearch                *bool       `json:"supportsFullTextSearch,omitempty"`
	FullTextSearchCapabilities            interface{} `json:"fullTextSearchCapabilities,omitempty"`
	FullTextSearchableFields              []string    `json:"fullTextSearchableFields,omitempty"`
	SupportedCurveTypes                   []string    `json:"supportedCurveTypes,omitempty"`
}

// AdvancedQueryAnalyticCapabilities defines advanced query analytic capabilities
type AdvancedQueryAnalyticCapabilities struct {
	SupportsPercentileAnalytic *bool `json:"supportsPercentileAnalytic,omitempty"`
}

// HeightModelInfo contains height model information
type HeightModelInfo struct {
	HeightModel *string `json:"heightModel,omitempty"`
	VertCRS     *string `json:"vertCRS,omitempty"`
	HeightUnit  *string `json:"heightUnit,omitempty"`
}

// DrawingInfo contains drawing/rendering information
type DrawingInfo struct {
	Renderer     interface{} `json:"renderer,omitempty"`
	Transparency *int        `json:"transparency,omitempty"`
	LabelingInfo interface{} `json:"labelingInfo,omitempty"`
}

// TimeInfo contains temporal information
type TimeInfo struct {
	StartTimeField    *string                  `json:"startTimeField,omitempty"`
	EndTimeField      *string                  `json:"endTimeField,omitempty"`
	TrackIDField      *string                  `json:"trackIdField,omitempty"`
	TimeExtent        []int64                  `json:"timeExtent,omitempty"`
	TimeReference     *DateFieldsTimeReference `json:"timeReference,omitempty"`
	TimeInterval      *int                     `json:"timeInterval,omitempty"`
	TimeIntervalUnits *string                  `json:"timeIntervalUnits,omitempty"`
}

// Type represents a feature type
type Type struct {
	ID        *string                `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Domains   map[string]interface{} `json:"domains,omitempty"`
	Templates []Template             `json:"templates,omitempty"`
}

// Template represents a feature template
type Template struct {
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Prototype   interface{} `json:"prototype,omitempty"`
	DrawingTool *string     `json:"drawingTool,omitempty"`
}

// Subtype represents a feature subtype
type Subtype struct {
	Code          interface{}            `json:"code,omitempty"`
	Name          *string                `json:"name,omitempty"`
	DefaultValues map[string]interface{} `json:"defaultValues,omitempty"`
	Domains       map[string]interface{} `json:"domains,omitempty"`
}
