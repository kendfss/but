package but

import "github.com/kendfss/oprs"

// lined returns a line-separator-terminated string
func lined(arg string) string {
	chars := []rune(arg)
	if chars[len(chars)-1] == '\n' {
		return arg
	}
	return arg + "\n"
}

// ParseArgs parses an []any in accordance with the FmtArgs constraint
// see go doc but FmtArgs for more info
func ParseArgs(args ...any) (code *int, msg *string, rem []any) {
	for i := 0; len(args) >= 1; i++ {
		code = oprs.Assert[int](args[0])
		if code == nil || i > 1 {
			msg = oprs.Assert[string](args[0])
		}
		args = args[1:]
	}
	rem = args
	return
}
