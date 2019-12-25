package postgres

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"

	"private-conda-repo/store/models"
)

func TestStore_UserOperations(t *testing.T) {
	name := "danny"
	password := "password123"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		assert := require.New(t)
		store, err := newTestDb()
		assert.NoError(err)

		user, err := store.GetUser(name)
		assert.EqualError(err, gorm.ErrRecordNotFound.Error())

		user, err = store.AddUser(name, password)
		assert.NoError(err)
		assert.IsType(*user, models.User{})
		assert.Equal(user.Name, name)
		assert.NotEqual(user.Password, password)
		assert.True(user.HasValidPassword(password))
		assert.False(user.HasValidPassword(password + "abc"))

		user2, err := store.GetUser(name)
		assert.NoError(err)
		assert.Equal(user.Id, user2.Id)

		err = store.RemoveUser(name, "BadPassword")
		assert.EqualError(err, "incorrect credentials supplied to delete user")

		err = store.RemoveUser(name, password)
		assert.NoError(err)
	})
}
