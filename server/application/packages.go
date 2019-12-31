package application

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"private-conda-repo/conda/condatypes"
	"private-conda-repo/store/models"
)

type ChannelDetails struct {
	Channel  string              `json:"channel"`
	Password string              `json:"password"`
	Package  *condatypes.Package `json:"package"`
}

type ChannelPackageDetails struct {
	Channel string                               `json:"channel"`
	Package string                               `json:"package"`
	Details []*models.PackageCount               `json:"details"`
	Latest  *condatypes.ChannelMetaPackageOutput `json:"latest"`
}

func (c *ChannelDetails) Validate() error {
	c.Channel = strings.TrimSpace(c.Channel)
	if c.Channel == "" {
		return errors.New("channel name cannot be empty or whitespaces")
	}
	return nil
}

func ListAllPackages(w http.ResponseWriter, r *http.Request) {
	channels, err := repo.ListAllChannels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var output []*condatypes.ChannelMetaPackageOutput
	for _, c := range channels {
		meta, err := c.GetMetaInfo()
		if err != nil {
			http.Error(w, "error getting meta info", http.StatusInternalServerError)
			return
		}

		for _, m := range meta.NormalizedPackagesOutput(c.Name()) {
			output = append(output, m)
		}
	}

	toJson(w, output)
}

func ListPackagesByUser(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	chn, err := repo.GetChannel(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not find user/channel with name '%s'", user), http.StatusBadRequest)
		return
	}

	meta, err := chn.GetMetaInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, meta.NormalizedPackagesOutput(user))
}

func ListPackageDetails(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	pkg := chi.URLParam(r, "pkg")

	details, err := db.GetPackageCounts(user, pkg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if len(details) == 0 {
		http.Error(w, "package does not exist in channel", http.StatusNotFound)
		return
	}

	chn, err := repo.GetChannel(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	meta, err := chn.GetMetaInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get latest package
	var latest *condatypes.ChannelMetaPackageOutput
	for _, m := range meta.NormalizedPackagesOutput(user) {
		if m.Name != pkg {
			continue
		}
		if latest == nil {
			latest = m
		} else if (latest.Version == nil || m.Version == nil) && latest.Timestamp < m.Timestamp {
			latest = m
		} else if *latest.Version < *m.Version {
			latest = m
		}
	}

	toJson(w, ChannelPackageDetails{user, pkg, details, latest})
}

func UploadPackage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 24)
	if err != nil {
		http.Error(w, "could not parse form data", http.StatusInternalServerError)
		return
	}

	// read and validate inputs
	channel := r.FormValue("channel")
	if strings.TrimSpace(channel) == "" {
		http.Error(w, "channel name cannot be empty", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, errors.Wrap(err, "could not parse uploaded file. Please ensure that you have "+
			"uploaded a valid file with 'file' as the form key").Error(), http.StatusBadRequest)
		return
	}
	defer func() { _ = file.Close() }()

	// retrieve package meta data
	pkg, err := dcp.RetrieveMetadata(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer pkg.Close()

	if err := pkg.Package.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check channel and password
	if status, err := validateCredentials(channel, r.FormValue("password")); err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	// get channel and upload file
	chn, err := repo.GetChannel(channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := pkg.Open()
	if err != nil {
		http.Error(w, errors.Wrap(err, "could not open temporary package archive").Error(), http.StatusInternalServerError)
		return
	}

	p, err := chn.AddPackage(f, pkg.Package)
	if err != nil {
		http.Error(w, errors.Wrap(err, "upload failed").Error(), http.StatusInternalServerError)
		return
	}

	if _, err := db.CreatePackageCount(p.ToPackageCount(channel)); err != nil {
		http.Error(w, errors.Wrap(err, "could not create package count record").Error(), http.StatusInternalServerError)
		return
	}

	// return outcome
	toJson(w, p)
}

func RemovePackage(w http.ResponseWriter, r *http.Request) {
	var c ChannelDetails
	if err := readJson(r, &c); err != nil {
		http.Error(w, errors.Wrap(err, "could not parse input JSON").Error(), http.StatusBadRequest)
		return
	}

	if c.Package == nil {
		http.Error(w, "package details must be defined", http.StatusBadRequest)
		return
	}

	if status, err := validateCredentials(c.Channel, c.Password); err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	chn, err := repo.GetChannel(c.Channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := chn.RemoveSinglePackage(c.Package); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.RemovePackageCount(&models.PackageCount{
		Channel:     c.Channel,
		Package:     c.Package.Name,
		BuildString: c.Package.BuildString,
		BuildNumber: c.Package.BuildNumber,
		Version:     c.Package.Version,
		Platform:    c.Package.Platform,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok(w)
}

func RemoveAllPackages(w http.ResponseWriter, r *http.Request) {
	var c ChannelDetails
	if err := readJson(r, &c); err != nil {
		http.Error(w, errors.Wrap(err, "could not parse input JSON").Error(), http.StatusBadRequest)
		return
	}

	if status, err := validateCredentials(c.Channel, c.Password); err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	chn, err := repo.GetChannel(c.Channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	numDeleted, err := chn.RemovePackageAllVersions(chi.URLParam(r, "pkg"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	toJson(w, numDeleted)
}

func validateCredentials(channel, password string) (int, error) {
	user, err := db.GetUser(channel)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if isValid := user.HasValidPassword(password); !isValid {
		return http.StatusForbidden, errors.Errorf("password given for channel '%s' is incorrect", channel)
	}
	return http.StatusOK, nil
}
