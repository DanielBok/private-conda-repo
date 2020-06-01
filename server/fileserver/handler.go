package fileserver

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/entity"
	"private-conda-repo/domain/enum"
)

var nameRegex = regexp.MustCompile(`([\w\-_]+)-([\w.]+)-(\w+)_(\d+).tar.bz2`)

type FileHandler struct {
	DB interfaces.DataAccessLayer
}

func (h *FileHandler) Server(mountFolder string) http.HandlerFunc {
	root := http.Dir(mountFolder)
	fs := http.FileServer(root)

	return func(w http.ResponseWriter, r *http.Request) {
		components := strings.Split(strings.Trim(r.RequestURI, "/"), "/")
		if len(components) != 3 {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		channel := components[0]
		platform := components[1]
		file := components[2]
		if _, err := enum.MapPlatform(platform); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch {
		case file == "current_repodata.json" || file == "repodata.json":
			log.Infof("requesting repodata from '%s/%s'", channel, platform)
		case strings.HasSuffix(file, ".tar.bz2"):
			err := h.updateCount(channel, file, platform)
			if err != nil {
				http.Error(w, "could not update package count", http.StatusBadRequest)
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

func (h *FileHandler) updateCount(channel, file, platform string) error {
	chn, err := h.DB.GetChannel(channel)
	if err != nil {
		return err
	}

	m := nameRegex.FindStringSubmatch(file)
	if len(m) != 5 {
		return errors.Errorf("could not parse package name from '%s'", file)
	}

	buildNo, err := strconv.Atoi(m[4])
	if err != nil {
		return errors.Errorf("could not parse build number from '%s'", file)
	}

	_, err = h.DB.IncreasePackageCount(&entity.PackageCount{
		ChannelId:   chn.Id,
		Package:     m[1],
		Version:     m[2],
		BuildString: m[3],
		BuildNumber: buildNo,
		Platform:    platform,
	})
	return err
}
