package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

func ListUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, user := range users {
		user.Password = ""
	}

	toJson(w, &users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.IsValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, err := db.GetUser(u.Channel); err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if user != nil {
		http.Error(w, fmt.Sprintf("channel '%s' already exists", u.Channel), http.StatusBadRequest)
		return
	}

	u, err = db.AddUser(u.Channel, u.Password, u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.Password = ""

	chn, err := repo.CreateChannel(u.Channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := chn.Index(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, &u)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user")
	user, err := db.GetUser(username)
	if err != nil {
		http.Error(w, errors.Wrapf(err, "could not find user: %s", username).Error(), http.StatusBadRequest)
		return
	}
	user.Password = ""

	toJson(w, user)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.RemoveUser(u.Channel, u.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repo.RemoveChannel(u.Channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok(w)
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actual, err := db.GetUser(u.Channel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if actual.HasValidPassword(u.Password) {
		ok(w)
	} else {
		http.Error(w, "invalid credentials", http.StatusForbidden)
	}
}

func getUser(r *http.Request) (*models.User, error) {
	var u *models.User
	if err := readJson(r, &u); err != nil {
		return nil, err
	}
	return u, nil
}
