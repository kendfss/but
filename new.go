package but

import "fmt"

// Error is a errors.New && fmt.Errorf hybrid.
// returns nil if text is empty string
// seen go doc but FmtArgs for more info
func New(msg string, fmtArgs ...any) error {
	if msg != "" {
		return fmt.Errorf(msg, fmtArgs...)
	}
	return nil
}

// Create new notes
// see go doc but FmtArgs for more info
func NewNote(fmtArgs ...any) Note {
	_, msg, args := ParseArgs(fmtArgs)
	return Note(fmt.Sprintf(*msg, args))
}
