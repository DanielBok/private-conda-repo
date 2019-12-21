package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/go-multierror"

	"private-conda-repo/config"
)

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"` // this will be the channel name as well
	Password string `json:"password,omitempty"`
}

func (u *User) HasValidPassword(password string) bool {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	return hashPassword(password, conf.Salt) == u.Password
}

func (u *User) IsValid() (err error) {
	u.Name = strings.TrimSpace(u.Name)
	u.Password = strings.TrimSpace(u.Password)

	if len(u.Name) < 4 {
		err = multierror.Append(err, errors.New("username must be >= 4 characters"))
	}
	if len(u.Password) < 4 {
		err = multierror.Append(err, errors.New("password must be >= 4 characters"))
	}

	return
}

func NewUser(name, password string) (*User, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Password: hashPassword(password, conf.Salt),
	}, nil
}

func hashPassword(plainPassword, salt string) string {
	h := sha256.New()
	h.Write([]byte(plainPassword + salt))
	return fmt.Sprintf("%x", h.Sum(nil))[:64]
}
