package postgres

import (
	"errors"

	"private-conda-repo/store/models"
)

func (s *Store) AddUser(name, password string) (*models.User, error) {
	user, err := models.NewUser(name, password)
	if err != nil {
		return nil, err
	}

	if errs := s.db.Create(user).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return user, nil
}

func (s *Store) GetUser(name string) (*models.User, error) {
	var user models.User
	if errs := s.db.Where("name = ?", name).First(&user).GetErrors(); len(errs) > 0 {
		if len(errs) == 1 {
			return nil, errs[0]
		}
		return nil, joinErrors(errs)
	}

	return &user, nil
}

func (s *Store) RemoveUser(name, password string) error {
	var user models.User
	if errs := s.db.Where("name = ?", name).First(&user).GetErrors(); len(errs) > 0 {
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
