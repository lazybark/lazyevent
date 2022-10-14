package logger

import (
	"fmt"
	"os"

	"github.com/lazybark/go-helpers/fsw"
	"github.com/lazybark/lazyevent/v2/events"
)

type PlaintextFileLogger struct {
	//pureText means logger will print out only log text itself
	pureText bool
	lTypes   []events.LogType
	file     *os.File
}

// NewPlaintext returns logger capable of appending strings to text file
func NewPlaintext(path string, pureText bool, truncate bool, lTypes ...events.LogType) (*PlaintextFileLogger, error) {
	f, err := fsw.MakePathToFile(path, truncate)
	if err != nil {
		return nil, fmt.Errorf("[NewPlaintext] error making log path: %w", err)
	}

	return &PlaintextFileLogger{pureText: pureText, lTypes: lTypes, file: f}, nil
}

// Log pushes event data into default output
func (l PlaintextFileLogger) Log(e events.Event, timeFormat string) error {
	log := ""
	if l.pureText {
		log = FormatPureText(e)
	} else {
		log = Format(e, timeFormat)
	}
	_, err := l.file.WriteString(log)
	if err != nil {
		return err
	}

	return nil
}

// Type returns set of types supported by the logger
func (l PlaintextFileLogger) Type() []events.LogType { return l.lTypes }
