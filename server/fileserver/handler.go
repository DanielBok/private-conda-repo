package fileserver

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda/condatypes"
)

var nameRegex = regexp.MustCompile(`([\w\-_]+)-([\w.]+)-(\w+)_(\d+).tar.bz2`)

func FileHandler(mountFolder string) http.HandlerFunc {
	root := http.Dir(mountFolder)
	fs := http.FileServer(root)

	log.WithField("Repository mount folder used to serve packages", root).Info()

	return func(w http.ResponseWriter, r *http.Request) {
		components := strings.Split(strings.Trim(r.RequestURI, "/"), "/")
		if len(components) != 3 {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		channel := components[0]
		platform := components[1]
		file := components[2]
		if _, err := condatypes.MapPlatform(platform); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch {
		case file == "current_repodata.json":
			log.Infof("requesting repodata from '%s/%s'", channel, platform)
		case strings.HasSuffix(file, ".tar.bz2"):
			p, err := getPackageDetail(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if _, err := db.IncreasePackageCount(channel, p.Name, platform, p.Version); err != nil {
				http.Error(w, errors.Wrap(err, "could not increment package count").Error(), http.StatusInternalServerError)
				return
			}
			log.Infof("serving '%s' to remote '%s'", file, r.RemoteAddr)

		default:
			http.Error(w, "request is not made by a conda agent", http.StatusBadRequest)
			return
		}
		fs.ServeHTTP(w, r)
	}
}

func getPackageDetail(file string) (*condatypes.Package, error) {
	m := nameRegex.FindStringSubmatch(file)
	if len(m) != 5 {
		return nil, errors.Errorf("could not parse package name from '%s'", file)
	}

	bNo, err := strconv.Atoi(m[4])
	if err != nil {
		return nil, errors.Errorf("could not parse build number from '%s'", file)
	}

	return &condatypes.Package{
		Name:        m[1],
		Version:     m[2],
		BuildString: m[3],
		BuildNumber: bNo,
	}, nil
}
