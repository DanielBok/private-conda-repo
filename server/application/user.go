package application

import (
	"net/http"

	"private-conda-repo/store/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u *models.User
	if err := readJson(r, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := u.IsValid()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err = db.AddUser(u.Name, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.Password = ""

	if _, err := repo.CreateChannel(u.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJson(w, &u)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	var u *models.User
	if err := readJson(r, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.RemoveUser(u.Name, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repo.RemoveChannel(u.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok(w)
}
