package postgres

import "github.com/hashicorp/go-multierror"

func joinErrors(errors []error) error {
	var err error
	for _, e := range errors {
		err = multierror.Append(err, e)
	}
	return err
}
