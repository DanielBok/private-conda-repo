package models

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestUser_IsValid(t *testing.T) {
	goodEmail := "daniel@gmail.com"
	goodPassword := "good" // Must be >= 4 characters
	goodUsername := "good" // Must be >= 4 characters

	tests := []struct {
		user     User
		domain   string
		hasError bool
	}{
		{
			User{
				Channel:  "bad",
				Password: goodPassword,
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			User{
				Channel:  "A-really-long-name-with-valid-characters-but-is-more-than-the-limit-of-50-characters",
				Password: goodPassword,
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			User{
				Channel:  goodUsername,
				Password: "bad",
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			User{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    "badEmail",
			},
			"",
			true,
		},
		{
			User{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    "email@bad-domain.com",
			},
			"yahoo.com",
			true,
		},
		{
			User{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			"yahoo.com",
			true,
		},
		{
			User{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			"gmail.com",
			false,
		},
		{
			User{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			"",
			false,
		},
	}

	for _, test := range tests {
		viper.Set("user.email_domain", test.domain)
		err := test.user.IsValid()
		if test.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestUser_HasValidPassword(t *testing.T) {
	viper.Set("salt", "test-salt")

	u, err := NewUser("daniel", "good-password", "daniel@gmail.com")
	require.NoError(t, err)

	valid := u.HasValidPassword("bad-password")
	require.False(t, valid)

	valid = u.HasValidPassword("good-password")
	require.True(t, valid)
}
