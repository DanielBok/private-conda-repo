package postgres

import (
	"github.com/hashicorp/go-multierror"
)

func joinErrors(errors []error) error {
	if len(errors) == 1 {
		return errors[0]
	}

	var err error

	for _, e := range errors {
		err = multierror.Append(err, e)
	}
	return err
}
