package api

type Profile struct {
	User User `json:"user"`
	Followers int `json:"followers"`
	Following int `json:"following"`
}
