package but

/*

	The contents of this file are ported from the builtin "errors" library for documentational/expediential purposes

*/

import "errors"

// As finds the first error in err's chain that matches target, and if one is found, sets target to that error value and returns true. Otherwise, it returns false.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Is reports whether any error in err's chain matches target.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Join returns an error that wraps the given errors. Any nil error values are
// discarded. Join returns nil if errs contains no non-nil values. The error
// formats as the concatenation of the strings obtained by calling the Error
// method of each element of errs, with a newline between each string.
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's type
// contains an Unwrap method returning error. Otherwise, Unwrap returns nil.
//
// Unwrap returns nil if the Unwrap method returns []error.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
