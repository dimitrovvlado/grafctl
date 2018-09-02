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
