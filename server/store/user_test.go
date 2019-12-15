package store

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestStore_UserOperations(t *testing.T) {
	name := "danny"
	password := "password123"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		assert := assert.New(t)
		store, err := newTestDb()
		assert.NoError(err)

		user, err := store.GetUser(name)
		assert.EqualError(err, gorm.ErrRecordNotFound.Error())

		user, err = store.AddUser(name, password)
		if assert.NoError(err) {
			assert.IsType(*user, User{})
		}
		assert.Equal(user.Name, name)
		assert.NotEqual(user.Password, password)
		assert.True(user.IsValid(password))
		assert.False(user.IsValid(password + "abc"))

		user2, err := store.GetUser(name)
		if assert.NoError(err) {
			assert.Equal(user.Id, user2.Id)
		}

		err = store.RemoveUser(name, "BadPassword")
		assert.EqualError(err, "incorrect credentials supplied to delete user")

		err = store.RemoveUser(name, password)
		assert.NoError(err)
	})
}
