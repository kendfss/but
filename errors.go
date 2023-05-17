package but

import (
	"fmt"
	"os"
)

type FmtArgs interface {
	int | string | []any
	// in that order
	// the int is only looked for in context where it is meaningful
}

var Separator rune = '=' // used by Must to separate output

// HaveWant produces "want, have" error in cases where difference is already known
// see go doc but Template for more general use cases
func HaveWant(have, want any) error {
	return fmt.Errorf("have: %v, want: %v", have, want)
}

// HaveNeedf produces "want, have" error if the two values are not equal
// nil otherwise
// see go doc but Template for more general use cases
func HaveNeedc[T comparable](have, want T) error {
	if have == want {
		return nil
	}
	return fmt.Errorf("have: %#v, want: %#v", have, want)
}

// HaveNeedf produces "want, have" error
// see go doc but Template for more general use cases
func HaveNeedf[T any](eq func(T, T) bool, have, want T) error {
	if eq(have, want) {
		return nil
	}
	return fmt.Errorf("have: %#v, want: %#v", have, want)
}

// Exif calls os.Exit and, optionally, logs a message to the output
// see go doc but FmtArgs for more info
func Exit(code int, fmtArgs ...any) {
	defer os.Exit(code)
	_, msg, args := ParseArgs(fmtArgs...)
	*msg = fmt.Sprintf(*msg, args...)
	output.WriteString(*msg)
}

// Exif calls os.Exit if pred == true, which implies err != nil
// see go doc but FmtArgs for more info
func Exif(pred bool, fmtArgs ...any) {
	c := 1
	if pred {
		code, msg, args := ParseArgs(fmtArgs...)
		if code == nil {
			code = &c
		}
		if msg != nil {
			*msg = fmt.Sprintf(*msg, args...)
			fmt.Fprintln(output, *msg)
		}
		os.Exit(*code)
	}
}

// panic if an error is not nil
// see go doc but FmtArgs for more info
func Must(err error, fmtArgs ...any) {
	if err != nil {
		_, msg, args := ParseArgs(fmtArgs...)
		panic(fmt.Sprintf(*msg, args...))
	}
}

// panic if an error is not nil
// see go doc but FmtArgs for more info
func MustBool(pred bool, fmtArgs ...any) {
	if !pred {
		_, msg, args := ParseArgs(fmtArgs...)
		panic(fmt.Sprintf(*msg, args...))
	}
}

// Mustv unpacks the return value of a function that might fail
// panicking if it does
// 		Mustv(myfunc(...))
// see go doc but FmtArgs for more info
func Mustv[T any](val T, err error, fmtArgs ...any) T {
	Must(err, fmtArgs...)
	return val
}

// see go doc but FmtArgs for more info
func Mustv2[I, O any](fn func(I) (O, error), fmtArgs ...any) func(I) O {
	return func(i I) O {
		a, b := fn(i)
		return Mustv(a, b, fmtArgs...)
	}
}

// print a non-nil error to stderr
// return false if err != nil
// see go doc but FmtArgs for more info
func Should(pred bool, fmtArgs ...any) bool {
	if pred {
		_, msg, args := ParseArgs(fmtArgs...)
		*msg = fmt.Sprintf(*msg, args...)
		output.WriteString(*msg)
	}
	return pred
}

// Create a template for errors using a format string
// see go doc but FmtArgs for more info
func Fmt(format string) func(...any) error {
	return func(fmtArgs ...any) error {
		return New(format, fmtArgs...)
	}
}
