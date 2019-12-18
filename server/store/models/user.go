package models

import (
	"crypto/sha256"
	"fmt"
	"log"

	"private-conda-repo/config"
)

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"` // this will be the channel name as well
	Password string `json:"password,omitempty"`
}

func (u *User) IsValid(password string) bool {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	return hashPassword(password, conf.Salt) == u.Password
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
