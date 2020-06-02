package libs

import (
	"os"
)

// Check that path or file exists. This is a simple check and should suffice
// for most use cases.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
