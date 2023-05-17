package but

import (
	"fmt"
	"os"
	"strings"
)

type Note string

// Appendf concatenates n with the given args using the given %-form and sep
func (n Note) Appendf(form, sep string, args ...any) Note {
	form = form + strings.Repeat(sep+form, len(args))
	args = append([]any{n}, args...)
	return Note(fmt.Sprintf(form, args...))
}

func (n Note) Fmt(args ...any) Note {
	return Note(fmt.Sprintf(n.String(), args...))
}

func (n Note) String() string {
	return string(n)
}
func (n Note) Error() string {
	return string(n)
}

// Exit the entire program with the given code
// exits with 0 if the Note is empty
// if you only want to exit on non-nil errors, use Exif
func (n Note) Exit(code int) {
	if n == "" {
		os.Exit(0)
	}
	os.Exit(1)
}

// Exif much like exit, except that it will only exit
// if the underlying string is empty
// accepts format parameters as described in the but.FmtArgs interface
func (n Note) Exif(code int, fmtArgs ...any) {
	if n != "" {
		_, msg, args := ParseArgs(fmtArgs)
		output.WriteString(fmt.Sprintf(*msg, args...))
		os.Exit(code)
	}
}
