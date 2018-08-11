package grafana

//User defines the data struct for a User object
type User struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Email       string `json:"email,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	IsAdmin     bool   `json:"isAdmin,omitempty"`
	LastSeenAge string `json:"lastSeenAtAge,omitempty"`
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
