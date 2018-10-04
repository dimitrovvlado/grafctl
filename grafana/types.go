package grafana

//User defines the data struct for a User object
type User struct {
	ID          int    `json:"id,omitempty"`
	UserID      int    `json:"userId,omitempty"`
	OrgID       int    `json:"orgId,omitempty"`
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Email       string `json:"email,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	IsAdmin     bool   `json:"isAdmin,omitempty"`
	LastSeenAge string `json:"lastSeenAtAge,omitempty"`
}

//ListUserOptions options for querying users
type ListUserOptions struct {
	CurrentOrg bool
}

//Org defines the data struct for an Organization object
type Org struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Address Address `json:"address,omitempty"`
}

//Address defines the data struct for an Address object
type Address struct {
	Address1 string `json:"address1,omitempty"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city,omitempty"`
	ZipCode  string `json:"zipCode,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
}

//Datasource defines the data struct for an Datasource object
type Datasource struct {
	ID                int    `json:"id,omitempty"`
	OrgID             int    `json:"orgId,omitempty"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type,omitempty"`
	TypeLogoURL       string `json:"typeLogoUrl,omitempty"`
	Access            string `json:"access,omitempty"`
	URL               string `json:"url,omitempty"`
	Password          string `json:"password,omitempty"`
	User              string `json:"user,omitempty"`
	Database          string `json:"database,omitempty"`
	BasicAuth         bool   `json:"basicAuth,omitempty"`
	IsDefault         bool   `json:"isDefault,omitempty"`
	IsDeReadOnlyfault bool   `json:"readOnly,omitempty"`
}

//SearchTeamsOptions options for querying teams
type SearchTeamsOptions struct {
	Query string
}

//TeamPage defines a search page for Teams
type TeamPage struct {
	TotalCount int    `json:"totalCount,omitempty"`
	Teams      []Team `json:"teams,omitempty"`
	Page       int    `json:"page,omitempty"`
	PerPage    int    `json:"perPage,omitempty"`
}

//Team defines the data struct for a Team
type Team struct {
	ID          int    `json:"id,omitempty"`
	OrgID       int    `json:"orgId,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	MemberCount int    `json:"memberCount,omitempty"`
}

//Dashboard defines the data struct for a Dashboard search result
type Dashboard struct {
	ID        int      `json:"id,omitempty"`
	UID       string   `json:"uid,omitempty"`
	Title     string   `json:"title,omitempty"`
	URI       string   `json:"uri,omitempty"`
	URL       string   `json:"url,omitempty"`
	Type      string   `json:"type,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	IsStarred bool     `json:"isStarred,omitempty"`
}

//DashboardRequest defines the data struct for create/update of a Dashbord
type DashboardRequest struct {
	Dashboard interface{} `json:"dashboard,omitempty"`
	Overwrite *bool       `json:"overwrite,omitempty"`
	FolderID  int         `json:"folderId,omitempty"`
	Message   string      `json:"message,omitempty"`
}

//DashboardResponse defines the data struct for a response after successful Dashboard create
type DashboardResponse struct {
	PluginID         string `json:"pluginId,omitempty"`
	Title            string `json:"title,omitempty"`
	Imported         bool   `json:"imported,omitempty"`
	ImportedURI      string `json:"importedUri,omitempty"`
	ImportedURL      string `json:"importedUrl,omitempty"`
	Slug             string `json:"slug,omitempty"`
	DashboardID      int    `json:"dashboardId,omitempty"`
	ImportedRevision int    `json:"importedRevision,omitempty"`
	Revision         int    `json:"revision,omitempty"`
	Description      string `json:"description,omitempty"`
	Path             string `json:"path,omitempty"`
	Removed          bool   `json:"removed,omitempty"`
}

//DashboardExport struct
type DashboardExport struct {
	Meta      DashboardMetadata `json:"meta,omitempty"`
	Dashboard interface{}       `json:"dashboard,omitempty"`
}

//DashboardMetadata defines data struct for the metadata when exporting a dashboard
type DashboardMetadata struct {
	Type        string `json:"type,omitempty"`
	CanSave     bool   `json:"canSave,omitempty"`
	CanEdit     bool   `json:"canEdit,omitempty"`
	CanAdmin    bool   `json:"canAdmin,omitempty"`
	CanStar     bool   `json:"canStar,omitempty"`
	Slug        string `json:"slug,omitempty"`
	URL         string `json:"url,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty"`
	CreatedBy   string `json:"createdBy,omitempty"`
	Version     int    `json:"version,omitempty"`
	HasACL      bool   `json:"hasAcl,omitempty"`
	FolderID    int    `json:"folderId,omitempty"`
	FolderTitle string `json:"folderTitle,omitempty"`
}

//Folder defines data struct for folders.
type Folder struct {
	ID        int    `json:"id,omitempty"`
	UID       string `json:"uid,omitempty"`
	Title     string `json:"title,omitempty"`
	URL       string `json:"url,omitempty"`
	HasACL    bool   `json:"hasAcl,omitempty"`
	CanSave   bool   `json:"canSave,omitempty"`
	CanEdit   bool   `json:"canEdit,omitempty"`
	CanAdmin  bool   `json:"canAdmin,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	Version   int    `json:"version,omitempty"`
}
