package entity

import "time"

type PackageCount struct {
	Id          int       `json:"-"`
	ChannelId   int       `json:"-"`
	Package     string    `json:"package"`
	BuildString string    `json:"build_string"`
	BuildNumber int       `json:"build_number"`
	Version     string    `json:"version"`
	Platform    string    `json:"platform"`
	Count       int       `json:"count"`
	UploadDate  time.Time `json:"upload_date"`
}
