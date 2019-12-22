package application

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func ListPackagesByUser(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	chn, err := repo.GetChannel(user)
	if err != nil {
		http.Error(w, errors.Wrapf(err, "could not find user/channel with name '%s'", user).Error(), http.StatusBadRequest)
		return
	}

	meta, err := chn.GetMetaInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, meta.NormalizedPackagesOutput())
}

func ListPackageDetails(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	pkg := chi.URLParam(r, "pkg")

	repo, err := repo.GetChannel(user)
	if err != nil {
		http.Error(w, errors.Wrapf(err, "could not find user/channel with name '%s'", user).Error(), http.StatusBadRequest)
		return
	}

	details, err := repo.GetPackageDetails(pkg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, details)
}

func UploadPackage(w http.ResponseWriter, r *http.Request) {

}

func RemovePackage(w http.ResponseWriter, r *http.Request) {

}

func RemoveAllPackages(w http.ResponseWriter, r *http.Request) {

}
