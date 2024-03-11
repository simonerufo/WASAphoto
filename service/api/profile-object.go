package api

type Profile struct {
	User User `json:"user"`
	//Photos    []Photo `json:"photos"`
	Followers int `json:"followers"`
	Following int `json:"following"`
}
