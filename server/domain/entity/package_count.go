package entity

import "time"

type PackageCount struct {
	Id          int       `json:"-"`
	ChannelId   int       `json:"-"`
	Package     string    `json:"package"`
	BuildString string    `json:"buildString"`
	BuildNumber int       `json:"buildNumber"`
	Version     string    `json:"version"`
	Platform    string    `json:"platform"`
	Count       int       `json:"count"`
	UploadDate  time.Time `json:"uploadDate"`
}
