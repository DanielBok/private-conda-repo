package api

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"private-conda-repo/api/dto"
	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/entity"
	"private-conda-repo/libs"
)

type PackageHandler struct {
	DB           interfaces.DataAccessLayer
	Decompressor interfaces.Decompressor
	FileSys      interfaces.FileSys
}

func (h *PackageHandler) ListAllPackages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels, err := h.FileSys.ListAllChannels()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var output []*dto.ChannelData
		for _, c := range channels {
			channelData, err := c.GetChannelData()
			if err != nil {
				http.Error(w, "error getting meta info", http.StatusInternalServerError)
				return
			}

			output = append(output, dto.ToChannelDataDto(channelData, c.Name())...)
		}

		toJson(w, output)
	}
}

func (h *PackageHandler) ListPackagesInChannel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "channel")
		chn, err := h.FileSys.GetChannel(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not find name/channel with name '%s'", name), http.StatusBadRequest)
			return
		}

		channelData, err := chn.GetChannelData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		toJson(w, dto.ToChannelDataDto(channelData, name))
	}
}

func (h *PackageHandler) FetchPackageDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "channel")
		pkg := chi.URLParam(r, "pkg")

		chn, err := h.DB.GetChannel(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		details, err := h.DB.GetPackageCounts(chn.Id, pkg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if len(details) == 0 {
			http.Error(w, "package does not exist in channel", http.StatusNotFound)
			return
		}

		folder, err := h.FileSys.GetChannel(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		channelData, err := folder.GetChannelData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get latest package
		var latest *dto.ChannelData

		for _, c := range dto.ToChannelDataDto(channelData, name) {
			if c.Name != pkg {
				continue
			}
			if latest == nil {
				latest = c
			} else if (latest.Version == nil || c.Version == nil) && latest.Timestamp < c.Timestamp {
				latest = c
			} else if *latest.Version < *c.Version {
				latest = c
			}
		}

		toJson(w, dto.PackageDetails{
			Channel: name,
			Package: pkg,
			Details: details,
			Latest:  latest,
		})
	}
}

func (h *PackageHandler) RemoveAllPackages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chn, _, err := h.getChannelPackage(r)
		if errors.Is(err, ErrInvalidCredential) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		folder, err := h.FileSys.GetChannel(chn.Channel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		numDeleted, err := folder.RemovePackageAllVersions(chi.URLParam(r, "pkg"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		toJson(w, numDeleted)
	}
}

func (h *PackageHandler) RemovePackage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chn, pkg, err := h.getChannelPackage(r)
		if errors.Is(err, ErrInvalidCredential) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if pkg == nil {
			http.Error(w, "package details must be defined", http.StatusBadRequest)
			return
		}

		folder, err := h.FileSys.GetChannel(chn.Channel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := folder.RemoveSinglePackage(pkg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := h.DB.RemovePackageCount(&entity.PackageCount{
			ChannelId:   chn.Id,
			Package:     pkg.Name,
			BuildString: pkg.BuildString,
			BuildNumber: pkg.BuildNumber,
			Version:     pkg.Version,
			Platform:    pkg.Platform,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ok(w)
	}
}

func (h *PackageHandler) UploadPackage() http.HandlerFunc {
	getChannel := func(name, password string) (*entity.Channel, error) {
		c, err := h.DB.GetChannel(name)
		if err != nil {
			return nil, err
		}

		if !c.HasValidPassword(password) {
			return nil, ErrInvalidCredential
		}

		return c, nil
	}

	savePackageToDisk := func(file multipart.File, channelName string, fixes []string) (*dto.PackageDto, error) {
		// retrieve package meta data
		pkg, err := h.Decompressor.RetrieveMetadata(file)
		if err != nil {

			return nil, err
		}
		defer pkg.Close()

		if err := pkg.Package.Validate(); err != nil {
			return nil, err
		}

		// get channel and upload file
		chn, err := h.FileSys.GetChannel(channelName)
		if err != nil {
			return nil, err
		}

		f, err := pkg.Open()
		if err != nil {
			return nil, ErrOpeningCondaPackage
		}
		defer pkg.Close()

		// save package to disk
		p, err := chn.AddPackage(f, pkg.Package, fixes)
		if err != nil {
			return nil, ErrSavingPackageToDisk
		}

		return p, nil
	}

	createPackageCount := func(p *dto.PackageDto, id int) error {
		if _, err := h.DB.CreatePackageCount(p.ToPackageCount(id)); err != nil {
			return errors.New("could not create package count record")
		}

		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 24)
		if err != nil {
			http.Error(w, "could not parse form data", http.StatusInternalServerError)
			return
		}

		// check channel and password
		chn, err := getChannel(r.FormValue("channel"), r.FormValue("password"))
		if errors.Is(err, ErrInvalidCredential) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, ErrParsingFormFile.Error(), http.StatusBadRequest)
			return
		}
		defer libs.IOCloser(file)

		// save package
		fixes := strings.Split(strings.TrimSpace(strings.ToLower(r.FormValue("fixes"))), ",")
		p, err := savePackageToDisk(file, chn.Channel, fixes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = createPackageCount(p, chn.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// return outcome
		toJson(w, p)
	}
}

func (h *PackageHandler) getChannelPackage(r *http.Request) (*entity.Channel, *dto.PackageDto, error) {
	var d *dto.ChannelPackage
	if err := readJson(r, &d); err != nil {
		return nil, nil, err
	}

	c, err := h.DB.GetChannel(d.Channel)
	if err != nil {
		return nil, nil, err
	}

	if !c.HasValidPassword(d.Password) {
		return nil, nil, ErrInvalidCredential
	}

	return c, d.Package, nil
}
