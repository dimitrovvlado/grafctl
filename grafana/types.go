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

// "id": 1,
// "name": "",
// "login": "usagemeter",
// "email": "usagemeter@localhost",
// "avatarUrl": "/avatar/78e9bc2c2b0992a40a058aafe115b065",
// "isAdmin": true,
// "lastSeenAt": "2018-08-08T09:24:24Z",
// "lastSeenAtAge": "< 1m"
