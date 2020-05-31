package condatypes

type RepoData struct {
	Info struct {
		Platform             string `json:"platform,omitempty"`
		Subdir               string `json:"subdir"`
		DefaultPythonVersion string `json:"default_python_version,omitempty"`
		Arch                 string `json:"arch,omitempty"`
		DefaultNumpyVersion  string `json:"default_numpy_version"`
	} `json:"info"`
	Packages map[string]*PackageDetail `json:"packages"`
}

type PackageDetail struct {
	Arch          string   `json:"arch,omitempty"`
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
