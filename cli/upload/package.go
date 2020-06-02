package upload

import "fmt"

type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	BuildString string `json:"buildString"`
	BuildNumber int    `json:"buildNumber"`
	Platform    string `json:"platform"`
}

// Returns the package's full filename (i.e. perfana-0.0.6-py_0.tar.bz2)
func (p *Package) Filename() string {
	return fmt.Sprintf("%s-%s-%s_%d.tar.bz2", p.Name, p.Version, p.BuildString, p.BuildNumber)
}
