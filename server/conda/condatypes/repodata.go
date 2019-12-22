package condatypes

import "regexp"

type RepoData struct {
	Packages map[string]ChannelPackageDetail `json:"packages"`
}

type ChannelPackageDetail struct {
	Build         string   `json:"build"`
	BuildNumber   int      `json:"build_number"`
	Depends       []string `json:"depends"`
	License       string   `json:"license"`
	LicenseFamily string   `json:"license_family,omitempty"`
	MD5           string   `json:"md5"`
	Name          string   `json:"name"`
	NoArch        string   `json:"noarch,omitempty"`
	SHA256        string   `json:"sha256"`
	Size          int      `json:"size"`
	Subdir        string   `json:"subdir"`
	Timestamp     uint64   `json:"timestamp"`
	Version       string   `json:"version"`
}

func (c *ChannelPackageDetail) ToPackage() *Package {
	p := Package{
		Name:        c.Name,
		Version:     c.Version,
		BuildNumber: c.BuildNumber,
		Platform:    c.Subdir,
	}

	r := regexp.MustCompile(`.*_(\d+)$`)
	matches := r.FindStringSubmatch(c.Build)
	if len(matches) == 2 {
		p.BuildString = matches[1]
	} else {
		p.BuildString = matches[0]
	}

	return &p
}
