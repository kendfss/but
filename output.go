package but

import (
	"io"
	"os"
)

// Writer is the type needed to log to std{i,e,o} and files
type Writer interface {
	io.Writer
	io.StringWriter
}

var output Writer = os.Stderr

// SetOutput configures the io.Writer to which all logs are made
func SetOutput[T Writer](dst T) {
	output = new(stringWriter).from(dst)
}

// SetOutput acquires the io.Writer to which all logs are made
func GetOutput() io.Writer {
	return output
}

type stringWriter struct {
	target io.Writer
}

func (*stringWriter) from(w Writer) *stringWriter {
	return &stringWriter{
		target: w,
	}
}

func (sw *stringWriter) Write(buf []byte) (n int, err error) {
	return sw.target.Write(buf)
}

func (sw *stringWriter) WriteString(msg string) (n int, err error) {
	return io.WriteString(sw.target, lined(msg))
}
