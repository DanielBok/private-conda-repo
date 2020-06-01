package api

import (
	"errors"
)

var (
	ErrInvalidCredential   = errors.New("invalid credential")
	ErrOpeningCondaPackage = errors.New("could not open compressed package archive")
	ErrSavingPackageToDisk = errors.New("could not save conda package to disk")
	ErrParsingFormFile     = errors.New("could not parse uploaded file. Please ensure that you have uploaded a valid file with 'file' as the form key")
)
