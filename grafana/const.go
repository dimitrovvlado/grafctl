package grafana

const (
	// AuthorizationHeader value
	AuthorizationHeader = "Authorization"

	//AuthorizationTypeBasic value
	AuthorizationTypeBasic = "Basic "

	// UsersEndpoint is the API endoint for users.
	UsersEndpoint = "/api/users"

	// OrgsUsersEndpoint is the API endoint for organizations.
	OrgsUsersEndpoint = "/api/org/users"

	// OrgsEndpoint is the API endoint for organizations.
	OrgsEndpoint = "/api/org"

	// DatasourcesEndpoint is the API endoint for datasources.
	DatasourcesEndpoint = "/api/datasources"

	// TeamsSearchEndpoint is the API endpoint for searching Teams
	TeamsSearchEndpoint = "/api/teams/search"

	// TeamsEndpoint is the API endpoint for searching Teams
	TeamsEndpoint = "/api/teams"

	//SearchEndpoint is the API endpoint for searching
	SearchEndpoint = "/api/search"

	//DashboardsImportEndpoint is the API endpoint for create/update of dashboards
	DashboardsImportEndpoint = "/api/dashboards/import"

	//DashboardsUIDEndpoint is the API endpoint for create/update of dashboards
	DashboardsUIDEndpoint = "/api/dashboards/uid"

	//FoldersEndpoint is the API endpoint for folders
	FoldersEndpoint = "/api/folders"
)
