package simplelog

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// Info prints `v` into standard output (via log) with a green prefix "info:".
func Info(v ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	m := fmt.Sprintln(v...)
	log.Printf("%s %s", green("[info]"), m)
}

// Error prints `v` into standard output (via log) with a red prefix "error:".
func Error(v ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	m := fmt.Sprintln(v...)
	log.Printf("%s %s", red("[error]"), m)
}

// ErrorD prints `v` into standard output (via log) with a red prefix "error:" along with stack trace.
func ErrorD(v ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	m := fmt.Sprintln(v...)
	log.Printf("%s %s", red("[error]"), m)
	err := errors.WithStack(fmt.Errorf(m))
	fmt.Printf("%+v\n", err)
}

// Infof is the formatted version of Info().
func Infof(format string, v ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	m := fmt.Sprintf(format, v...)
	log.Printf("%s %s", green("[info]"), m)
}

// Errorf is the formatted version of Error().
func Errorf(format string, v ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	m := fmt.Sprintf(format, v...)
	log.Printf("%s %s", red("[error]"), m)
}
