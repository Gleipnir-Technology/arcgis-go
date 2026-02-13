package response

type UserInfo struct {
	Username             string   `json:"username"`
	Udn                  *string  `json:"udn"`
	ID                   string   `json:"id"`
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
	OrgID                string   `json:"orgId"`
	Role                 string   `json:"role"`
	Privileges           []string `json:"privileges"`
	RoleId               string   `json:"roleId"`
	Level                string   `json:"level"`
	UserLicenseTypeID    string   `json:"userLicenseTypeId"`
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
