package models

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestUser_IsValid(t *testing.T) {
	u := User{
		Channel:  "bad",
		Password: "bad",
	}
	err := u.IsValid()
	require.Error(t, err, "Channel and password are both too short")

	u.Channel = "good"
	err = u.IsValid()
	require.Error(t, err, "Password should still be short")

	u.Password = "good"
	err = u.IsValid()
	require.NoError(t, err)
}

func TestUser_HasValidPassword(t *testing.T) {
	viper.Set("salt", "test-salt")

	u, err := NewUser("daniel", "good-password")
	require.NoError(t, err)

	valid := u.HasValidPassword("bad-password")
	require.False(t, valid)

	valid = u.HasValidPassword("good-password")
	require.True(t, valid)
}
