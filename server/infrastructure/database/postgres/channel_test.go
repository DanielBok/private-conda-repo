package postgres_test

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"

	"private-conda-repo/domain/entity"
)

func TestPostgres_ChannelOperations(t *testing.T) {
	name := "daniel"
	password := "Password123"
	email := "daniel@gmail.com"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		assert := require.New(t)
		store, err := newTestDb(info)
		assert.NoError(err)

		chn, err := store.GetChannel(name)
		assert.EqualError(err, gorm.ErrRecordNotFound.Error())

		chn, err = store.CreateChannel(name, password, email)
		assert.NoError(err)
		assert.IsType(*chn, entity.Channel{})
		assert.Equal(chn.Channel, name)
		assert.NotEqual(chn.Password, password)
		assert.True(chn.HasValidPassword(password))
		assert.False(chn.HasValidPassword(password + "abc"))

		chn2, err := store.GetChannel(name)
		assert.NoError(err)
		assert.Equal(chn.Id, chn2.Id)

		err = store.RemoveChannel(0)
		assert.Error(err, "invalid name id")

		err = store.RemoveChannel(99999)
		assert.Error(err, "name with id does not exist")

		err = store.RemoveChannel(chn.Id)
		assert.NoError(err)

		chn, err = store.GetChannel(name)
		assert.Error(err)
		assert.Nil(chn)
	})
}
