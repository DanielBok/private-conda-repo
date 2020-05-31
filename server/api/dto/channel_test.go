package dto_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/api/dto"
)

func TestChannelDto_IsValid(t *testing.T) {
	t.Parallel()

	goodEmail := "daniel@gmail.com"
	goodPassword := "good" // Must be >= 2 characters
	goodUsername := "good" // Must be >= 4 characters

	tests := []struct {
		Channel  ChannelDto
		HasError bool
		Message  string
	}{
		{
			ChannelDto{
				Channel:  "b", // too short a name
				Password: goodPassword,
				Email:    goodEmail,
			},
			true,
			"channel name is too short",
		},
		{
			ChannelDto{
				Channel:  "A-really-long-name-with-valid-characters-but-is-more-than-the-limit-of-50-characters",
				Password: goodPassword,
				Email:    goodEmail,
			},
			true,
			"channel name is more than 50 characters (too long)",
		},
		{
			ChannelDto{
				Channel:  goodUsername,
				Password: "bad",
				Email:    goodEmail,
			},
			true,
			"password is too short",
		},
		{
			ChannelDto{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    "badEmail",
			},
			true,
			"email is invalid",
		},
		{
			ChannelDto{
				Channel:  goodUsername,
				Password: goodPassword,
				Email:    goodEmail,
			},
			false,
			"should not fail as DTO is valid",
		},
	}

	for _, test := range tests {
		err := test.Channel.IsValid()
		if test.HasError {
			require.Error(t, err, test.Message)
		} else {
			require.NoError(t, err, test.Message)
		}
	}
}
