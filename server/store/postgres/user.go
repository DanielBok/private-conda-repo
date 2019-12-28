package postgres

import (
	"errors"
	"strings"

	"private-conda-repo/store/models"
)

func (s *Store) AddUser(channel, password string) (*models.User, error) {
	user, err := models.NewUser(channel, password)
	if err != nil {
		return nil, err
	}

	if errs := s.db.Create(user).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return user, nil
}

func (s *Store) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if errs := s.db.Find(&users).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}
	return users, nil
}

func (s *Store) GetUser(channel string) (*models.User, error) {
	var user models.User
	if errs := s.db.
		Where("channel = ?", strings.ToLower(channel)).
		First(&user).
		GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return &user, nil
}

func (s *Store) RemoveUser(channel, password string) error {
	var user models.User
	if errs := s.db.
		Where("channel = ?", strings.ToLower(channel)).
		First(&user).
		GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	if !user.HasValidPassword(password) {
		return errors.New("incorrect credentials supplied to delete user")
	}

	if errs := s.db.Delete(&user).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	return nil
}
