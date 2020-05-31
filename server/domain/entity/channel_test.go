package entity

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestUser_IsValid(t *testing.T) {
	goodEmail := "daniel@gmail.com"
	goodPassword := "good" // Must be >= 2 characters
	goodUsername := "good" // Must be >= 4 characters

	tests := []struct {
		user     Channel
		domain   string
		hasError bool
	}{
		{
			Channel{
				Channel:  "b", // too short a name
				Password: goodPassword,
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			Channel{
				Channel:  "A-really-long-name-with-valid-characters-but-is-more-than-the-limit-of-50-characters",
				Password: goodPassword,
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			Channel{
				Channel:  goodUsername,
				Password: "bad",
				Email:    goodEmail,
			},
			"",
			true,
		},
		{
			Channel{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    "badEmail",
			},
			"",
			true,
		},
		{
			Channel{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    "email@bad-domain.com",
			},
			"yahoo.com",
			true,
		},
		{
			Channel{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			"yahoo.com",
			true,
		},
		{
			Channel{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			"gmail.com",
			false,
		},
		{
			Channel{
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
	c := NewChannel("daniel", "good-password", "daniel@gmail.com")
	valid := c.HasValidPassword("bad-password")
	require.False(t, valid)

	valid = c.HasValidPassword("good-password")
	require.True(t, valid)
}
