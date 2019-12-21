package models

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUser_IsValid(t *testing.T) {
	u := User{
		Name:     "bad",
		Password: "bad",
	}
	err := u.IsValid()
	assert.Error(t, err, "Name and password are both too short")

	u.Name = "good"
	err = u.IsValid()
	assert.Error(t, err, "Password should still be short")

	u.Password = "good"
	err = u.IsValid()
	assert.NoError(t, err)
}

func TestUser_HasValidPassword(t *testing.T) {
	viper.Set("salt", "test-salt")

	u, err := NewUser("daniel", "good-password")
	assert.NoError(t, err)

	valid := u.HasValidPassword("bad-password")
	assert.False(t, valid)

	valid = u.HasValidPassword("good-password")
	assert.True(t, valid)
}
