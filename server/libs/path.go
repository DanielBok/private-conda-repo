package libs

import "os"

// Checks that path or file exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return err == nil
	}
}
