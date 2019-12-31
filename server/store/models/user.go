package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"

	"private-conda-repo/config"
)

type User struct {
	Id       int       `json:"-"`
	Channel  string    `json:"channel"` // this will be the channel name as well
	Password string    `json:"password,omitempty"`
	Email    string    `json:"email"`
	JoinDate time.Time `json:"join_date"`
}

func (u *User) HasValidPassword(password string) bool {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	return hashPassword(password, conf.UserConfig.Salt) == u.Password
}

func (u *User) IsValid() (err error) {
	nameRegex := regexp.MustCompile(`^\w[\w\-]{3,49}$`)

	conf, err := config.New()
	if err != nil {
		return err
	}

	u.Channel = strings.TrimSpace(u.Channel)
	u.Password = strings.TrimSpace(u.Password)
	u.Email = strings.TrimSpace(u.Email)

	if !nameRegex.MatchString(u.Channel) {
		err = multierror.Append(err, errors.New("user/channel name length must be between [4, 50] characters and can only be alphanumeric with dashes"))
	}
	if len(u.Password) < 4 {
		err = multierror.Append(err, errors.New("password must be >= 4 characters"))
	}
	if !conf.UserConfig.ValidateEmail(u.Email) {
		err = multierror.Append(err, errors.New("invalid email address"))
	}

	return
}

func NewUser(name, password, email string) (*User, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	return &User{
		Channel:  strings.ToLower(strings.TrimSpace(name)),
		Password: hashPassword(password, conf.UserConfig.Salt),
		Email:    strings.TrimSpace(strings.ToLower(email)),
		JoinDate: time.Now().UTC(),
	}, nil
}

func hashPassword(plainPassword, salt string) string {
	h := sha256.New()
	h.Write([]byte(plainPassword + salt))
	return fmt.Sprintf("%x", h.Sum(nil))[:64]
}
