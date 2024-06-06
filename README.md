but
---

Just a few utilities for error management

Why?
It's called _but_ because a but is an _and_ with a non-standard expectation



How?
You can create error templates as string constants (see but.Note) and format them later.
You can exit with/out codes if templates are filled, errors are nil and predicates are satisfied

```go
package but // import "github.com/kendfss/but"


// VARIABLES

var Separator rune = '=' // used by Must to separate output

// FUNCTIONS

func As(err error, target any) bool
    // As finds the first error in err's chain that matches target, and if one
    // is found, sets target to that error value and returns true. Otherwise,
    // it returns false.

func Exif(pred bool, fmtArgs ...any)
    // Exif calls os.Exit if pred == true, which implies err != nil see go doc but
    // FmtArgs for more info

func Exit(code int, fmtArgs ...any)
    // Exif calls os.Exit and, optionally, logs a message to the output see go doc
    // but FmtArgs for more info

func Fmt(format string) func(...any) error
    // Create a template for errors using a format string see go doc but FmtArgs
    // for more info

func GetOutput() io.Writer
    // SetOutput acquires the io.Writer to which all logs are made

func HaveNeedc[T comparable](have, want T) error
    // HaveNeedf produces "want, have" error if the two values are not equal nil
    // otherwise see go doc but Template for more general use cases

func HaveNeedf[T any](eq func(T, T) bool, have, want T) error
    // HaveNeedf produces "want, have" error see go doc but Template for more
    // general use cases

func HaveWant(have, want any) error
    // HaveWant produces "want, have" error in cases where difference is already
    // known see go doc but Template for more general use cases

func Is(err, target error) bool
    // Is reports whether any error in err's chain matches target.

func Join(errs ...error) error
    // Join returns an error that wraps the given errors. Any nil error values are
    // discarded. Join returns nil if errs contains no non-nil values. The error
    // formats as the concatenation of the strings obtained by calling the Error
    // method of each element of errs, with a newline between each string.

func LST(arg string) bool
    // LST checks if a string is line-separator-terminated

func Must(err error, fmtArgs ...any)
    // panic if an error is not nil see go doc but FmtArgs for more info

func MustBool(pred bool, fmtArgs ...any)
    // panic if an error is not nil see go doc but FmtArgs for more info

func Mustv[T any](val T, err error, fmtArgs ...any) T
    // Mustv unpacks the return value of a function that might fail panicking if it
    // does

    //     Mustv(myfunc(...))

    // see go doc but FmtArgs for more info

func Mustv2[I, O any](fn func(I) (O, error), fmtArgs ...any) func(I) O
    // see go doc but FmtArgs for more info

func New(msg string, fmtArgs ...any) error
    // Error is a errors.New && fmt.Errorf hybrid. returns nil if text is empty
    // string seen go doc but FmtArgs for more info

func ParseArgs(args ...any) (code *int, msg *string, rem []any)
    // ParseArgs parses an []any in accordance with the FmtArgs constraint see go
    // doc but FmtArgs for more info

func SetOutput[T Writer](dst T)
    // SetOutput configures the io.Writer to which all logs are made

func Should(pred bool, fmtArgs ...any) bool
    // print a non-nil error to stderr return false if err != nil see go doc but
    // FmtArgs for more info

func Strr(arg any) string
func Strrf(arg any, format string) string
    // Strrf checks if arg is a stringer, error, or string, and returns:

    //     the associated string

    // else a representation (if format != "") else ""

func Unwrap(err error) error
    // Unwrap returns the result of calling the Unwrap method on err, if err's type
    // contains an Unwrap method returning error. Otherwise, Unwrap returns nil.

    // Unwrap returns nil if the Unwrap method returns []error.


// TYPES

type FmtArgs interface {
	int | string | []any
} // Fmt functions expect trailing arguments to have one of the above types

type Note string // error message

func NewNote(fmtArgs ...any) Note
    // Create new notes see go doc but FmtArgs for more info

func (n Note) Appendf(form, sep string, args ...any) Note
    // Appendf concatenates n with the given args using the given %-form and sep

func (n Note) Error() string

func (n Note) Exif(code int, fmtArgs ...any)
    // Exif much like exit, except that it will only exit if the underlying string
    // is empty accepts format parameters as described in the but.FmtArgs interface

func (n Note) Exit(code int)
    // Exit the entire program with the given code exits with 0 if the Note is
    // empty if you only want to exit on non-nil errors, use Exif

func (n Note) Fmt(args ...any) Note

func (n Note) String() string

type Writer interface {
	io.Writer
	io.StringWriter
}
    // Writer is the type needed to log to std{in,err,out} and files



```

