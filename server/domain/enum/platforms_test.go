package enum_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/domain/enum"
)

func TestMapPlatform(t *testing.T) {
	tests := []struct {
		input    string
		expected Platform
		hasError bool
	}{
		{"linux32", LINUX32, false},
		{"linux-32", LINUX32, false},
		{"linux64", LINUX64, false},
		{"linux-64", LINUX64, false},
		{"win32", WIN32, false},
		{"win-32", WIN32, false},
		{"win64", WIN64, false},
		{"win-64", WIN64, false},
		{"osx64", OSX64, false},
		{"osx-64", OSX64, false},
		{"noarch", NOARCH, false},
		{"bad-value", "", true},
	}

	for _, test := range tests {
		p, err := MapPlatform(test.input)
		if test.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.EqualValues(t, test.expected, p)
		}
	}
}
