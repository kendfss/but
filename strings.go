package but

import (
	"fmt"
)

// LST checks if a string is line-separator-terminated
func LST(arg string) bool {
	chars := []rune(arg)
	return chars[len(chars)-1] == '\n'
}

// Strrf checks if arg is a stringer, error, or string, and returns:
//
//	the associated string
//
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
