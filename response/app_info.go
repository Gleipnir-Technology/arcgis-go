package response

type AppInfo struct {
	AppId      string   `json:"appId"`
	ItemId     string   `json:"itemId"`
	AppOwner   string   `json:"appOwner"`
	OrgId      string   `json:"orgId"`
	AppTitle   string   `json:"appTitle"`
	Privileges []string `json:"privileges"`
}
