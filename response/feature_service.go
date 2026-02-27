package response

// https://developers.arcgis.com/rest/services-reference/enterprise/feature-service/#json-response-syntax
/*
  "currentVersion": <currentVersion>,
  "serviceDescription": "<serviceDescription>",
  "hasVersionedData": <true | false>,
  "supportsDisconnectedEditing": <true | false>,
  "supportsDatumTransformation": <true | false>, //Added at 10.8
  "supportsReturnDeleteResults": <true | false>,
  "hasStaticData" : <true | false>,
  "maxRecordCount" : "<maxRecordCount>",
  "supportedQueryFormats": "<supportedQueryFormats>",
  "supportsRelationshipsResource": <true | false>,
  "supportsAppend":  <true | false>,
  "supportedAppendFormats": "<supportedAppendFormats>",
  "supportsTrueCurve": <true | false>,
  "supportedCurveTypes": [<Curve Types>] //Added at 11.5
  "capabilities": "<capabilities>",
  "description": "<description>",
  "copyrightText": "<copyrightText>",
  "userTypeExtensions": [<Extension Types>], //Added at 10.8
  "advancedEditingCapabilities": {<advancedEditingCapabilities>},
  "spatialReference": {<spatialReference>},
  "initialExtent": {<envelope>},
  "fullExtent": {<envelope>},
  "allowGeometryUpdates": <true | false>,
  "units": "<units>",
  "syncEnabled" : <true | false>,
  "supportedExportFormats": "<supported formats>", //Added at 10.9.1
  "returnServiceEditsHaveSR": <true | false>, //Added at 10.7.1
  //Added at 10.7
  "validationSystemLayers": {
    "validationPointErrorlayerId": <validationPointErrorlayerId>,
    "validationLineErrorlayerId": <validationLineErrorlayerId>,
    "validationPolygonErrorlayerId": <validationPolygonErrorlayerId>,
    "validationObjectErrortableId": <validationObjectErrortableId>
  },
  //Added at 10.6.1
  "extractChangesCapabilities": {
    "supportsReturnIdsOnly": <true | false>,
    "supportsReturnExtentOnly": <true | false>,
    "supportsReturnAttachments": <true | false>,
    "supportsLayerQueries": <true | false>,
    "supportsSpatialFilter": <true | false>,
    "supportsReturnFeature": <true | false>,
  },
  "syncCapabilities": {
    "supportsASync": <true | false>,
    "supportsRegisteringExistingData": <true | false>,
    "supportsSyncDirectionControl": <true | false>,
    "supportsPerLayerSync": <true | false>,
    "supportsPerReplicaSync": <true | false>,
    "supportsRollbackOnFailure": <true | false>,
    "supportedSyncDataOptions": <supportedSyncDataOptions>,
    "supportsQueryWithDatumTransformatiom": <true | false>, //Added at 10.8
  },
  "editorTrackingInfo": {
    "enableEditorTracking": <true | false>,
    "enableOwnershipAccessControl": <true | false>,
    "allowOthersToUpdate": <true | false>,
    "allowOthersToDelete": <true | false>
  },
  "documentInfo": {
   "<key1>": "<value1>",
   "<key2>": "<value2>"
   },
  //the feature layers published by this service
  "layers": [
    { "id": <layerId1>, "name": "<layerName1>" },
    { "id": <layerId2>, "name": "<layerName2>" }
  ],
  //the non-spatial tables published by this service
  "tables": [
    { "id": <tableId1>, "name": "<tableName1>" },
    { "id": <tableId2>, "name": "<tableName2>" }
  ],
  "relationships": [
    { "id": <relationshipId1>, "name": "<relationshipName1>" },
    { "id": <relationshipId2>, "name": "<relationshipName2>" }
  ],
  "datumTransformations": [<datumTransformations>] //Added at 10.7.1
  "enableZDefaults": <true | false>,
  "isLocationTrackingService": <true | false>,
  "isLocationTrackingView": <true | false>,
  "isIndoorsService": <true | false>,
  "zDefault": <zDefaultValue>
*/

type FeatureService struct {
	CurrentVersion                              float64                     `json:"currentVersion"`
	ServiceItemId                               string                      `json:"serviceItemId,omitempty"`
	SourceSchemaChangesAllowed                  bool                        `json:"sourceSchemaChangesAllowed,omitempty"`
	HasViews                                    bool                        `json:"hasViews,omitempty"`
	HasSyncEnabledViews                         bool                        `json:"hasSyncEnabledViews,omitempty"`
	ServiceDescription                          string                      `json:"serviceDescription"`
	HasVersionedData                            bool                        `json:"hasVersionedData"`
	SupportsDisconnectedEditing                 bool                        `json:"supportsDisconnectedEditing"`
	SupportsDatumTransformation                 bool                        `json:"supportsDatumTransformation,omitempty"`
	SupportsReturnDeleteResults                 bool                        `json:"supportsReturnDeleteResults"`
	HasStaticData                               bool                        `json:"hasStaticData"`
	HasSharedDomains                            bool                        `json:"hasSharedDomains,omitempty"`
	MaxRecordCount                              uint                        `json:"maxRecordCount,omitempty"`
	SupportedQueryFormats                       string                      `json:"supportedQueryFormats"`
	SupportsRelationshipsResource               bool                        `json:"supportsRelationshipsResource,omitempty"`
	SupportsAppend                              bool                        `json:"supportsAppend"`
	SupportedAppendFormats                      string                      `json:"supportedAppendFormats,omitempty"`
	SupportedExportFormats                      string                      `json:"supportedExportFormats,omitempty"`
	SupportsTrueCurve                           bool                        `json:"supportsTrueCurve"`
	SupportedCurveTypes                         []string                    `json:"supportedCurveTypes,omitempty"`
	SupportsVCSProjection                       bool                        `json:"supportsVCSProjection,omitempty"`
	SupportedConvertFileFormats                 string                      `json:"supportedConvertFileFormats,omitempty"`
	SupportedConvertContentFormats              string                      `json:"supportedConvertContentFormats,omitempty"`
	SupportedFullTextLocales                    []string                    `json:"supportedFullTextLocales,omitempty"`
	SupportsSharedTemplates                     bool                        `json:"supportsSharedTemplates,omitempty"`
	HasSharedTemplates                          bool                        `json:"hasSharedTemplates,omitempty"`
	Capabilities                                string                      `json:"capabilities"`
	Description                                 string                      `json:"description"`
	CopyrightText                               string                      `json:"copyrightText"`
	UserTypeExtensions                          []string                    `json:"userTypeExtensions,omitempty"`
	AdvancedEditingCapabilities                 AdvancedEditingCapabilities `json:"advancedEditingCapabilities"`
	SpatialReference                            SpatialReference            `json:"spatialReference"`
	InitialExtent                               Envelope                    `json:"initialExtent"`
	FullExtent                                  Envelope                    `json:"fullExtent"`
	AllowGeometryUpdates                        bool                        `json:"allowGeometryUpdates"`
	AllowTrueCurvesUpdates                      bool                        `json:"allowTrueCurvesUpdates,omitempty"`
	OnlyAllowTrueCurveUpdatesByTrueCurveClients bool                        `json:"onlyAllowTrueCurveUpdatesByTrueCurveClients,omitempty"`
	Units                                       string                      `json:"units"`
	SyncEnabled                                 bool                        `json:"syncEnabled"`
	ReturnServiceEditsHaveSR                    bool                        `json:"returnServiceEditsHaveSR,omitempty"`
	ValidationSystemLayers                      ValidationSystemLayers      `json:"validationSystemLayers,omitempty"`
	ExtractChangesCapabilities                  ExtractChangesCapabilities  `json:"extractChangesCapabilities,omitempty"`
	SyncCapabilities                            SyncCapabilities            `json:"syncCapabilities,omitempty"`
	EditorTrackingInfo                          EditorTrackingInfo          `json:"editorTrackingInfo,omitempty"`
	ChangeTrackingInfo                          ChangeTrackingInfo          `json:"changeTrackingInfo,omitempty"`
	DocumentInfo                                map[string]string           `json:"documentInfo,omitempty"`
	Layers                                      []Layer                     `json:"layers,omitempty"`
	Tables                                      []TableInfo                 `json:"tables,omitempty"`
	Relationships                               []RelationshipInfo          `json:"relationships,omitempty"`
	DatumTransformations                        []interface{}               `json:"datumTransformations,omitempty"`
	EnableZDefaults                             bool                        `json:"enableZDefaults,omitempty"`
	IsLocationTrackingService                   bool                        `json:"isLocationTrackingService,omitempty"`
	IsLocationTrackingView                      bool                        `json:"isLocationTrackingView,omitempty"`
	IsIndoorsService                            bool                        `json:"isIndoorsService,omitempty"`
	ZDefault                                    float64                     `json:"zDefault,omitempty"`
	XssPreventionInfo                           XssPreventionInfo           `json:"xssPreventionInfo,omitempty"`
	SupportsApplyEditsWithGlobalIds             bool                        `json:"supportsApplyEditsWithGlobalIds,omitempty"`
	SupportsLayerOverrides                      bool                        `json:"supportsLayerOverrides,omitempty"`
	SupportsTilesAndBasicQueriesMode            bool                        `json:"supportsTilesAndBasicQueriesMode,omitempty"`
	SupportsQueryContingentValues               bool                        `json:"supportsQueryContingentValues,omitempty"`
	SupportedContingentValuesFormats            string                      `json:"supportedContingentValuesFormats,omitempty"`
	SupportsContingentValuesJson                int                         `json:"supportsContingentValuesJson,omitempty"`
	Size                                        int64                       `json:"size,omitempty"`
	LayerOverridesEnabled                       bool                        `json:"layerOverridesEnabled,omitempty"`
}

type Envelope struct {
	Xmin             float64          `json:"xmin"`
	Ymin             float64          `json:"ymin"`
	Xmax             float64          `json:"xmax"`
	Ymax             float64          `json:"ymax"`
	SpatialReference SpatialReference `json:"spatialReference"`
}

type ValidationSystemLayers struct {
	ValidationPointErrorLayerId   int `json:"validationPointErrorlayerId"`
	ValidationLineErrorLayerId    int `json:"validationLineErrorlayerId"`
	ValidationPolygonErrorLayerId int `json:"validationPolygonErrorlayerId"`
	ValidationObjectErrorTableId  int `json:"validationObjectErrortableId"`
}

type ExtractChangesCapabilities struct {
	SupportsReturnIdsOnly            bool `json:"supportsReturnIdsOnly"`
	SupportsReturnExtentOnly         bool `json:"supportsReturnExtentOnly"`
	SupportsReturnAttachments        bool `json:"supportsReturnAttachments"`
	SupportsLayerQueries             bool `json:"supportsLayerQueries"`
	SupportsSpatialFilter            bool `json:"supportsSpatialFilter,omitempty"`
	SupportsFeatureReturn            bool `json:"supportsFeatureReturn,omitempty"`
	SupportsReturnFeature            bool `json:"supportsReturnFeature,omitempty"`
	SupportsGeometry                 bool `json:"supportsGeometry,omitempty"`
	SupportsReturnHasGeometryUpdates bool `json:"supportsReturnHasGeometryUpdates,omitempty"`
	SupportsReturnDeletedFeatures    bool `json:"supportsReturnDeletedFeatures,omitempty"`
	SupportsServerGens               bool `json:"supportsServerGens,omitempty"`
	SupportsFieldsToCompare          bool `json:"supportsFieldsToCompare,omitempty"`
}

type SyncCapabilities struct {
	SupportsAsync                        bool `json:"supportsAsync"`
	SupportsRegisteringExistingData      bool `json:"supportsRegisteringExistingData"`
	SupportsSyncDirectionControl         bool `json:"supportsSyncDirectionControl"`
	SupportsPerLayerSync                 bool `json:"supportsPerLayerSync"`
	SupportsPerReplicaSync               bool `json:"supportsPerReplicaSync"`
	SupportsRollbackOnFailure            bool `json:"supportsRollbackOnFailure"`
	SupportedSyncDataOptions             int  `json:"supportedSyncDataOptions"`
	SupportsQueryWithDatumTransformation bool `json:"supportsQueryWithDatumTransformatiom,omitempty"`
	SupportsSyncModelNone                bool `json:"supportsSyncModelNone,omitempty"`
	SupportsAttachmentsSyncDirection     bool `json:"supportsAttachmentsSyncDirection,omitempty"`
	SupportsBiDirectionalSyncForServer   bool `json:"supportsBiDirectionalSyncForServer,omitempty"`
}

type EditorTrackingInfo struct {
	EnableEditorTracking         bool `json:"enableEditorTracking"`
	EnableOwnershipAccessControl bool `json:"enableOwnershipAccessControl"`
	AllowOthersToUpdate          bool `json:"allowOthersToUpdate"`
	AllowOthersToDelete          bool `json:"allowOthersToDelete"`
	AllowOthersToQuery           bool `json:"allowOthersToQuery,omitempty"`
	AllowAnonymousToUpdate       bool `json:"allowAnonymousToUpdate,omitempty"`
	AllowAnonymousToDelete       bool `json:"allowAnonymousToDelete,omitempty"`
	AllowAnonymousToQuery        bool `json:"allowAnonymousToQuery,omitempty"`
}

type ServerGen struct {
	ID           int `json:"id"`
	MinServerGen int `json:"minServerGen"`
	ServerGen    int `json:"serverGen"`
}

type ChangeTrackingInfo struct {
	LastSyncDate    int64       `json:"lastSyncDate"`
	LayerServerGens []ServerGen `json:"layerServerGens"`
}

type TableInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ParentLayerId     int    `json:"parentLayerId"`
	DefaultVisibility bool   `json:"defaultVisibility"`
	SubLayerIds       any    `json:"subLayerIds"`
	MinScale          int    `json:"minScale"`
	MaxScale          int    `json:"maxScale"`
}

type RelationshipInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type XssPreventionInfo struct {
	XssPreventionEnabled bool   `json:"xssPreventionEnabled"`
	XssPreventionRule    string `json:"xssPreventionRule"`
	XssInputRule         string `json:"xssInputRule"`
}
