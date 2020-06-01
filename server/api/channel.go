package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"private-conda-repo/api/dto"
	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/entity"
)

type ChannelHandler struct {
	DB      interfaces.DataAccessLayer
	FileSys interfaces.FileSys
}

func (h *ChannelHandler) ListChannels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels, err := h.DB.GetAllChannels()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, user := range channels {
			user.Password = ""
		}

		toJson(w, &channels)
	}
}

func (h *ChannelHandler) CreateChannel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get channel details from json request
		d, err := h.getChannelDto(r, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ensure that channel does not already exist
		if c, err := h.DB.GetChannel(d.Channel); err != nil && err != gorm.ErrRecordNotFound {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if c != nil {
			http.Error(w, fmt.Sprintf("channel '%s' already exists", c.Channel), http.StatusBadRequest)
			return
		}

		// create channel in database
		channel, err := h.DB.CreateChannel(d.Channel, d.Password, d.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		d = dto.NewChannelDto(channel)

		// create channel's folder
		folder, err := h.FileSys.CreateChannel(d.Channel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := folder.Index(nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		toJson(w, &d)
	}
}

func (h *ChannelHandler) GetChannelInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "channel")
		channel, err := h.DB.GetChannel(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		d := dto.NewChannelDto(channel)
		toJson(w, &d)
	}
}

func (h *ChannelHandler) RemoveChannel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := h.getChannel(r, false)
		if errors.Is(err, ErrInvalidCredential) {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.DB.RemoveChannel(c.Id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.FileSys.RemoveChannel(c.Channel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ok(w)
	}
}

func (h *ChannelHandler) CheckChannel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := h.getChannel(r, false)
		if errors.Is(err, ErrInvalidCredential) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ok(w)
	}
}

func (h *ChannelHandler) getChannel(r *http.Request, validateDto bool) (*entity.Channel, error) {
	d, err := h.getChannelDto(r, validateDto)
	if err != nil {
		return nil, err
	}

	c, err := h.DB.GetChannel(d.Channel)
	if err != nil {
		return nil, err
	}

	if !c.HasValidPassword(d.Password) {
		return nil, ErrInvalidCredential
	}

	return c, nil
}

func (h *ChannelHandler) getChannelDto(r *http.Request, validateDto bool) (*dto.ChannelDto, error) {
	var d *dto.ChannelDto
	if err := readJson(r, &d); err != nil {
		return nil, err
	}

	if validateDto {
		// validate that request payload is okay
		if err := d.IsValid(); err != nil {
			return nil, err
		}
	}

	return d, nil
}
