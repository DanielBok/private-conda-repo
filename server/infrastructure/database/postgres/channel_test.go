package postgres

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"

	"private-conda-repo/domain/entity"
)

func TestStore_UserOperations(t *testing.T) {
	channel := "daniel"
	password := "Password123"
	email := "daniel@gmail.com"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		assert := require.New(t)
		store, err := newTestDb(info)
		assert.NoError(err)

		chn, err := store.GetChannel(channel)
		assert.EqualError(err, gorm.ErrRecordNotFound.Error())

		chn, err = store.CreateChannel(channel, password, email)
		assert.NoError(err)
		assert.IsType(*chn, entity.Channel{})
		assert.Equal(chn.Channel, channel)
		assert.NotEqual(chn.Password, password)
		assert.True(chn.HasValidPassword(password, store.salt))
		assert.False(chn.HasValidPassword(password+"abc", store.salt))

		chn2, err := store.GetChannel(channel)
		assert.NoError(err)
		assert.Equal(chn.Id, chn2.Id)

		err = store.RemoveChannel(channel, "BadPassword")
		assert.EqualError(err, "incorrect credentials supplied to delete chn")

		err = store.RemoveChannel(channel, password)
		assert.NoError(err)
	})
}
