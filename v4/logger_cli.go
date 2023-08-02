package logger

import (
	"fmt"
	"regexp"
)

// CLI is a default logger that formats event data and pushes to default
// output via fmt.Print
type CLI struct {
	//pureText means logger will print out only log text itself
	//Without time, date, etc, but with color formatting.
	pureText bool
	lTypes   []LogType
}

// NewCLI returns instance of CLI logger with desired log types support
func NewCLI(pureText bool, lTypes ...LogType) *CLI {
	return &CLI{pureText: pureText, lTypes: lTypes}
}

// AnsiEscaper is a regexp to find ANSI escape characters in text
var AnsiEscaper = regexp.MustCompile(`\033\[\d*m`)

// Log pushes event data into default output
func (l *CLI) Log(e Event, timeFormat string) error {
	log := ""
	if l.pureText {
		log = FormatOutputPureText(e)
	} else {
		log = FormatOutput(e, timeFormat)
	}
	if e.Format != None {
		log = FormatColors(e.Format, log)
	}
	fmt.Print(log)

	return nil
}

// Type returns set of types supported by the logger
func (l *CLI) Type() []LogType { return l.lTypes }

// AvoidANSI uses regex to drop all ANSI escape sequences in event text.
func AvoidANSI(s string) string { return AnsiEscaper.ReplaceAllString(s, "") }
