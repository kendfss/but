package but

import (
	"fmt"
	"os"
	"reflect"
)

type (
	Bug struct {
		code uint8
		msg  string
	}
)

func (b Bug) Code() uint8 {
	return b.code
}

func (b Bug) Error() string {
	return b.msg
}

func (this Bug) Wrapf(msg string) Bug {
	return Bug{
		code: this.code,
		msg:  fmt.Sprintf(msg, this.msg),
	}
}

func (b Bug) F(index int, msg string, args ...any) Bug {
	args = insert(args, index)
	return Bug{
		code: b.code,
		msg:  fmt.Sprintf(msg, args...),
	}
}

func (b Bug) Warn() {
	if b.msg != "" {
		os.Stderr.WriteString(lined(b.msg))
	}
}
func (b Bug) Exit() {
	b.Exif()
	os.Exit(b.Int())
}
func (b Bug) Exif() {
	if b.code != 0 && b.msg == "" {
		output.WriteString(b.msg)
		os.Exit(int(b.code))
	}
}

func (b Bug) changed(ctr int, args ...any) *Bug {
	out := b.Clone()
	switch arg := args[0]; arg.(type) {
	case uint8, int:
		out.code = arg.(uint8)
	case string:
		out.msg = arg.(string)
	default:
		Exit(1, "bug.Change: arg #%d, %q, is %q want %q or %q", ctr, Strr(arg), reflect.TypeOf(arg), "uint8", "string")
	}
	return b.changed(ctr, args[1:]...)
}

func (b *Bug) Change(args ...any) {
	b = b.changed(0, args...)
}

// Clone returns a copy of a Bug at a new memory location
func (b Bug) Clone() *Bug {
	return &Bug{
		code: b.code,
		msg:  b.msg,
	}
}

// Explain appends a colon, space char, and the given string to a Bug's message
func (b Bug) Explain(s string) Bug {
	return Bug{
		code: b.code,
		msg:  fmt.Sprintf("%s: %s", b.msg, s),
	}
}

// Name prepends the given string, a colon, and a space char to a Bug's message
func (b Bug) Name(s string) Bug {
	return Bug{
		code: b.code,
		msg:  fmt.Sprintf("%s: %s", s, b.msg),
	}
}

// Int returns the Bug's exit code as an int
func (b Bug) Int() int { return int(b.code) }

// String returns the Bug's message
func (b Bug) String() string { return b.msg }

// From derives a Bug from a std.string
func (*Bug) From(err string) *Bug {
	return &Bug{
		code: 1,
		msg:  err,
	}

}

// For derives a Bug from a std.error
func (*Bug) For(err error) *Bug {
	return new(Bug).From(err.Error())
}

// Eq checks if this Bug has the same pointer, code, or message, as that one
func (this *Bug) Eq(that *Bug) bool {
	return this == that || this.code == that.code || this.msg == that.msg
}

// LST checks if a string is line-separator-terminated
func LST(arg string) bool {
	chars := []rune(arg)
	return chars[len(chars)-1] == '\n'
}

// Strrf checks if arg is a stringer, error, or string, and returns:
//		the associated string
// else	a representation (if format != "")
// else ""
func Strrf(arg any, format string) string {
	switch arg.(type) {
	case fmt.Stringer:
		return arg.(fmt.Stringer).String()
	case error:
		return arg.(error).Error()
	case string:
		return arg.(string)
	default:
		if format == "" {
			return ""
		}
		return fmt.Sprint(format, arg)
	}
}

func Strr(arg any) string {
	return Strrf(arg, "%#v")
}

// insert inserts the values v... into s at index i,
// returning the modified slice.
// In the returned slice r, r[i] == v[0].
// insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func insert[E any](s []E, i int, args ...E) []E {
	tot := len(s) + len(args)
	if tot <= cap(s) {
		s2 := s[:tot]
		copy(s2[i+len(args):], s[i:])
		copy(s2[i:], args)
		return s2
	}
	s2 := make([]E, tot)
	copy(s2, s[:i])
	copy(s2[i:], args)
	copy(s2[i+len(args):], s[i:])
	return s2
}
