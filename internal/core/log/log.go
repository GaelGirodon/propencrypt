package log

import (
	"fmt"
	"io"
	"os"
)

// Output is the application logging output.
var Output io.Writer = os.Stderr

// Print prints a message on the log output.
func Print(v ...interface{}) {
	_, _ = fmt.Fprint(Output, v...)
}

// Println prints a message and a new line on the log output.
func Println(v ...interface{}) {
	_, _ = fmt.Fprintln(Output, v...)
}

// Printf prints a formatted message on the log output.
func Printf(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(Output, format, v...)
}

// Warn prints a warning message on the log output.
func Warn(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(Output, "Warn: "+format+"\n", v...)
}

// Error prints an error message on the log output.
func Error(format string, v ...interface{}) int {
	_, _ = fmt.Fprintf(Output, "Error: "+format+"\n", v...)
	return 1
}
