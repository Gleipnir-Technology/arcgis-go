package arcgis

type Basemap struct {
	BasemapLayers []LayerResource `json:"baseMapLayers"`
	Title         string          `json:"title"`
}

type Extent struct {
	XMin             float64 `json:"xmin"`
	YMin             float64 `json:"ymin"`
	XMax             float64 `json:"xmax"`
	YMax             float64 `json:"ymax"`
	SpatialReference SpatialReference
}

type LayerResource struct {
	ID        string `json:"id"`
	LayerType string `json:"layerType"`
	//ResourceInfo ResourceInfo `json:"resourceInfo"`
	URL        string `json:"url"`
	Visibility bool   `json:"visibility"`
	Opacity    int    `json:"opacity"`
	Title      string `json:"title"`
}

type ResourceInfo struct {
	CurrentVersion        float32 `json:"currentVersion"`
	MapName               string  `json:"mapName"`
	SupportsDynamicLayers bool    `json:"supportsDynamicLayers"`
	//Layers []LayerResource
	//Tables TableResource
	//SpatialReference SpatialReference `json:"spatialReference"`
	SingleFusedMapCache bool `json:"singleFusedMapCache"`
	//TileInfo TileInfo `json:"tileInfo"`
	//InitialExtent Extent `json:"initialExtent"`
	//FullExtent Extent `json:"fullExtent"`
}

type PortalsResponse struct {
	TwoDSketchStylesGroupQuery     string   `json:"2DSketchStylesGroupQuery"`
	TwoDStylesGroupQuery           string   `json:"2DStylesGroupQuery"`
	ThreeDBasemapGalleryGroupQuery string   `json:"3DBasemapGalleryGroupQuery"`
	Access                         string   `json:"access"`
	AllSSL                         bool     `json:"allSSL"`
	AllowedRedirectURIs            []string `json:"allowedRedirectUris"`
	AnalysisLayersGroupQuery       string   `json:"analysisLayersGroupQuery"`
	AuthorizedCrossOriginDomains   []string `json:"authorizedCrossOriginDomains"`
	AvailableCredits               float32  `json:"availableCredits"`
	BackgroundImage                string   `json:"backgroundImage"`
	BasemapGalleryGroupQuery       string   `json:"basemapGalleryGroupQuery"`
	CanListApps                    bool     `json:"canListApps"`
	CanListConsultingServices      bool     `json:"canListConsultingServices"`
	CanListData                    bool     `json:"canListData"`
	CanListPreProvisionedItems     bool     `json:"canListPreProvisionedItems"`
	CanListSolutions               bool     `json:"canListSolutions"`
	CanProvisionDirectPurchase     bool     `json:"canProvisionDirectPurchase"`
	CanSearchPublic                bool     `json:"canSearchPublic"`
	CanSetCustomBuyLink            bool     `json:"canSetCustomBuyLink"`
	CanSetQuestionnaire            bool     `json:"canSetQuestionnaire"`
	CanShareBingPublic             bool     `json:"canShareBingPublic"`
	CanSharePublic                 bool     `json:"canSharePublic"`
	CanSignInArcGIS                bool     `json:"canSignInArcGIS"`
	CanSignInIDP                   bool     `json:"canSignInIDP"`
	CanSignInOIDC                  bool     `json:"canSignInOIDC"`
	CanSignInSocial                bool     `json:"canSignInSocial"`
	CDNUrl                         string   `json:"cdnUrl"`
	ColorSetsGroupQuery            string   `json:"colorSetsGroupQuery"`
	CommentsEnabled                bool     `json:"commentsEnabled"`
	ContentCategorySetsGroupQuery  string   `json:"contentCategorySetsGroupQuery"`
	Created                        int64    `json:"created"`
	Culture                        string   `json:"culture"`
	CultureFormat                  string   `json:"cultureFormat"`
	CustomBaseUrl                  string   `json:"customBaseUrl"`
	DatabaseQuota                  int      `json:"databaseQuota"`
	DatabaseUsage                  int      `json:"databaseUsage"`
	Default3DBasemapQuery          string   `json:"default3DBasemapQuery"`
	DefaultBasemap                 Basemap  `json:"defaultBasemap"`
	DefaultDevBasemap              Basemap  `json:"defaultDevBasemap"`
	DefaultExtent                  Extent   `json:"defaultExtent"`
	DefaultUserCreditAssignment    int      `json:"defaultUserCreditAssignment"`
	DefaultVectorBasemap           Basemap  `json:"defaultVectorBasemap"`
	Description                    string
	Dev3DBasemapGalleryGroupQuery  string
	DevBasemapGalleryGroupQuery    string
	EueiEnabled                    bool
	FeaturedGroups                 []Group
	FeaturedGroupsId               string `json:"featuredGroupsId"`
	FeaturedItemsGroupQuery        string `json:"featuredItemsGroupQuery"`
	FontManifestUrl                string `json:"fontManifestUrl"`
	G3DTilesGalleryGroupQuery      string `json:"g3DTilesGalleryGroupQuery"`
	G3dTilesEnabled                bool   `json:"g3dTilesEnabled"`
	GalleryTemplatesGroupQuery     string `json:"galleryTemplatesGroupQuery"`
	HasCategorySchema              bool   `json:"hasCategorySchema"`
	HasMemberCategorySchema        bool   `json:"hasMemberCategorySchema"`
	HelpBase                       string `json:"helpBase"`
	//HelperServices []HelperService
	HomePageFeaturedContent      string   `json:"homePageFeaturedContent"`
	HomePageFeaturedContentCount int      `json:"homePageFeaturedContentCount"`
	ID                           string   `json:"id"`
	InactivityTimeout            int      `json:"inactivityTimeout"`
	IsPortal                     bool     `json:"isPortal"`
	IsVerified                   bool     `json:"isVerified"`
	LayerTemplatesGroupQuery     string   `json:"layerTemplatesGroupQuery"`
	LivingAtlasGroupQuery        string   `json:"livingAtlasGroupQuery"`
	MaxTokenExpirationMinutes    int      `json:"maxTokenExpirationMinutes"`
	MetadataEditable             bool     `json:"metadataEditable"`
	MetadataFormats              []string `json:"metadataFormats"`
	Modified                     int64    `json:"modified"`
	Name                         string   `json:"name"`
	NotificationsEnabled         bool     `json:"notificationsEnabled"`
	PlatformSSO                  bool     `json:"platformSSO"`
	PortalHostname               string   `json:"portalHostname"`
	PortalMode                   string   `json:"portalMode"`
	PortalName                   string   `json:"portalName"`
	//PortalProperties []PortalProperty `json:"portalProperties"`
	PortalThumbnail                   *string `json:"portalThumbnail"`
	RasterFunctionTemplatesGroupQuery string  `json:"rasterFunctionTemplatesGroupQuery"`
	Region                            string  `json:"region"`
	//RotatorPanels []RotatorPanel `json:"rotatorPanels"`
	ShowHomePageDescription        bool   `json:"showHomePageDescription"`
	StaticImagesUrl                string `json:"staticImagesUrl"`
	StorageQuota                   int    `json:"storageQuota"`
	StorageUsage                   int    `json:"storageUsage"`
	StylesGroupQuery               string `json:"stylesGroupQuery"`
	Supports3DTilesServices        bool   `json:"supports3DTilesServices"`
	SupportsHostedServices         bool   `json:"supportsHostedServices"`
	SymbolSetsGroupQuery           string `json:"symbolSetsGroupQuery"`
	TemplatesGroupQuery            string `json:"templatesGroupQuery"`
	Thumbnail                      string `json:"thumbnail"`
	Units                          string `json:"units"`
	UpdateUserProfileDisabled      bool   `json:"updateUserProfileDisabled"`
	UrlKey                         string `json:"urlKey"`
	Use3dBasemaps                  bool   `json:"use3dBasemaps"`
	UseDefault3dBasemap            bool   `json:"useDefault3dBasemap"`
	UseStandardizedQuery           bool   `json:"useStandardizedQuery"`
	UseVectorBasemaps              bool   `json:"useVectorBasemaps"`
	VectorBasemapGalleryGroupQuery string `json:"vectorBasemapGalleryGroupQuery"`
	//subscriptionInfo SubscriptionInfo `json:"subscriptionInfo"`
	RecycleBinSupported bool     `json:"recycleBinSupported"`
	RecycleBinEnabled   bool     `json:"recycleBinEnabled"`
	IpCntryCode         string   `json:"ipCntryCode"`
	HttpPort            int      `json:"httpPort"`
	HttpsPort           int      `json:"httpsPort"`
	SupportsOAuth       bool     `json:"supportsOAuth"`
	IsReadOnly          bool     `json:"isReadOnly"`
	CurrentVersion      string   `json:"currentVersion"`
	MfaEnabled          bool     `json:"mfaEnabled"`
	Contacts            []string `json:"contacts"`
	User                UserInfo `json:"user"`
	AppInfo             AppInfo  `json:"appInfo"`
}

type SpatialReference struct {
	LatestWKID int `json:"wkid"`
	WKID       int `json:"latestWkid"`
}

type AppInfo struct {
	AppId      string   `json:"appId"`
	ItemId     string   `json:"itemId"`
	AppOwner   string   `json:"appOwner"`
	OrgId      string   `json:"orgId"`
	AppTitle   string   `json:"appTitle"`
	Privileges []string `json:"privileges"`
}

type Group struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
}

type UserInfo struct {
	Username             string   `json:"username"`
	Udn                  *string  `json:"udn"`
	Id                   string   `json:"id"`
	FullName             string   `json:"fullName"`
	Categories           []string `json:"categories"`
	EmailStatus          string   `json:"emailStatus"`
	EmailStatusDate      int64    `json:"emailStatusDate"`
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	PreferredView        *string  `json:"preferredView"`
	Description          *string  `json:"description"`
	Email                string   `json:"email"`
	UserType             string   `json:"userType"`
	IdpUsername          *string  `json:"idpUsername"`
	FavGroupId           string   `json:"favGroupId"`
	LastLogin            int64    `json:"lastLogin"`
	MfaEnabled           bool     `json:"mfaEnabled"`
	MfaEnforcementExempt bool     `json:"mfaEnforcementExempt"`
	StorageUsage         int64    `json:"storageUsage"`
	StorageQuota         int64    `json:"storageQuota"`
	OrgId                string   `json:"orgId"`
	Role                 string   `json:"role"`
	Privileges           []string `json:"privileges"`
	RoleId               string   `json:"roleId"`
	Level                string   `json:"level"`
	UserLicenseTypeId    string   `json:"userLicenseTypeId"`
	Disabled             bool     `json:"disabled"`
	Tags                 []string `json:"tags"`
	Culture              string   `json:"culture"`
	CultureFormat        string   `json:"cultureFormat"`
	Region               string   `json:"region"`
	Units                string   `json:"units"`
	Thumbnail            *string  `json:"thumbnail"`
	Access               string   `json:"access"`
	Created              int64    `json:"created"`
	Modified             int64    `json:"modified"`
	Provider             string   `json:"provider"`
}

type SearchResponse struct {
	Total     int            `json:"total"`
	Start     int            `json:"start"`
	Num       int            `json:"num"`
	NextStart int            `json:"nextStart"`
	Results   []SearchResult `json:"results"`
}
type SearchResult struct {
	ID string `json:"id"`
	//Owner string `json:"owner"`
	//OrgID string `json:"orgId"`
	//Created int64 `json:"created"`
	//IsOrgItem bool `json:"isOrgItem"`
	//Modified int64 `json:"modified"`
	//Guid *string `json:"guid"`
	Name string `json:"name"`
	//Title string `json:"title"`
	//Type string `json:"type"`
	//TypeKeywords []string `json:"typeKeywords"`
	//Description string `json:"description"`
	//Tags []string `json:"tags"`
	//Snippet string `json:"snippet"`
	//Thumbnail string `json:"thumbnail"`
	//Documentation *string `json:"documentation"`
	//Extent [][]float32 `json:"extent"`
	//Categories []string `json:"categories"`
	//SpatialReference string `json:"spatialReference"`
	//AccessInformation string `json:"accessInformation"`
	//Classification *string `json:"classification"`
	//LicenseInfo string `json:"licenseInfo"`
	//Culture string `json:"culture"`
	//Properties *string `json:"properties"`
	//AdvancedSettings *string `json:"advancedSettings"`
	URL string `json:"url"`
	//ProxyFilter *string `json:"proxyFilter"`
	//Access string `json:"access"`
	//Size int `json:"size"`
	//SubInfo int `json:"subInfo"`
	//AppCategories []string `json:"appCategories"`
	//Industries []string `json:"industries"`
	//Languages []string `json:"languages"`
	//LargeThumbnail *string `json:"largeThumbnail"`
	//Banner *string `json:"banner"`
	//Screenshots []string `json:"screenshots"`
	//Listed bool `json:"listed"`
	//NumComments int `json:"numComments"`
	//NumRatings int `json:"numRatings"`
	//AvgRating int `json:"avgRating"`
	//NumViews int `json:"numViews"`
	//ScoreCompleteness int `json:"scoreCompleteness"`
	//GroupDesignations *string `json:"groupDesignations"`
	//ApiToken1ExpirationDate int `json:"apiToken1ExpirationDate"`
	//ApiToken2ExpirationDate int `json:"apiToken2ExpirationDate"`
	//LastViewed int64 `json:"lastViewed"`
}
