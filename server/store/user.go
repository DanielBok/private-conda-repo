package store

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"private-conda-repo/config"
)

type User struct {
	Id       int
	Name     string // this will be the channel name as well
	Password string
}

func (s *Store) AddUser(name, password string) (*User, error) {
	user := User{
		Name:     name,
		Password: hashPassword(password, s.conf.Salt),
	}

	if errs := s.db.Create(&user).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return &user, nil
}

func (s *Store) GetUser(name string) (*User, error) {
	var user User
	if errs := s.db.Where("name = ?", name).First(&user).GetErrors(); len(errs) > 0 {
		if len(errs) == 1 {
			return nil, errs[0]
		}
		return nil, joinErrors(errs)
	}

	return &user, nil
}

func (s *Store) RemoveUser(name, password string) error {
	var user User
	if errs := s.db.Where("name = ?", name).First(&user).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	if !user.IsValid(password) {
		return errors.New("incorrect credentials supplied to delete user")
	}

	if errs := s.db.Delete(&user).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	return nil
}

func (u *User) IsValid(password string) bool {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	return hashPassword(password, conf.Salt) == u.Password
}

func hashPassword(plainPassword, salt string) string {
	h := sha256.New()
	h.Write([]byte(plainPassword + salt))
	return fmt.Sprintf("%x", h.Sum(nil))[:64]
}
