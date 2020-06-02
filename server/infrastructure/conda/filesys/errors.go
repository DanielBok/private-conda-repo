package filesys

import "errors"

var (
	ErrPackageNotFound = errors.New("package specified does not exist")
)
