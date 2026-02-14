package response

type Permission struct {
	ChildURL   string          `json:"childURL"`
	Operation  string          `json:"operation"`
	Permission PermissionEntry `json:"permission"`
	Principal  string          `json:"principal"`
}
type PermissionEntry struct {
	Constraint string `json:"constraint"`
	IsAllowed  bool   `json:"isAllowed"`
}

type PermissionListResponse struct {
	Permissions []Permission `json:"permissions"`
}
type PermissionSlice = []Permission
