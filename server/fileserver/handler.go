package fileserver

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda/condatypes"
)

var nameRegex = regexp.MustCompile(`([\w\-_]+)-[\w.]+-\w+_\d+.tar.bz2`)

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
			name, err := getPackageName(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if _, err := db.IncreasePackageCount(channel, name, platform); err != nil {
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

func getPackageName(file string) (string, error) {
	matches := nameRegex.FindStringSubmatch(file)
	if len(matches) != 2 {
		return "", errors.Errorf("could not detect package name from '%s'", file)
	}
	return matches[1], nil
}
