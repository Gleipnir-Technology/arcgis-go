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

type LayerMetadata struct {
	CurrentVersion                         *float64                                `json:"currentVersion,omitempty"`
	ID                                     *int                                    `json:"id,omitempty"`
	Name                                   *string                                 `json:"name,omitempty"`
	Type                                   *string                                 `json:"type,omitempty"`
	ParentLayer                            *int                                    `json:"parentLayer,omitempty"`
	DisplayField                           *string                                 `json:"displayField,omitempty"`
	Description                            *string                                 `json:"description,omitempty"`
	CopyrightText                          *string                                 `json:"copyrightText,omitempty"`
	SubtypeField                           *string                                 `json:"subtypeField,omitempty"`
	DefaultSubtypeCode                     *int                                    `json:"defaultSubtypeCode,omitempty"`
	DefaultVisibility                      *bool                                   `json:"defaultVisibility,omitempty"`
	EditFieldsInfo                         *EditFieldsInfo                         `json:"editFieldsInfo,omitempty"`
	OwnershipBasedAccessControlForFeatures *OwnershipBasedAccessControlForFeatures `json:"ownershipBasedAccessControlForFeatures,omitempty"`
	SyncCanReturnChanges                   *bool                                   `json:"syncCanReturnChanges,omitempty"`
	Relationships                          []Relationship                          `json:"relationships,omitempty"`
	IsDataVersioned                        *bool                                   `json:"isDataVersioned,omitempty"`
	IsDataArchived                         *bool                                   `json:"isDataArchived,omitempty"`
	IsDataBranchVersioned                  *bool                                   `json:"isDataBranchVersioned,omitempty"`
	IsDataReplicaTracked                   *bool                                   `json:"isDataReplicaTracked,omitempty"`
	IsCoGoEnabled                          *bool                                   `json:"isCoGoEnabled,omitempty"`
	SupportsRollbackOnFailureParameter     *bool                                   `json:"supportsRollbackOnFailureParameter,omitempty"`
	DateFieldsTimeReference                *DateFieldsTimeReference                `json:"dateFieldsTimeReference,omitempty"`
	PreferredTimeReference                 *DateFieldsTimeReference                `json:"preferredTimeReference,omitempty"`
	DatesInUnknownTimezone                 *bool                                   `json:"datesInUnknownTimezone,omitempty"`
	ArchivingInfo                          *ArchivingInfo                          `json:"archivingInfo,omitempty"`
	SupportsStatistics                     *bool                                   `json:"supportsStatistics,omitempty"`
	SupportsAdvancedQueries                *bool                                   `json:"supportsAdvancedQueries,omitempty"`
	SupportsCoordinatesQuantization        *bool                                   `json:"supportsCoordinatesQuantization,omitempty"`
	SupportsDatumTransformation            *bool                                   `json:"supportsDatumTransformation,omitempty"`
	GeometryType                           *string                                 `json:"geometryType,omitempty"`
	GeometryProperties                     *GeometryProperties                     `json:"geometryProperties,omitempty"`
	MinScale                               *float64                                `json:"minScale,omitempty"`
	MaxScale                               *float64                                `json:"maxScale,omitempty"`
	EffectiveMinScale                      *float64                                `json:"effectiveMinScale,omitempty"`
	EffectiveMaxScale                      *float64                                `json:"effectiveMaxScale,omitempty"`
	SupportsQuantizationEditMode           *bool                                   `json:"supportsQuantizationEditMode,omitempty"`
	SupportsAppend                         *bool                                   `json:"supportsAppend,omitempty"`
	SupportedAppendFormats                 *string                                 `json:"supportedAppendFormats,omitempty"`
	HasContingentValuesDefinition          *bool                                   `json:"hasContingentValuesDefinition,omitempty"`
	SpatialReference                       *SpatialReference                       `json:"spatialReference,omitempty"`
	AdvancedQueryCapabilities              *AdvancedQueryCapabilities              `json:"advancedQueryCapabilities,omitempty"`
	StandardMaxRecordCountNoGeometry       *int                                    `json:"standardMaxRecordCountNoGeometry,omitempty"`
	SupportsAsyncCalculate                 *bool                                   `json:"supportsAsyncCalculate,omitempty"`
	SupportsFieldDescriptionProperty       *bool                                   `json:"supportsFieldDescriptionProperty,omitempty"`
	AdvancedEditingCapabilities            *AdvancedEditingCapabilities            `json:"advancedEditingCapabilities,omitempty"`
	AdvancedQueryAnalyticCapabilities      *AdvancedQueryAnalyticCapabilities      `json:"advancedQueryAnalyticCapabilities,omitempty"`
	UserTypeExtensions                     []string                                `json:"userTypeExtensions,omitempty"`
	Extent                                 *Extent                                 `json:"extent,omitempty"`
	HeightModelInfo                        *HeightModelInfo                        `json:"heightModelInfo,omitempty"`
	SourceHeightModelInfo                  *HeightModelInfo                        `json:"sourceHeightModelInfo,omitempty"`
	SourceSpatialReference                 *SpatialReference                       `json:"sourceSpatialReference,omitempty"`
	DrawingInfo                            *DrawingInfo                            `json:"drawingInfo,omitempty"`
	HasM                                   *bool                                   `json:"hasM,omitempty"`
	HasZ                                   *bool                                   `json:"hasZ,omitempty"`
	EnableZDefaults                        *bool                                   `json:"enableZDefaults,omitempty"`
	ZDefault                               *float64                                `json:"zDefault,omitempty"`
	AllowGeometryUpdates                   *bool                                   `json:"allowGeometryUpdates,omitempty"`
	TimeInfo                               *TimeInfo                               `json:"timeInfo,omitempty"`
	HasAttachments                         *bool                                   `json:"hasAttachments,omitempty"`
	HTMLPopupType                          *string                                 `json:"htmlPopupType,omitempty"`
	ObjectIDField                          *string                                 `json:"objectIdField,omitempty"`
	GlobalIDField                          *string                                 `json:"globalIdField,omitempty"`
	TypeIDField                            *string                                 `json:"typeIdField,omitempty"`
	Fields                                 []Field                                 `json:"fields,omitempty"`
	GeometryField                          *Field                                  `json:"geometryField,omitempty"`
	Types                                  []Type                                  `json:"types,omitempty"`
	Templates                              []Template                              `json:"templates,omitempty"`
	Subtypes                               []Subtype                               `json:"subtypes,omitempty"`
	MaxRecordCount                         *int                                    `json:"maxRecordCount,omitempty"`
	StandardMaxRecordCount                 *int                                    `json:"standardMaxRecordCount,omitempty"`
	TileMaxRecordCount                     *int                                    `json:"tileMaxRecordCount,omitempty"`
	MaxRecordCountFactor                   *int                                    `json:"maxRecordCountFactor,omitempty"`
	SupportedQueryFormats                  *string                                 `json:"supportedQueryFormats,omitempty"`
	SupportedExportFormats                 *string                                 `json:"supportedExportFormats,omitempty"`
	SupportedSpatialRelationships          []string                                `json:"supportedSpatialRelationships,omitempty"`
	HasMetadata                            *bool                                   `json:"hasMetadata,omitempty"`
	HasStaticData                          *bool                                   `json:"hasStaticData,omitempty"`
	SQLParserVersion                       *string                                 `json:"sqlParserVersion,omitempty"`
	IsUpdatableView                        *bool                                   `json:"isUpdatableView,omitempty"`
	Capabilities                           *string                                 `json:"capabilities,omitempty"`
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

// Field represents a field in the layer or table
type Field struct {
	Name         *string     `json:"name,omitempty"`
	Type         *string     `json:"type,omitempty"`
	Alias        *string     `json:"alias,omitempty"`
	Domain       interface{} `json:"domain,omitempty"`
	Editable     *bool       `json:"editable,omitempty"`
	Nullable     *bool       `json:"nullable,omitempty"`
	Length       *int        `json:"length,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	ModelName    *string     `json:"modelName,omitempty"`
}

// Type represents a feature type
type Type struct {
	ID        *int                   `json:"id,omitempty"`
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
