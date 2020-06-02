package libs

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// Closes the stream and logs a message at level Error on the standard logger
// if an error is encountered while closing the stream.
func IOCloser(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Error(err)
	}
}

// Closes the stream and logs a message at level Error on the standard logger
// if an error is encountered while closing the stream. The error is appended
// to the back of the formatted string so there is no need to add a "%v" verb
// for the error
func IOCloserf(closer io.Closer, format string, args ...interface{}) {
	err := closer.Close()
	if err != nil {
		args = append(args, err)
		log.Errorf(format+": %v", args...)
	}
}
