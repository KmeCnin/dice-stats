package input

import "errors"

// Validate checks that user input is well formatted.
func validate(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments given (expected 1)")
	}
	return nil
}
