package logger

import (
	"fmt"
	"regexp"

	"github.com/lazybark/lazyevent/v2/events"
)

// CLI is a default logger that formats event data and pushes to default
// output via fmt.Print
type CLI struct {
	lTypes []events.LogType
}

// NewCLI returns instance of CLI logger with desired log types support
func NewCLI(lTypes ...events.LogType) *CLI {
	return &CLI{lTypes: lTypes}
}

// AnsiEscaper is a regexp to find ANSI escape characters in text
var AnsiEscaper = regexp.MustCompile(`\033\[\d*m`)

// Log pushes event data into default output
func (l CLI) Log(e events.Event, timeFormat string) error {
	fmt.Print(Format(e, timeFormat))
	return nil
}

// Type returns set of types supported by the logger
func (l CLI) Type() []events.LogType { return l.lTypes }

// AvoidANSI uses regex to drop all ANSI escape sequences in event text.
func AvoidANSI(s string) string { return AnsiEscaper.ReplaceAllString(s, "") }
