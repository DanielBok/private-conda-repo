package application

import "net/http"

type ManagerImageInfo struct {
	Image   string `json:"image"`
	Version int    `json:"version"`
}

func GetImageManagerVersion(w http.ResponseWriter, _ *http.Request) {
	version, err := mgr.CheckCurrentVersion()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, ManagerImageInfo{
		Image:   mgr.Image,
		Version: version,
	})
}

func UpdateImageManagerVersion(w http.ResponseWriter, _ *http.Request) {
	version, err := mgr.UpdateImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, ManagerImageInfo{
		Image:   mgr.Image,
		Version: version,
	})
}
