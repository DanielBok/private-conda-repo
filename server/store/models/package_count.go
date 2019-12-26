package models

type PackageCount struct {
	Id          int    `json:"-"`
	Channel     string `json:"channel"` // this will be the channel name as well
	Package     string `json:"package"`
	BuildString string `json:"build_string"`
	BuildNumber int    `json:"build_number"`
	Version     string `json:"version"`
	Platform    string `json:"platform"`
	Count       int    `json:"count"`
}
