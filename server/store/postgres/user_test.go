package postgres

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"

	"private-conda-repo/store/models"
)

func TestStore_UserOperations(t *testing.T) {
	channel := "daniel"
	password := "Password123"
	email := "daniel@gmail.com"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		assert := require.New(t)
		store, err := newTestDb()
		assert.NoError(err)

		user, err := store.GetUser(channel)
		assert.EqualError(err, gorm.ErrRecordNotFound.Error())

		user, err = store.AddUser(channel, password, email)
		assert.NoError(err)
		assert.IsType(*user, models.User{})
		assert.Equal(user.Channel, channel)
		assert.NotEqual(user.Password, password)
		assert.True(user.HasValidPassword(password))
		assert.False(user.HasValidPassword(password + "abc"))

		user2, err := store.GetUser(channel)
		assert.NoError(err)
		assert.Equal(user.Id, user2.Id)

		err = store.RemoveUser(channel, "BadPassword")
		assert.EqualError(err, "incorrect credentials supplied to delete user")

		err = store.RemoveUser(channel, password)
		assert.NoError(err)
	})
}
