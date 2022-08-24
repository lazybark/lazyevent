package logger

import (
	"fmt"
	"os"

	"github.com/lazybark/go-helpers/fsw"
	"github.com/lazybark/lazyevent/v2/events"
)

type PlaintextFileLogger struct {
	lTypes []events.LogType
	file   *os.File
}

// NewPlaintext returns logger capable of appending strings to text file
func NewPlaintext(path string, truncate bool, lTypes ...events.LogType) (*PlaintextFileLogger, error) {
	f, err := fsw.MakePathToFile(path, truncate)
	if err != nil {
		return nil, fmt.Errorf("[NewPlaintext] error making log path: %w", err)
	}

	return &PlaintextFileLogger{lTypes: lTypes, file: f}, nil
}

// Log pushes event data into default output
func (l PlaintextFileLogger) Log(e events.Event, timeFormat string) error {
	_, err := l.file.WriteString(Format(e, timeFormat))
	if err != nil {
		return err
	}

	return nil
}

// Type returns set of types supported by the logger
func (l PlaintextFileLogger) Type() []events.LogType { return l.lTypes }
